package mastodon

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	GET  = "GET"
	POST = "POST"
)

type RequestParams map[string]string

type Client interface {
	// GetTimeline returns a names specific timeline
	GetTimeline(timeline, minId, maxId string, local, onlyMedia bool, limit int) ([]Status, error)
	// Post posts a status
	Post(status PostStatus) (*Status, error)
	// VerifyCredentials tests that the access token is valid for connecting to the server
	VerifyCredentials() (*Application, error)
}

type client struct {
	config Config
}

func (c *client) IsDebug() bool {
	return c.config.Debug
}

func (c *client) getUri(endpoint string) string {
	return strings.TrimSuffix(c.config.Server, "/") + "/" + strings.TrimPrefix(endpoint, "/")
}

type RequestHandler func(req *http.Request, payload interface{}) error

func (c *client) request(method, endpoint string, reqHandler RequestHandler, params RequestParams, resp interface{}) error {
	var body io.Reader

	contentType := ""
	if params != nil {
		form := make(url.Values)
		for key, val := range params {
			form.Add(key, val)
		}
		body = strings.NewReader(form.Encode())
		contentType = "application/x-www-form-urlencoded"
	}

	req, err := http.NewRequest(method, c.getUri(endpoint), body)
	if err != nil {
		return err
	}

	if reqHandler != nil {
		err = reqHandler(req, nil)
		if err != nil {
			return err
		}
	}

	// Always set this AFTER the RequestHandler is called
	req.Header.Set("Authorization", "Bearer "+c.config.AccessToken)

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	if c.IsDebug() {
		log.Printf("Req %s %s", method, req.URL)
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if c.IsDebug() {
		log.Printf("Status %d %q", response.StatusCode, response.Status)
	}

	if response.StatusCode == 404 {
		return ErrNotFound
	}

	if resp != nil {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return json.Unmarshal(body, resp)
	}

	return nil
}
