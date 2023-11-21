package session

import (
	"fmt"
	"os"

	_ "github.com/ARC5RF/DiscordAnalyticsBot/internals/cfg"

	"github.com/bwmarrin/discordgo"
)

var (
	DISCORD_TOKEN = os.Getenv("DISCORD_TOKEN")
)

type DiscordClientInstance struct {
	*discordgo.Session
}

func New() *DiscordClientInstance {
	if DISCORD_TOKEN == "" {
		panic("DISCORD_TOKEN env empty")
	}

	session, err := discordgo.New("Bot " + DISCORD_TOKEN)
	if err != nil {
		panic(err)
	}
	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	return &DiscordClientInstance{Session: session}
}

func (session *DiscordClientInstance) GoInvisible() {
	err := session.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "invisible",
	})
	if err != nil {
		fmt.Println("error updating status,", err)
		panic(err)
	}
}
