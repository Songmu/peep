package peepnotifier

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	outStream io.Writer = os.Stdout
	errStream io.Writer = os.Stderr
)

const (
	exitCodeOK = iota
	exitCodeErr
)

// Run the peepnotifier
func Run(args []string) int {
	err := (&peepnotifier{outStream: os.Stdout, errStream: os.Stderr}).run(args)
	if err != nil {
		if err == flag.ErrHelp {
			return exitCodeOK
		}
		log.Printf("%s\n", err)
		return exitCodeErr
	}
	return exitCodeOK
}

type peepnotifier struct {
	outStream, errStream io.Writer
}

func (pe *peepnotifier) run(args []string) error {
	outStream = pe.outStream
	errStream = pe.errStream

	log.SetOutput(pe.errStream)
	log.SetPrefix("[peepnotify] ")
	log.SetFlags(0)

	return pe.dispatch(args)
}

var runnerMap = map[string]runner{
	"stdout":     &stdout{},
	"pushbullet": &pushbullet{},
}

func (pe *peepnotifier) dispatch(args []string) error {
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

type runner interface {
	run(*result, []string) error
}

type result struct {
	User    string    `json:"user"`
	Command string    `json:"command"`
	Started time.Time `json:"started"`
	Ended   time.Time `json:"ended"`
	Host    string    `json:"host"`
}
