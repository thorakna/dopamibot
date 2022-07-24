package main

import (
	"dpmmusicbot/utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Dg *discordgo.Session

func initialiseSession() {
	discord, err := discordgo.New("Bot " + utils.GoDotEnvVariable("DCTOKEN"))
	if err != nil {
		fmt.Println("Error creating discord session,", err)
		return
	}
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	Dg = discord

	fmt.Println("Started")
}

func ConnectToDiscord() {
	initialiseSession()

	Dg.AddHandler(messageCreate)

	Dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Ctrl + C to disconnect
	fmt.Println("To quit, use Ctrl + C")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "dpmhelp" {
		_, err := s.ChannelMessageSend(m.ChannelID, "I have nothing to offer but the storm that is approacting... yet.")
		if err != nil {
			fmt.Println("error sending message:", err)
		}
	}

}
