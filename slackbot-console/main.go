package main

import (
	"github.com/cadena-monde/slackbot"
	"os"
	"flag"
)

var (
	channel, message string
	slackURL        = os.Getenv("SLACKBOT_URL")
	token           = os.Getenv("SLACKBOT_TOKEN")
)

func main() {
	parseFlags()
	postMessage();
}

func parseFlags() {
	flag.StringVar(&channel, "channel", "", "Channel. Ex: #random")
	flag.StringVar(&message, "message", "", "Message to be sent to the channel")
	flag.Parse()

	if channel == "" || message == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func postMessage() {
	b := slackbot.New(slackURL, token)
	b.PostMessage(channel, message)
}
