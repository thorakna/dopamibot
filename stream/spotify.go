package stream

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// Playlist çekilecek
// example URI https://open.spotify.com/playlist/37i9dQZF1EZT8cSzhJA9lL

type SpotifyPlaylistResponse struct {
	Href  string `json:"href"`
	Items []struct {
		AddedAt time.Time `json:"added_at"`
		AddedBy struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"added_by"`
		IsLocal      bool        `json:"is_local"`
		PrimaryColor interface{} `json:"primary_color"`
		Track        struct {
			Album struct {
				AlbumType string `json:"album_type"`
				Artists   []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"artists"`
				AvailableMarkets []string `json:"available_markets"`
				ExternalUrls     struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href   string `json:"href"`
				ID     string `json:"id"`
				Images []struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"images"`
				Name                 string `json:"name"`
				ReleaseDate          string `json:"release_date"`
				ReleaseDatePrecision string `json:"release_date_precision"`
				TotalTracks          int    `json:"total_tracks"`
				Type                 string `json:"type"`
				URI                  string `json:"uri"`
			} `json:"album"`
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			DiscNumber       int      `json:"disc_number"`
			DurationMs       int      `json:"duration_ms"`
			Episode          bool     `json:"episode"`
			Explicit         bool     `json:"explicit"`
			ExternalIds      struct {
				Isrc string `json:"isrc"`
			} `json:"external_ids"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href        string `json:"href"`
			ID          string `json:"id"`
			IsLocal     bool   `json:"is_local"`
			Name        string `json:"name"`
			Popularity  int    `json:"popularity"`
			PreviewURL  string `json:"preview_url"`
			Track       bool   `json:"track"`
			TrackNumber int    `json:"track_number"`
			Type        string `json:"type"`
			URI         string `json:"uri"`
		} `json:"track"`
		VideoThumbnail struct {
			URL interface{} `json:"url"`
		} `json:"video_thumbnail"`
	} `json:"items"`
	Limit    int         `json:"limit"`
	Next     interface{} `json:"next"`
	Offset   int         `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int         `json:"total"`
}

func ParsePlaylistId(playlistUrl string) (playlistId string, isSuccess bool) {
	spotifyRegexp := regexp.MustCompile("https?://open.spotify.com/playlist/(.+)")
	matches := spotifyRegexp.FindStringSubmatch(playlistUrl)
	if len(matches) > 1 {
		return matches[1], true
	}
	return "", false
}

func GetPlaylistTracks(playlistId string, oAuthToken string) []string {
	client := &http.Client{}
	url := "https://api.spotify.com/v1/playlists/" + playlistId + "/tracks"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+oAuthToken)

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("GetPlaylistTracks: Request Failed with status code %d!\n", resp.StatusCode)

		return nil
	}

	var response SpotifyPlaylistResponse
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("GetPlaylistTracks: Couldn't unmarshal JSON")
		return nil
	}

	var tracks []string = make([]string, len(response.Items))
	for i, rec := range response.Items {
		tracks[i] = rec.Track.Name
	}

	return tracks
}

func GetSpotifyPlaylistTracksByUrl(playlistUrl string, oAuthToken string) []string {
	playlistId, isSuccess := ParsePlaylistId(playlistUrl)
	if isSuccess {
		tracks := GetPlaylistTracks(playlistId, oAuthToken)
		return tracks
	}

	fmt.Println("GetSpotifyPlaylistTracksByUrl: Error while parsing playlist id!")
	return nil
}
