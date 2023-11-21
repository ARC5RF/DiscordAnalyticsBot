package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/ARC5RF/DiscordAnalyticsBot/internals/command"
	"github.com/ARC5RF/DiscordAnalyticsBot/internals/download"
	"github.com/ARC5RF/fsup"
)

func main() {
	p, err := fsup.FromWD("cmd")
	if err != nil {
		panic(err)
	}
	e := filepath.Join(p, "twcss", "tailwind")
	if _, err := os.Stat(e); errors.Is(err, os.ErrNotExist) {
		l := "https://github.com/tailwindlabs/tailwindcss/releases/download/v3.3.5/tailwindcss-linux-x64"
		err = download.Executable(e, l)
		if err != nil {
			panic(err)
		}
	}

	log.Fatal(command.Exec(e, "--watch=always", "-i", "./assets/private/twmain.css", "-o", "./assets/public/tw.css"))
}
