package stream

import (
	"bytes"
	"dpmmusicbot/utils"
	"fmt"
	"os/exec"
	"time"
)

// TODO: Soundcloud search & get stream url

func GetStreamFromSCWithId(input string) (*string, error) {
	start := time.Now()
	cmd := exec.Command("youtube-dl", "--get-url", "https://soundcloud.com/"+input)
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
