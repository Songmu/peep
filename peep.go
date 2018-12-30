package peep

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Songmu/wrapcommander"
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
		log.Printf("%s\n", err)
		return exitCodeErr
	}
	return exitCodeOK
}

type peep struct {
	pid         int
	commandArgs []string
	out         io.Writer

	psStat *psStat
}

func (pe *peep) run() error {
	p, err := pe.getPsStat()
	if err != nil {
		return err
	}
	if p == nil {
		return fmt.Errorf("no processes found for pid %d", pe.pid)
	}
	pe.psStat = p

	ret, err := pe.watch()
	if err != nil {
		return err
	}
	return json.NewEncoder(pe.out).Encode(ret)
}

func (pe *peep) getPsStat() (*psStat, error) {
	c := pe.genWatcherCmd()
	out, err := c.Output()
	o := string(out)
	if err != nil {
		exitCode := wrapcommander.ResolveExitCode(err)
		if exitCode != 1 ||
			!strings.Contains(o, "USER") ||
			!strings.Contains(o, "STARTED") ||
			!strings.Contains(o, "COMMAND") {
			return nil, err
		}
	}
	return parsePsStat(o)
}

func (pe *peep) watch() (*result, error) {
	for {
		time.Sleep(time.Second)
		p, err := pe.getPsStat()
		if err != nil {
			return nil, err
		}
		if p == nil || p.Command != pe.psStat.Command {
			return &result{
				psStat: *pe.psStat,
				Ended:  time.Now(),
			}, nil
		}
	}
}

type psStat struct {
	User    string    `json:"user"`
	Command string    `json:"command"`
	Started time.Time `json:"started"`
}

type result struct {
	psStat
	Ended time.Time `json:"ended"`
}

var reg = regexp.MustCompile(`\s+`)

func parsePsStat(str string) (*psStat, error) {
	lines := strings.Split(str, "\n")
	if len(lines) < 2 || lines[1] == "" {
		return nil, nil
	}
	// ex. root     Thu Feb 18 11:12:38 2016 /sbin/init  hoge
	stuff := reg.Split(lines[1], 7)
	if len(stuff) != 7 {
		return nil, fmt.Errorf("invalid ps line: %s", lines[1])
	}
	day, _ := strconv.Atoi(stuff[3])
	dateStr := fmt.Sprintf("%s %s %02d %s %s", stuff[1], stuff[2], day, stuff[4], stuff[5])
	t, err := time.ParseInLocation(time.ANSIC, dateStr, time.Local)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %s", dateStr)
	}
	return &psStat{
		User:    stuff[0],
		Command: stuff[6],
		Started: t,
	}, nil
}

func (pe *peep) genWatcherCmd() *exec.Cmd {
	return exec.Command("ps", "-p", fmt.Sprintf("%d", pe.pid), "-o", "user,lstart,command")
}
