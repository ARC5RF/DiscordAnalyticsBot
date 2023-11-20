package perms

import (
	"github.com/ARC5RF/DiscordAnalyticsBot/bot/lookup"
	"github.com/bwmarrin/discordgo"
)

func ComputeBasePermissions(s *discordgo.Session, member *discordgo.Member, guild *discordgo.Guild) int64 {
	if guild.OwnerID == member.User.ID {
		return ALL
	}

	var everyone *discordgo.Role
	for _, role := range guild.Roles {
		if role.ID == guild.ID {
			everyone = role
			break
		}
	}

	permissions := everyone.Permissions

	for _, mrole := range member.Roles {
		var role *discordgo.Role
		for _, grole := range guild.Roles {
			if mrole == grole.ID {
				role = grole
				break
			}
		}

		permissions |= role.Permissions
	}

	if (permissions & ADMINISTRATOR) == ADMINISTRATOR {
		return ALL
	}

	return permissions
}

func ComputeOverwrites(base_permissions int64, member *discordgo.Member, channel *discordgo.Channel) int64 {
	// ADMINISTRATOR overrides any potential permission overwrites, so there is nothing to do here.
	if base_permissions&ADMINISTRATOR == ADMINISTRATOR {
		return ALL
	}

	permissions := base_permissions
	overwrites := channel.PermissionOverwrites

	// Find (@everyone) role overwrite and apply it.
	var overwrite_everyone *discordgo.PermissionOverwrite
	for _, overwrite := range overwrites {
		if overwrite.ID == channel.GuildID {
			overwrite_everyone = overwrite
			break
		}
	}
	if overwrite_everyone != nil {
		permissions &= ^overwrite_everyone.Deny
		permissions |= overwrite_everyone.Allow
	}

	// Apply role specific overwrites.
	var allow int64 = NONE
	var deny int64 = NONE
	for _, role_id := range member.Roles {
		//overwrite_role = overwrites.get(role_id)

		var overwrite_role *discordgo.PermissionOverwrite
		for _, overwrite := range overwrites {
			if overwrite.ID == role_id {
				overwrite_role = overwrite
				break
			}
		}

		if overwrite_role != nil {
			allow |= overwrite_role.Allow
			deny |= overwrite_role.Deny
		}
	}

	permissions &= ^deny
	permissions |= allow

	// Apply member specific overwrite if it exist.
	// overwrite_member = overwrites.get(member.User.ID)
	var overwrite_member *discordgo.PermissionOverwrite
	for _, overwrite := range overwrites {
		if overwrite.ID == member.User.ID {
			overwrite_member = overwrite
			break
		}
	}

	if overwrite_member != nil {
		permissions &= ^overwrite_member.Deny
		permissions |= overwrite_member.Allow
	}

	return permissions
}

func ComputePermissions(s *discordgo.Session, member *discordgo.Member, channel *discordgo.Channel) (int64, error) {
	guild, err := lookup.Guild(s, channel.GuildID)
	if err != nil {
		return 0, err
	}
	base_permissions := ComputeBasePermissions(s, member, guild)
	return ComputeOverwrites(base_permissions, member, channel), nil
}
