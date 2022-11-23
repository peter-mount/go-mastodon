package mastodon

import (
	"path/filepath"
	"strings"
)

type UrlBuilder struct {
	host  string
	path  string
	query []param
}

type param struct {
	key   string
	value string
}

func NewUrlBuilder() *UrlBuilder {
	return &UrlBuilder{}
}

func (b *UrlBuilder) Host(host string) *UrlBuilder {
	b.host = host
	return b
}

func (b *UrlBuilder) Path(path string) *UrlBuilder {
	for strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	if b.path == "" {
		b.path = path
	} else {
		b.path = filepath.Join(b.path, path)
	}
	return b
}

func (b *UrlBuilder) Param(key, value string) *UrlBuilder {
	b.query = append(b.query, param{key: key, value: value})
	return b
}

func (b *UrlBuilder) ParamBool(key string, value bool) *UrlBuilder {
	if value {
		return b.Param(key, "true")
	}
	return b.Param(key, "false")
}

func (b *UrlBuilder) ParamBoolIfTrue(key string, value bool) *UrlBuilder {
	if value {
		return b.ParamBool(key, true)
	}
	return b
}

func (b *UrlBuilder) Build() string {
	var url string

	if b.host != "" {
		url = "https://" + b.host
	}

	url = url + "/" + b.path

	if len(b.query) > 0 {
		url = url + "?"
		for i, param := range b.query {
			if i > 0 {
				url = url + "&"
			}
			// TODO add URL encoding here
			url = url + param.key + "=" + param.value
		}
	}

	return url
}
