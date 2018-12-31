package peepnotifier

import (
	"fmt"
	"os"

	slacli "github.com/monochromegane/slack-incoming-webhooks"
)

type slack struct {
}

func (sl *slack) run(re *result, args []string) error {
	u := os.Getenv("SLACK_WEBHOOK_URL")
	if u == "" {
		return fmt.Errorf("please specify environment variable `SLACK_WEBHOOK_URL`")
	}
	cli := slacli.Client{WebhookURL: u}
	msg := fmt.Sprintf(
		"The command %q by %s is finished around %s, which started at %s on %s",
		re.Command, re.User, re.Ended, re.Started, re.Host)
	return cli.Post(&slacli.Payload{
		Username: "peepnotifier",
		Text:     msg,
	})
}
