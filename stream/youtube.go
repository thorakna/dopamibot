package stream

import (
	"bytes"
	"dpmmusicbot/utils"
	"fmt"
	"os/exec"
	"strings"
	"time"

	ytsearch "github.com/AnjanaMadu/YTSearch"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

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
	if !utils.IsAnUrl(url) {
		return nil, &errorString{"Hata: Url Gelmedi"}
	}
	return &url, nil
}

func GetStreamFromYTMusicId(musicId string) (*string, error) {
	start := time.Now()
	cmd := exec.Command("youtube-dl", musicId, "--get-url", "-f", "251")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Println("Getting stream with ytdl from ytmusic took ", time.Since(start))
	if err != nil {
		return nil, err
	}
	url := out.String()
	if !utils.IsAnUrl(url) {
		return nil, &errorString{"Hata: Url Gelmedi"}
	}
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
	if !utils.IsAnUrl(str) {
		return nil, &errorString{"Hata: Url Gelmedi"}
	}
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
