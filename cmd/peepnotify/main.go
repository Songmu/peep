package main

import (
	"os"

	"github.com/Songmu/peep/peepnotify"
)

func main() {
	os.Exit(peepnotify.Run(os.Args[1:]))
}
