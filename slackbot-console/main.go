package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/monde-sistemas/slackbot"
)

var (
	channel, message string
	stdin            bool
	slackURL         = os.Getenv("SLACKBOT_URL")
)

func main() {
	if strings.TrimSpace(slackURL) == "" {
		fmt.Println("Slackbot URL not found in environment")
		os.Exit(1)
	}
	parseFlags()
	if stdin {
		readFromStdin()
	}
	postMessage()
}

func readFromStdin() {
	input := bytes.NewBuffer([]byte{})
	input.ReadFrom(os.Stdin)

	inputString := input.String()
	if inputString != "" {
		message = fmt.Sprintf("%s\n```%s```", message, inputString)
		message = strings.Replace(message, "\r", "", -1)
	}
}

func parseFlags() {
	flag.StringVar(&channel, "channel", "", "Channel. Ex: #random")
	flag.StringVar(&message, "message", "", "Message to be sent to the channel")
	flag.BoolVar(&stdin, "stdin", false, "Read message from stdin")
	flag.Parse()

	if channel == "" || (message == "" && !stdin) {
		flag.Usage()
		os.Exit(1)
	}
}

func postMessage() {
	b := slackbot.New(slackURL)
	b.PostMessage(channel, message)
}
