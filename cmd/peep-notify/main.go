package main

import (
	"os"

	"github.com/Songmu/peep/peepnotifier"
)

func main() {
	os.Exit(peepnotifier.Run(os.Args[1:]))
}
