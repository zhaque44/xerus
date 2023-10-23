package datapreparation

import (
	"net/url"
	"strings"
)

func SanitizeURL(urlStr string, defaultPort bool) (string, error) {
	URL := strings.TrimSpace(urlStr)

	if strings.HasPrefix(strings.ToLower(URL), "http://") || strings.HasPrefix(strings.ToLower(URL), "https://") {
		if defaultPort {
			u, err := url.Parse(URL)
			if err != nil {
				return "", err
			}
			if u.Port() == "" {
				u.Host += ":80"
			}
			return u.String(), nil
		}
		return URL[strings.Index(URL, "//")+2:], nil
	}

	return URL, nil
}
