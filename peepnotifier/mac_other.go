// +build !darwin

package peepnotifier

import "fmt"

func (ma *mac) run(re *result, args []string) error {
	return fmt.Errorf("mac notification is not supported on non-mac environment")
}
