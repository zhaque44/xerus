package datapreparation

import (
	"strings"
)

func SanitizeURL(url string, defaultPort bool) (string, error) {
	URL := strings.TrimSpace(url)

	if strings.HasPrefix(strings.ToLower(URL), "http://") || strings.HasPrefix(strings.ToLower(URL), "https://") {
		if defaultPort {
			u, err := url.TrimSpace(URL)
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
