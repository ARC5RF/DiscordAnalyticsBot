package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"reflect"

	"net/http"

	"github.com/ARC5RF/DiscordAnalyticsBot/discord"
	"github.com/ARC5RF/DiscordAnalyticsBot/webui"
	"github.com/ARC5RF/DiscordAnalyticsBot/webui/hub"
	"github.com/bwmarrin/discordgo"
)

// formatting and printing values to the console.
// logging messages to the console.
// Used for build HTTP servers and clients.

// Port we listen on.
const portNum string = ":8080"

// Handler functions.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Info page")
}

func main() {

	wui := webui.New()

	discord_session := discord.New()
	if err := discord_session.Open(); err != nil {
		panic(err)
	}
	discord_session.GoInvisible()

	discord_session.AddHandler(func(s *discordgo.Session, data interface{}) {
		v := reflect.ValueOf(data)

		switch e := data.(type) {
		case *discordgo.Event:
			data, err := json.Marshal(&e)
			if err != nil {
				return
			}
			wui.Broadcast(data)
		default:
			fmt.Println(v.Type())
		}
	})

	wui.Receive(func(m *hub.ClientCommand) {
		fmt.Println("recieved command from web UI", m.Type)
	})

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	go wui.Start()

	_ = <-interrupt
	fmt.Println()
	fmt.Println("Bot Gracefully shutting down...")

	wui.Stop()

}
