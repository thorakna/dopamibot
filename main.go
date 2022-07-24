package main

import (
	"dpmmusicbot/stream"
	"fmt"
)

func main() {
	// discord, err := discordgo.New(utils.GoDotEnvVariable("DCTOKEN"))
	// if err != nil {
	// 	fmt.Println("Error creating discord session,", err)
	// 	return
	// }
	// err = discord.Open()
	// if err != nil {
	// 	fmt.Println("Error opening connection,", err)
	// 	return
	// }
	// fmt.Println("Started")
	// <-make(chan struct{})

	opusUrl, err := stream.GetStreamFromYT("Bury the light")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*opusUrl)
}
