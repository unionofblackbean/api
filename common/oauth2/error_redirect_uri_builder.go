package oauth2

import (
	"errors"
)

type ErrorRedirectURIBuilder struct {
	uri    URI
	err    Error
	dec    string
	errURI string
	state  string
}

func (b *ErrorRedirectURIBuilder) SetRedirectURI(uri URI) {
	b.uri = uri
}

func (b *ErrorRedirectURIBuilder) SetError(err Error) {
	b.err = err
}

func (b *ErrorRedirectURIBuilder) SetErrorDescription(dec string) {
	b.dec = dec
}

func (b *ErrorRedirectURIBuilder) SetErrorURI(uri string) {
	b.errURI = uri
}

func (b *ErrorRedirectURIBuilder) SetState(state string) {
	b.state = state
}

func (b *ErrorRedirectURIBuilder) Build() (string, error) {
	if b.uri == nil {
		return "", errors.New("missing source redirect uri")
	}

	if b.err == "" {
		return "", errors.New("missing error parameter")
	}

	url_ := b.uri.URL()
	query := url_.Query()

	query.Add("error", string(b.err))
	if b.dec != "" {
		query.Add("error_description", b.dec)
	}
	if b.errURI != "" {
		query.Add("error_uri", b.errURI)
	}
	if b.state != "" {
		query.Add("state", b.state)
	}

	return url_.String(), nil
}

func (b *ErrorRedirectURIBuilder) Reset() {
	b.uri = nil
	b.err = ""
	b.dec = ""
	b.errURI = ""
	b.state = ""
}
