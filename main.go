package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarin/discordgo"
)

func main() {
	var Session, _ = discordgo.New()
	Session.Token = ""

	err := Session.Open()
	if err != nil {
		log.Printf("Error establisihng connection to Discord, %s\n", err)
		os.Exit(1)
	}

	log.Printf("Now running, press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up
	Session.Close()
}
