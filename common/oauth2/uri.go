package oauth2

import "net/url"

type URI interface {
	String() string
	URL() *url.URL
}
