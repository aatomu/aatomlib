package disgord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type MessageData struct {
	GuildID   string
	Guild     *discordgo.Guild
	ChannelID string
	Channel   *discordgo.Channel
	User      *discordgo.User

	MessageID string
	Message   *discordgo.Message

	FormatText string
}

type VoiceStateData struct {
	GuildID   string
	Guild     *discordgo.Guild
	ChannelID string
	Channel   *discordgo.Channel
	User      *discordgo.User

	Status       VoiceStatus
	UpdateStatus VoiceStatus

	FormatText string
}

type VoiceStatus struct {
	ChannelJoin  bool
	ServerDeaf   bool
	ServerMute   bool
	ClientDeaf   bool
	ClientMute   bool
	ClientGoLive bool
	ClientCam    bool
}

type ReactionData struct {
	GuildID   string
	Guild     *discordgo.Guild
	ChannelID string
	Channel   *discordgo.Channel
	User      *discordgo.User

	MessageID string
	Message   *discordgo.Message

	Emoji     discordgo.Emoji
	EmojiIcon string

	FormatText string
}

type InteractionData struct {
	GuildID   string
	Guild     *discordgo.Guild
	ChannelID string
	Channel   *discordgo.Channel
	User      *discordgo.User

	InteractionType discordgo.InteractionType
	Command         discordgo.ApplicationCommandInteractionData
	CommandOptions  map[string]ApplicationCommandOptions
	Component       discordgo.MessageComponentInteractionData
	Modal           discordgo.ModalSubmitInteractionData
	ModalValues     map[string]string

	FormatText string
}

type ApplicationCommandOptions struct {
	*discordgo.ApplicationCommandInteractionDataOption
}

func MessageParse(discord *discordgo.Session, m *discordgo.Message) (md MessageData) {
	md.GuildID = m.GuildID
	md.Guild, _ = discord.Guild(md.GuildID)
	guildName := ""
	if md.Guild != nil {
		guildName = md.Guild.Name
	}

	md.ChannelID = m.ChannelID
	md.Channel, _ = discord.Channel(md.ChannelID)
	channelName := ""
	if md.Channel != nil {
		channelName = md.Channel.Name
	}

	md.User = m.Author
	userName := ""
	if md.User != nil {
		userName = md.User.String()
	}

	md.MessageID = m.ID
	md.Message, _ = discord.ChannelMessage(md.ChannelID, md.MessageID)
	content := ""
	if md.Message != nil {
		content = md.Message.Content
	}

	attachmentsURL := ""
	if len(m.Attachments) > 0 {
		attachmentURLs := []string{}
		for _, file := range m.Attachments {
			attachmentURLs = append(attachmentURLs, file.URL)
		}
		attachmentsURL = fmt.Sprintf("Attachments:\"%s\"  ", strings.Join(attachmentURLs, " "))
	}

	// Formatter
	md.FormatText = fmt.Sprintf(`Guild:"%s"  Channel:"%s"  %s<%s>: %s`, guildName, channelName, attachmentsURL, userName, content)
	return
}

func VoiceStateParse(discord *discordgo.Session, v *discordgo.VoiceStateUpdate) (vd VoiceStateData) {
	vd.GuildID = v.GuildID
	vd.Guild, _ = discord.Guild(vd.GuildID)
	guildName := ""
	if vd.Guild != nil {
		guildName = vd.Guild.Name
	}

	vd.ChannelID = v.ChannelID
	vd.Channel, _ = discord.Channel(vd.ChannelID)
	channelName := ""
	if vd.Channel != nil {
		channelName = vd.Channel.Name
	}

	vd.User, _ = discord.User(v.UserID)
	userName := ""
	if vd.User != nil {
		userName = vd.User.String()
	}

	vd.Status = VoiceStatus{
		ChannelJoin:  (v.ChannelID != ""),
		ServerDeaf:   v.Deaf,
		ServerMute:   v.Mute,
		ClientDeaf:   v.SelfDeaf,
		ClientMute:   v.SelfMute,
		ClientGoLive: v.SelfStream,
		ClientCam:    v.SelfVideo,
	}

	if v.BeforeUpdate == nil {
		vd.UpdateStatus.ChannelJoin = true
	} else {
		vd.UpdateStatus = VoiceStatus{
			ChannelJoin:  (v.ChannelID != v.BeforeUpdate.ChannelID),
			ServerDeaf:   (v.Deaf != v.BeforeUpdate.Deaf),
			ServerMute:   (v.Mute != v.BeforeUpdate.Mute),
			ClientDeaf:   (v.SelfDeaf != v.BeforeUpdate.SelfDeaf),
			ClientMute:   (v.SelfMute != v.BeforeUpdate.SelfMute),
			ClientGoLive: (v.SelfStream != v.BeforeUpdate.SelfStream),
			ClientCam:    (v.SelfVideo != v.BeforeUpdate.SelfVideo),
		}
	}

	// Formatter
	vd.FormatText = fmt.Sprintf(`Guild:"%s"  Channel:"%s"  <%s>`, guildName, channelName, userName)
	switch {
	case vd.UpdateStatus.ChannelJoin:
		vd.FormatText = fmt.Sprintf("%s ChangeTo ChannelJoin:\"%t\"", vd.FormatText, vd.Status.ChannelJoin)
	case vd.UpdateStatus.ServerDeaf:
		vd.FormatText = fmt.Sprintf("%s ChangeTo ServerDeaf:\"%t\"", vd.FormatText, vd.Status.ServerDeaf)
	case vd.UpdateStatus.ServerMute:
		vd.FormatText = fmt.Sprintf("%s ChangeTo ServerMute:\"%t\"", vd.FormatText, vd.Status.ServerMute)
	case vd.UpdateStatus.ClientDeaf:
		vd.FormatText = fmt.Sprintf("%s ChangeTo ClientDeaf:\"%t\"", vd.FormatText, vd.Status.ClientDeaf)
	case vd.UpdateStatus.ClientMute:
		vd.FormatText = fmt.Sprintf("%s ChangeTo ClientMute:\"%t\"", vd.FormatText, vd.Status.ClientMute)
	case vd.UpdateStatus.ClientGoLive:
		vd.FormatText = fmt.Sprintf("%s ChangeTo ClientGoLive:\"%t\"", vd.FormatText, vd.Status.ClientGoLive)
	case vd.UpdateStatus.ClientCam:
		vd.FormatText = fmt.Sprintf("%s ChangeTo ClientCam:\"%t\"", vd.FormatText, vd.Status.ClientCam)
	}
	return
}

