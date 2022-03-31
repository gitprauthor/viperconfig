package main

import (
	"log"

	"github.com/gitprauthor/viperconfig/pkg/config"
)

func main() {
	_, err := config.DefaultConfig()
	if err != nil {
		log.Fatal(err)
	}
}
