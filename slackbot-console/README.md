# Slackbot-Console

Go console application to send messages to a Slack channel using command line params.

## Configuration

The application uses the `SLACKBOT_URL` environment variable, ex:

```
export SLACKBOT_URL='https://my_account_name.slack.com/services/hooks/slackbot?token=<your_token>'
```

## Usage

Call the application passing the channel and the message through command line params:

```
slackbot-console -channel="#random" -message="All your base are belong to us!"
```

### Reading the message from stdin

You can also pass a message from the stding, example:

Windows:
```
type file.txt | slackbot-console.exe -channel="#random" -message="Title"
```
Linux:
```
cat file.txt | slackbot-console -channel="#random" -message="Title"
```

In this case, the contents from the stdin are enclosed by ``` to be displayed as pre formatted text on slack.

Tip: You can also mix `-message` with reading a file from stdin.
