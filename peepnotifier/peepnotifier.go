package peepnotifier

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
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
	log.SetPrefix("[peep-notify] ")
	log.SetFlags(0)

	return pe.dispatch(args)
}

var runnerMap = map[string]runner{
	"stdout":     &stdout{},
	"pushbullet": &pushbullet{},
	"slack":      &slack{},
	"mac":        &mac{},
}

var commands []string

func init() {
	for k := range runnerMap {
		commands = append(commands, k)
	}
	sort.Strings(commands)
}

func (pe *peepnotifier) dispatch(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no subcommands specified")
	}
	if strings.HasPrefix(args[0], "-") {
		fs := flag.NewFlagSet("peep-notify", flag.ContinueOnError)
		fs.Usage = func() {
			fmt.Fprintf(pe.outStream, `Usage: peep-notify <target> [<args>]

Following targets are available:
  %s`, strings.Join(commands, "\n  "))
		}
		if err := fs.Parse(args); err != nil {
			return err
		}
		args = fs.Args()
	}

	ru, ok := runnerMap[args[0]]
	if !ok {
		return fmt.Errorf("unknown subcommand: %s", args[0])
	}
	var re *result
	if !terminal.IsTerminal(syscall.Stdin) {
		re = &result{}
		err := json.NewDecoder(os.Stdin).Decode(re)
		if err != nil {
			return err
		}
	}
	return ru.run(re, args[1:])
}

type runner interface {
	run(*result, []string) error
}

type result struct {
	User    string    `json:"user"`
	Command string    `json:"command"`
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
	Host    string    `json:"host"`
}
