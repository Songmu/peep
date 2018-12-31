package peepnotifier

import (
	"flag"
	"io"
	"log"
	"os"
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
	err := (&cli{outStream: os.Stdout, errStream: os.Stderr}).run(args)
	if err != nil {
		if err == flag.ErrHelp {
			return exitCodeOK
		}
		log.Printf("%s\n", err)
		return exitCodeErr
	}
	return exitCodeOK
}
