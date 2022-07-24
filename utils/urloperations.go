package utils

import (
	"regexp"
)

func IsAnUrl(url string) bool {
	urlRegex := regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
	if len(urlRegex.FindStringIndex(url)) != 0 {
		return true
	} else {
		return false
	}
}

func SCUrlToId(url string) string {
	SCRegex := regexp.MustCompile(`https?:\/\/(www\.)?soundcloud.com\/(.*)`)
	return SCRegex.FindStringSubmatch(url)[2]
}
