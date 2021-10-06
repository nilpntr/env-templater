package main

import (
	"log"
)

func main() {
	cmd, err := newRootCmd()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
