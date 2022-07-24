package main

import (
	"dpmmusicbot/utils"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

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

	if strings.HasPrefix(m.Content, "dpm") {
		// Find the channel that the message came from.
		c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			// Could not find channel.
			return
		}

		if m.Content == "dpmhelp" {
			_, err := s.ChannelMessageSend(m.ChannelID, "I have nothing to offer but the storm that is approacting... yet.")
			if err != nil {
				fmt.Println("error sending message:", err)
			}
		} else if m.Content == "dpmjoin" {
			// Find the guild for that channel.
			g, err := s.State.Guild(c.GuildID)
			if err != nil {
				// Could not find guild.
				return
			}

			// Look for the message sender in that guild's current voice states.
			for _, vs := range g.VoiceStates {
				if vs.UserID == m.Author.ID {
					err = playSound(s, g.ID, vs.ChannelID)
					if err != nil {
						fmt.Println("Error playing sound:", err)
					}

					return
				}
			}
		}
	}
}

func playSound(s *discordgo.Session, guildID, channelID string) (err error) {
	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	vc.Speaking(true)

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(3000 * time.Millisecond)
	vc.Speaking(false)
	vc.Disconnect()

	return nil
}
