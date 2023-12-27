package disgord

import "github.com/bwmarrin/discordgo"

func HaveRole(discord *discordgo.Session, guildID, userID, roleID string) (have bool, err error) {
	//Check user role list
	userRoleList, err := discord.GuildMember(guildID, userID)
	if err != nil {
		return false, err
	}

	guildRoleList, err := discord.GuildRoles(guildID)
	if err != nil {
		return false, err
	}

	//Search by guild role list
	for _, guildRole := range guildRoleList {
		if guildRole.ID == roleID {
			for _, userRole := range userRoleList.Roles {
				if userRole == guildRole.ID {
					return true, nil
				}
			}
		}
	}
	return false, nil
}
