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
	pid                  int
	commandArgs          []string
	outStream, errStream io.Writer

	host string
	port int

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
		// XXX notify something?
		return err
	}
	if len(pe.commandArgs) > 0 {
		cmd := exec.Command(pe.commandArgs[0], pe.commandArgs[1:]...)
		cmd.Stdout = pe.outStream
		cmd.Stderr = pe.errStream
		err := func() error {
			stdin, err := cmd.StdinPipe()
			if err != nil {
				return err
			}
			defer stdin.Close()
			return json.NewEncoder(stdin).Encode(ret)
		}()
		if err != nil {
			return err
		}
		return cmd.Run()
	}
	return json.NewEncoder(pe.outStream).Encode(ret)
}

func (pe *peep) getPsStat() (*psStat, error) {
	ps := []string{"ps", "-p", fmt.Sprintf("%d", pe.pid), "-o", "user,lstart,command"}
	if pe.host != "" {
		ssh := []string{"ssh", pe.host}
		if pe.port > 0 {
			ssh = append(ssh, "-p", fmt.Sprintf("%d", pe.port))
		}
		ps = append(ssh, ps...)
	}
	c := exec.Command(ps[0], ps[1:]...)
	c.Stdin = os.Stdin
	out, err := c.Output()
	o := string(out)
	// XXX recover ssh error?
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
	interval := time.Second
	if pe.host != "" {
		interval *= 5
	}
	for {
		time.Sleep(interval)
		p, err := pe.getPsStat()
		if err != nil {
			return nil, err
		}
		if p == nil || !p.StartAt.Equal(pe.psStat.StartAt) {
			h := pe.host
			if h == "" {
				h = "localhost" // XXX retrieve from `hostname` command?
			}
			return &result{
				psStat: *pe.psStat,
				EndAt:  time.Now().Truncate(time.Second),
				Host:   h,
				Pid:    pe.pid,
			}, nil
		}
	}
}

type psStat struct {
	User    string    `json:"user"`
	Command string    `json:"command"`
	StartAt time.Time `json:"startAt"`
}

type result struct {
	psStat
	EndAt time.Time `json:"endAt"`
	Host  string    `json:"host"`
	Pid   int       `json:"pid"`
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
		StartAt: t,
	}, nil
}
