package disgord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Command生成
// all guild when guildID == ""
func InteractionCommandCreate(discord *discordgo.Session, guildID string, commands []*discordgo.ApplicationCommand) {
	for _, command := range commands {
		_, err := discord.ApplicationCommandCreate(discord.State.User.ID, guildID, command)
		if err != nil {
			fmt.Printf("Failed Create Command \"%s\"\n", command.Name)
			panic(err)
		}
	}
}

// Command削除
// all guild when guildID == ""
func InteractionCommandDelete(discord *discordgo.Session, guildID string, name string) error {
	cmd, err := discord.ApplicationCommands(discord.State.User.ID, guildID)
	if err != nil {
		return err
	}
	for _, cmdData := range cmd {
		if cmdData.Name == name {
			return discord.ApplicationCommandDelete(discord.State.User.ID, guildID, cmdData.ID)
		}
	}
	return fmt.Errorf("not found \"%s\"", name)
}

func FindValue(interaction *discordgo.InteractionCreate, key string) (value []interface{}) {
	for _, data := range interaction.ApplicationCommandData().Options {
		if data.Name == key {
			value = append(value, data.Value)
		}
	}
	return
}
