package stream

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	ytsearch "github.com/AnjanaMadu/YTSearch"
)

func GetStreamFromYT(query string) (*string, error) {
	start := time.Now()
	cmd := exec.Command("youtube-dl", "ytsearch:"+query, "--get-url", "-f", "251")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Println("Getting stream with ytdl search took ", time.Since(start))
	if err != nil {
		return nil, err
	}
	url := out.String()
	return &url, nil
}

func GetStreamFromYTWithId(input string) (*string, error) {
	start := time.Now()
	cmd := exec.Command("youtube-dl", "--get-url", "-f", "251", "https://www.youtube.com/watch?v="+input)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Println("Getting stream with id took ", time.Since(start))
	if err != nil {
		return nil, err
	}
	str := out.String()
	return &str, nil
}

func YTSearchWithAPI(query string) (string, error) {
	results, err := ytsearch.Search(query)
	var firstResult string
	if err != nil {
		return "", err
	}
	for _, v := range results {
		if strings.TrimSpace(v.VideoId) != "" {
			firstResult = v.VideoId
			break
		}
	}
	return firstResult, nil
}

// TODO: Pull Youtube Playlist
