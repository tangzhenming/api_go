package main

import (
	"log"

	"github.com/tang-projects/api_go/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
