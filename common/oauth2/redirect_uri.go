package oauth2

import (
	"errors"
	"fmt"
	"net/url"
)

type redirectURI struct {
	url *url.URL
}

func (uri *redirectURI) String() string {
	return uri.url.String()
}

func (uri *redirectURI) URL() *url.URL {
	return uri.url
}

func ParseRedirectURI(uri string) (URI, error) {
	rurl, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to parse uri -> %v", err)
	}

	if !rurl.IsAbs() {
		return nil, errors.New("redirect uri must be absolute")
	}

	ruri := new(redirectURI)
	ruri.url = rurl

	return ruri, nil
}
