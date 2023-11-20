package lookup

import "github.com/bwmarrin/discordgo"

func Guild(s *discordgo.Session, gid string) (*discordgo.Guild, error) {
	g, err := s.State.Guild(gid)
	if err == discordgo.ErrStateNotFound {
		g, err = s.Guild(gid)
		if err != nil {
			return nil, err
		}
		return g, nil
	}
	return g, err
}
