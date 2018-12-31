package peepnotifier

import "time"

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
