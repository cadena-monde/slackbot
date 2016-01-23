//Package slackbot allows posting messages to Slack using the Slackbot API
package slackbot

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// SlackBot stores the the current config and methods
type SlackBot struct {
	url string
}

// New configures a new SlackBot instance with the given url
func New(url string) SlackBot {
	return SlackBot{url: url}
}

func (bot SlackBot) buildURL(channel string) (*url.URL, error) {
	u, err := url.Parse(bot.url)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("channel", channel)
	u.RawQuery = q.Encode()

	return u, nil
}

// PostMessage sends the given message to the given channel as Slackbot
func (bot SlackBot) PostMessage(channel, message string) error {
	url, err := bot.buildURL(channel)
	if err != nil {
		log.Printf("Error building the URL: %v", err)
		return err
	}
	r, err := http.Post(url.String(), "text/plain", bytes.NewBufferString(message))
	if err != nil {
		return err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	log.Println(string(b))

	return nil
}
