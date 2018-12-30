package peep

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
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

	psStat *psStat
}

func (pe *peep) run() error {
	c := pe.genWatcherCmd()
	out, err := c.Output()
	if err != nil {
		return err
	}
	_ = parsePsStat(string(out))

	return nil
}

type psStat struct {
	user    string
	command string
	started time.Time
}

var reg = regexp.MustCompile(`\s+`)

func parsePsStat(str string) *psStat {
	lines := strings.Split(str, "\n")
	if len(lines) < 2 {
		return nil
	}
	// ex. root     Thu Feb 18 11:12:38 2016 /sbin/init  hoge
	stuff := reg.Split(lines[1], 7)
	if len(stuff) != 7 {
		return nil
	}
	day, _ := strconv.Atoi(stuff[3])
	dateStr := fmt.Sprintf("%s %s %02d %s %s", stuff[1], stuff[2], day, stuff[4], stuff[5])
	t, err := time.ParseInLocation(time.ANSIC, dateStr, time.Local)
	if err != nil {
		return nil
	}
	return &psStat{
		user:    stuff[0],
		command: stuff[6],
		started: t,
	}
}

func (pe *peep) genWatcherCmd() *exec.Cmd {
	return exec.Command("ps", "-p", fmt.Sprintf("%d", pe.pid), "-o", "user,lstart,command")
}
