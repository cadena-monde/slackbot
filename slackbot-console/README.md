# Slackbot-Console

Go console application to send messages to a Slack channel using command line params.

## Configuration

The application uses the following environment variables:

- SLACKBOT_URL - Ex: https://my_account_name.slack.com/services/hooks/slackbot.
- SLACKBOT_TOKEN - Slack generated token.

## Usage

Call the application passing the channel and the message through command line params:

```
    slackbot-console.exe -channel="#random" -message="All your base are belong to us!"
```

### Reading the message from stdin

You can also pass a message from the stding, example:

    type file.txt | slackbot-console.exe -channel="#random" -message="Title"

In this case, the contents from the stdin are enclosed by ``` to be displayed as pre formatted text on slack.
