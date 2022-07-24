package stream

import (
	"regexp"
)


// Playlist çekilecek
// example URI https://open.spotify.com/playlist/37i9dQZF1EZT8cSzhJA9lL

func ParsePlaylistId(playlistUrl string) (playlistId string, isSuccess bool) {
	spotifyRegexp := regexp.MustCompile("https?://open.spotify.com/playlist/(.+)")
	matches := spotifyRegexp.FindStringSubmatch(playlistUrl)
	if len(matches) > 1 {
		return matches[1], true
	}
	return "", false;
}
