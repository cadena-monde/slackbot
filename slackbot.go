package slackbot

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)


type SlackBot struct {
	url string
	token string
}

func New(url, token string) SlackBot {
	b := SlackBot{}
	b.url = url
	b.token = token
	return b
}

func (bot SlackBot) buildURL(channel string) *url.URL {
	u, err := url.Parse(bot.url)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("token", bot.token)
	q.Set("channel", channel)
	u.RawQuery = q.Encode()

	log.Println("URL: ", u.String())
	return u
}

func (bot SlackBot) PostMessage(channel, message string) error {
	r, err := http.Post(bot.buildURL(channel).String(), "text/plain", bytes.NewBufferString(message))
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
