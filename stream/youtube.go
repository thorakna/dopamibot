package stream

import (
	"bytes"
	"os/exec"

	ytsearch "github.com/AnjanaMadu/YTSearch"
)

func GetUrlFromYTWithId(input string) (*string, error) {
	cmd := exec.Command("youtube-dl", "--get-url", "-f", "251", input)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	str := out.String()
	return &str, nil
}

func YTSearchWithAPI(query string) (string, error) {
	results, err := ytsearch.Search("fiyakalı yok oluşlar")
	if err != nil {
		return "", err
	}
	return results[0].VideoId, nil
}