func ReactionParse(discord *discordgo.Session, r *discordgo.MessageReaction, eventName string) (rd ReactionData) {
	rd.GuildID = r.GuildID
	rd.Guild, _ = discord.Guild(rd.GuildID)
	guildName := ""
	if rd.Guild != nil {
		guildName = rd.Guild.Name
	}

	rd.ChannelID = r.ChannelID
	rd.Channel, _ = discord.Channel(rd.ChannelID)
	channelName := ""
	if rd.Channel != nil {
		channelName = rd.Channel.Name
	}

	rd.User, _ = discord.User(r.UserID)
	userName := ""
	if rd.User != nil {
		userName = rd.User.String()
	}

	rd.MessageID = r.MessageID
	rd.Message, _ = discord.ChannelMessage(rd.ChannelID, rd.MessageID)
	content := ""
	author := ""
	if rd.Message != nil {
		content = rd.Message.Content
		author = rd.Message.Author.String()
	}

	rd.Emoji = r.Emoji
	rd.EmojiIcon = rd.Emoji.Name

	// Delete After New lines
	if strings.Contains(content, "\n") {
		content = strings.SplitN(content, "\n", 2)[0]
	}

	split := strings.Split(content, "")
	if len(split) > 20 {
		content = strings.Join(split[:20], "") + ".."
	}

	// Formatter
	rd.FormatText = fmt.Sprintf(`Guild:"%s"  Channel:"%s"  <%s> Type:"%s" Emoji:"%s" => <%s> %s`, guildName, channelName, userName, eventName, rd.EmojiIcon, author, content)
	return
}

func InteractionParse(discord *discordgo.Session, i *discordgo.Interaction) (id InteractionData) {
	id.GuildID = i.GuildID
	id.Guild, _ = discord.Guild(id.GuildID)
	guildName := ""
	if id.Guild != nil {
		guildName = id.Guild.Name
	}

	id.ChannelID = i.ChannelID
	id.Channel, _ = discord.Channel(id.ChannelID)
	channelName := ""
	if id.Channel != nil {
		channelName = id.Channel.Name
	}

	if i.User != nil {
		id.User = i.User
	} else {
		id.User = i.Member.User
	}
	userName := ""
	if id.User != nil {
		userName = id.User.String()
	}

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		id.InteractionType = discordgo.InteractionApplicationCommand
		id.Command = i.ApplicationCommandData()
		id.CommandOptions = map[string]ApplicationCommandOptions{}
		// Optionデータ保存
		for _, optionData := range id.Command.Options {
			id.CommandOptions[optionData.Name] = ApplicationCommandOptions{optionData}
		}
		//表示
		id.FormatText = fmt.Sprintf(`Guild:"%s"  Channel:"%s"  <%s> /%s %v`, guildName, channelName, userName, id.Command.Name, id.CommandOptions)

	case discordgo.InteractionMessageComponent:
		id.InteractionType = discordgo.InteractionMessageComponent
		id.Component = i.MessageComponentData()
		//表示
		id.FormatText = fmt.Sprintf(`Guild:"%s"  Channel:"%s"  <%s> Component_ID:"%s"`, guildName, channelName, userName, id.Component.CustomID)

	case discordgo.InteractionModalSubmit:
		id.InteractionType = discordgo.InteractionModalSubmit
		id.Modal = i.ModalSubmitData()

		id.ModalValues = map[string]string{}
		for _, comp := range id.Modal.Components {
			data := comp.(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)
			id.ModalValues[data.CustomID] = data.Value
		}
		//表示
		id.FormatText = fmt.Sprintf(`Guild:"%s"  Channel:"%s"  <%s> Modal_ID:"%s" %+v`, guildName, channelName, userName, id.Modal.CustomID, id.ModalValues)
	}
	return
}

// discordgo unsupported Attachments
func (data ApplicationCommandOptions) AttachmentValue(i InteractionData) *discordgo.MessageAttachment {
	return i.Command.Resolved.Attachments[data.Value.(string)]
}
