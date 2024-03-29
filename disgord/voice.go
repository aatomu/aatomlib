package disgord

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func GetUserVoiceState(discord *discordgo.Session, userID string) *discordgo.VoiceState {
	for _, guild := range discord.State.Guilds {
		for _, vs := range guild.VoiceStates {
			if vs.UserID == userID {
				return vs
			}
		}
	}
	return nil
}

func JoinUserVCchannel(discord *discordgo.Session, userID string, micMute, speakerMute bool) (vc *discordgo.VoiceConnection, err error) {
	vs := GetUserVoiceState(discord, userID)
	if vs == nil {
		return nil, fmt.Errorf("user doesn't join voice chat")
	}

	vc, err = discord.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, micMute, speakerMute)
	return vc, err
}

// 音再生
// end := make(<-chan bool, 1)
func PlayAudioFile(vcSession *discordgo.VoiceConnection, filename string, speed float64, pitch float64, volume float64, isPlayback bool, end <-chan bool) error {
	if err := vcSession.Speaking(true); err != nil {
		return err
	}
	defer vcSession.Speaking(false)

	done := make(chan error)
	stream := NewFileEncodeStream(vcSession, filename, EncodeOpts{
		Compression: 1,
		AudioFilter: fmt.Sprintf("aresample=24000,asetrate=24000*%.2f/100,atempo=100/%.2f*%.2f,volume=%.2f", pitch*100, pitch*100, speed, volume),
	}, done)

	ticker := time.NewTicker(time.Second)
	if isPlayback {
		defer ticker.Stop()
	} else {
		ticker.Stop()
	}

	for {
		select {
		case err := <-done:
			if err != nil && err != io.EOF {
				return err
			}
			stream.Cleanup()
			return nil
		case <-ticker.C:
			log.Printf("Sending Now... : Playback:%.2f(x%.2f)", stream.Status.Time.Seconds(), stream.Status.Speed)
		case <-end:
			stream.Cleanup()
			return nil
		}
	}
}
