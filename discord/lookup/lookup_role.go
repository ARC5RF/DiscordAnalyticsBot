package lookup

import "github.com/bwmarrin/discordgo"

func Role(s *discordgo.Session, gid, rid string) (*discordgo.Role, error) {
	g, err := Guild(s, gid)
	if err != nil {
		return nil, err
	}

	r, err := s.State.Role(g.ID, rid)
	if err == discordgo.ErrStateNotFound {
		roles, err := s.GuildRoles(g.ID)
		if err != nil {
			return nil, err
		}
		var r *discordgo.Role
		for _, role := range roles {
			err = s.State.RoleAdd(gid, role)
			if err != nil {
				return nil, err
			}
			if role.ID == rid {
				r = role
			}
		}
		return r, nil
	}
	return r, err
}
