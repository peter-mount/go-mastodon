package mastodon

type Config struct {
	Server      string `json:"server" xml:"server,attr" yaml:"server"`                   // The Mastodon server to connect to
	AccessToken string `json:"access_token" xml:"access_token,attr" yaml:"access_token"` // The AccessToken
	Debug       bool   `json:"debug" xml:"debug,attr" yaml:"debug"`                      // true to enable debugging
}

func (c Config) Client() Client {
	return &client{config: c}
}
