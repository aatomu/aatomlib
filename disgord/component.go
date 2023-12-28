package disgord

import (
	"github.com/bwmarrin/discordgo"
)

type ComponentLine struct {
	row discordgo.ActionsRow
}

func ComponentNewLine() *ComponentLine {
	return &ComponentLine{}
}

// if Style:5 => Request: URL
// unles Style:5 => Request: "CustomID"
// max 5 buttons in 1line
// Request: Label, Style.
func (cl *ComponentLine) Button(button discordgo.Button) *ComponentLine {
	if len(cl.row.Components) > 5 {
		panic("Up to 5 buttons per line")
	}
	cl.row.Components = append(cl.row.Components, button)
	return cl
}

// FuncReq: AddLine()
// Request: CustomID, Options.
// SelectMenuOption Request: Label,Value.
// if MinValue or MaxValue != 0 multi select
func (cl *ComponentLine) SelectMenu(selectMenu discordgo.SelectMenu) *ComponentLine {
	if len(cl.row.Components) != 0 {
		panic("SelectMenu() only 1line in 1method")
	}
	cl.row.Components = append(cl.row.Components, selectMenu)
	return cl
}

// FuncReq: AddLine()
// Request: CustomID,Label,Style.
// Styles : discordgo.TextInputShort,discordgo.TextInputParagraph,.
// Interaction Response Only
func (cl *ComponentLine) Input(textInput discordgo.TextInput) *ComponentLine {
	if len(cl.row.Components) != 0 {
		panic("Input() only 1line in 1method")
	}
	cl.row.Components = append(cl.row.Components, textInput)
	return cl
}

// Componentをdiscordgoで使えるように
func ComponentParse(lines ...*ComponentLine) (result []discordgo.MessageComponent) {
	for _, line := range lines {
		result = append(result, line.row)
	}
	return
}
