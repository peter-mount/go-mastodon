package mastodon

import "log"

type Application struct {
	Name     string  `json:"name"`
	Website  *string `json:"website"`
	VapidKey string  `json:"vapid_key"`
}

// VerifyCredentials tests that the access token is valid for connecting to the server
func (c *client) VerifyCredentials() (*Application, error) {
	app := &Application{}

	err := c.request(GET, "/api/v1/apps/verify_credentials", nil, nil, app)
	if err != nil {
		return nil, err
	}

	if c.IsDebug() {
		website := "nil"
		if app.Website != nil {
			website = *app.Website
		}

		if c.IsDebug() {
			log.Printf("Logged in %q url %v vkey %q", app.Name, website, app.VapidKey)
		}
	}

	return app, nil
}
