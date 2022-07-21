package main

import (
	"dpmmusicbot/stream"
	"fmt"
	"time"
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
	start := time.Now()
	videoId, err := stream.YTSearchWithAPI("Fiyakalı Yok oluşlar")
	fmt.Printf("YT Search took %s\n", time.Since(start))
	opusUrl, err := stream.GetUrlFromYTWithId(videoId)
	fmt.Printf("YT Opus link extracting took %s\n\n", time.Since(start))

	if err != nil {
		return
	}

	fmt.Println(*opusUrl)
}
