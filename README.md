# Slackbot

A Slack notifier library for Go applications.

##Usage

Create a new instance of SlackBot specifying the URL of Slack and the access token, then call the method "PostMessage" passing the channel and message.

Example:

```
func SendSlackNotification() {
	b := slackbot.New("https://my_slack_hook_url", "my_slack_token")
	b.PostMessage("#random", "All your base ")
}
```