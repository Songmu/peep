package peepnotify

import (
	"fmt"
	"os"

	pbcli "github.com/xconstruct/go-pushbullet"
)

type pushbullet struct {
}

func (pu *pushbullet) run(re *result, args []string) error {
	token := os.Getenv("PUSHBULLET_TOKEN")
	if token == "" {
		return fmt.Errorf("no pushbullet token. please specify via environment variable `PUSHBULLET_TOKEN`")
	}
	pb := pbcli.New(token)
	return pb.PushNote("", "command finished!",
		fmt.Sprintf("The command %q by %s is finished at %s, which started at %s on %s", re.Command, re.User, re.Ended, re.Started, re.Host))
}
