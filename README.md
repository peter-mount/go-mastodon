# Mastodon client in go

This is a very basic mastodon client for use with bots.

I needed something for my [weather station bot](https://area51.social/@me15weather)
and this library is the result.

If a feature is not listed below then please raise an issue as I will be expanding
this library as and when I need features for various projects.

## Features
* Verify AccessToken is valid
* Submit plain text post to mastodon server
* Retrieve a timeline from the server

### To be implemented
* Upload media to attaching to posts (i.e. images)

## Example

    config := mastodon.Config{
      Server: "https://example.com/",
      AccessToken: "tokenfromserver",
    }
    
    client := config.Client()
    
    // Verify AccessToken is valid
    _, err := client.VerifyCredentials()
    if err != nil {
      panic(err)
    }
    
    // Post to the server
    status := PostStatus{Text: "This post is from golang"}
    _, error := client.Post(status)
    if err != nil {
      panic(err)
    }

You can get the access token by logging into your Mastodon account, going into settings,
Development and add a new application which will give you a client key, secret and access token.
You only need the latter for this library to work.
