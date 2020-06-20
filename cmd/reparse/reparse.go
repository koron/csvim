package main

import (
	"log"
	"os"

	"github.com/koron/csvim/internal/colorscheme"
)

func main() {
	cs, err := colorscheme.Read(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	err = cs.Marshal(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
