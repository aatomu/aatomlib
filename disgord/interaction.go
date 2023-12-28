package disgord

import (
	"github.com/bwmarrin/discordgo"
)

type InteractionResponse struct {
	discord     *discordgo.Session
	interaction *discordgo.Interaction

	Response *discordgo.InteractionResponse
}

// InteractionResponseのFlag
const Invisible uint64 = 1 << 6

func NewInteractionResponse(d *discordgo.Session, i *discordgo.Interaction) (ir *InteractionResponse) {
	ir.discord = d
	ir.interaction = i
	return
}

// Interaction Reply Message
// Flags Usual: Invisible
func (i *InteractionResponse) Reply(resData *discordgo.InteractionResponseData) error {
	i.Response = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: resData,
	}
	return i.discord.InteractionRespond(i.interaction, i.Response)
}

// Interaction Thinking Message
// Please after Follow()
func (i *InteractionResponse) Thinking(isInvisible bool) error {
	i.Response = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	}
	if isInvisible {
		i.Response.Data.Flags = discordgo.MessageFlagsEphemeral
	}
	return i.discord.InteractionRespond(i.interaction, i.Response)
}

// Interaction Window Message
// Component only usual: Input#string
func (i *InteractionResponse) Window(title, customID string, comps ...*ComponentLine) error {
	i.Response = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			Title:      title,
			CustomID:   customID,
			Components: ComponentParse(comps...),
		},
	}
	return i.discord.InteractionRespond(i.interaction, i.Response)
}

// Interaction Edit Message
func (i *InteractionResponse) Edit(newData *discordgo.WebhookEdit) (*discordgo.Message, error) {
	return i.discord.InteractionResponseEdit(i.interaction, newData)
}

// Interaction FollowUP Message
// Flags Usual: Invisible
func (i *InteractionResponse) Follow(newData *discordgo.WebhookParams) (*discordgo.Message, error) {
	return i.discord.FollowupMessageCreate(i.interaction, true, newData)
}
