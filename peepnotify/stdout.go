package peepnotify

import (
	"encoding/json"
)

type stdout struct {
}

func (st *stdout) run(re *result, args []string) error {
	return json.NewEncoder(outStream).Encode(re)
}
