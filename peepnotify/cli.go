package peepnotify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type cli struct {
	outStream, errStream io.Writer
}

func (cl *cli) run(args []string) error {
	outStream = cl.outStream
	errStream = cl.errStream

	log.SetOutput(cl.errStream)
	log.SetPrefix("[peepnotify] ")
	log.SetFlags(0)

	return cl.dispatch(args)
}

var runnerMap = map[string]runner{
	"stdout": &stdout{},
}

func (cl *cli) dispatch(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no subcommands specified")
	}
	ru, ok := runnerMap[args[0]]
	if !ok {
		return fmt.Errorf("unknown subcommand: %s", args[0])
	}
	re := &result{}
	err := json.NewDecoder(os.Stdin).Decode(re)
	if err != nil {
		return err
	}
	return ru.run(re, args[1:])
}
