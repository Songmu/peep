package peep

import (
	"flag"
	"fmt"
	"io"
	"log"
	"runtime"
	"strconv"
)

type cli struct {
	outStream, errStream io.Writer
}

func (cl *cli) run(args []string) error {
	log.SetOutput(cl.errStream)
	log.SetPrefix("[peep] ")
	log.SetFlags(0)

	pe, err := cl.parseArgs(args)
	if err != nil {
		return err
	}
	return pe.run()
}

func (cl *cli) parseArgs(args []string) (*peep, error) {
	pe := &peep{}
	fs := flag.NewFlagSet("peep", flag.ContinueOnError)
	fs.SetOutput(cl.errStream)
	fs.Usage = func() {
		fs.SetOutput(cl.outStream)
		defer fs.SetOutput(cl.errStream)
		fmt.Fprintf(cl.outStream, `peep - Process Peeper

Version: %s (rev: %s/%s)

Synopsis:
    %% peep $PID -- echo 1
Options:
`, version, revision, runtime.Version())
		fs.PrintDefaults()
	}
	fs.StringVar(&pe.host, "H", "", "ssh destination")
	fs.IntVar(&pe.port, "p", 0, "ssh port")

	err := fs.Parse(args)
	if err != nil {
		return nil, err
	}
	pidStr := fs.Arg(0)
	if pe.pid, err = strconv.Atoi(pidStr); err != nil {
		return nil, fmt.Errorf("invalid pid %q", pidStr)
	}
	if fs.NArg() > 2 {
		pe.commandArgs = fs.Args()[1:]
		if pe.commandArgs[0] == "--" {
			pe.commandArgs = pe.commandArgs[1:]
		}
	}
	pe.outStream = cl.outStream
	pe.errStream = cl.errStream
	return pe, nil
}
