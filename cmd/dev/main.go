package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ARC5RF/DiscordAnalyticsBot/internals/command"
)

func main() {
	go command.ExecPrefix("tailwind ", "go", "run", "./cmd/twcss")
	go command.ExecPrefix("analytics ", "go", "run", ".")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	_ = <-interrupt
	fmt.Println()
}
