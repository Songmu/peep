package peep

import (
	"flag"
	"log"
	"os"
)

const (
	exitCodeOK = iota
	exitCodeErr
)

// Run the peep
func Run(args []string) int {
	err := (&cli{outStream: os.Stdout, errStream: os.Stderr}).run(args)
	if err != nil {
		if err == flag.ErrHelp {
			return exitCodeOK
		}
		log.Printf("[!!ERROR!!] %s\n", err)
		return exitCodeErr
	}
	return exitCodeOK
}

type peep struct {
	pid         int
	commandArgs []string
}

func (pe *peep) run() error {
	return nil
}
