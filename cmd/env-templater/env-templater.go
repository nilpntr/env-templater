package main

import (
	"log"
	"os"
)

func main() {
	cmd, err := newRootCmd(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
