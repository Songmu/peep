package peepnotifier

import (
	"flag"
	"fmt"
	"os"
	"strings"

	slacli "github.com/monochromegane/slack-incoming-webhooks"
)

type slack struct {
	channel string
}

func (sl *slack) run(re *result, args []string) error {
	fs := flag.NewFlagSet("peepnotify slack", flag.ContinueOnError)
	fs.StringVar(&sl.channel, "c", "", "slack channel")

	err := fs.Parse(args)
	if err != nil {
		return err
	}

	if sl.channel != "" && !strings.HasPrefix(sl.channel, "#") {
		sl.channel = "#" + sl.channel
	}

	u := os.Getenv("SLACK_WEBHOOK_URL")
	if u == "" {
		return fmt.Errorf("please specify environment variable `SLACK_WEBHOOK_URL`")
	}
	cli := slacli.Client{WebhookURL: u}
	msg := fmt.Sprintf(
		"The command %q on %s by %s is finished around %s, which started at %s",
		re.Command, re.Host, re.User, re.Ended, re.Started)
	return cli.Post(&slacli.Payload{
		Username:  "peepnotifier",
		Text:      msg,
		IconEmoji: ":white_flower:",
		Channel:   sl.channel,
	})
}
