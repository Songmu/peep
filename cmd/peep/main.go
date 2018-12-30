package main

import (
	"os"

	"github.com/Songmu/peep"
)

func main() {
	os.Exit(peep.Run(os.Args[1:]))
}
