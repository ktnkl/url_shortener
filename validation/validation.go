package validation

import (
	"net/url"
)

func IsValidURL(link string) bool {
	parsed, err := url.Parse(link)

	if err != nil || parsed.Host == "" {
		return false
	} else {
		return true
	}
}
