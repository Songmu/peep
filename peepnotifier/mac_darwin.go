package peepnotifier

import (
	"fmt"

	gosxnotifier "github.com/deckarep/gosx-notifier"
)

func (ma *mac) run(re *result, args []string) error {
	if re == nil {
		return fmt.Errorf("please accept result json via stdin")
	}

	msg := fmt.Sprintf(
		"The command `%s` on %s by %s is finished around %s, which started at %s",
		re.Command, re.Host, re.User, re.Ended, re.Started)
	noti := gosxnotifier.NewNotification(msg)
	noti.Title = "Command Finished!"
	return noti.Push()
}
