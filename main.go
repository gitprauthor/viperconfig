package main

import (
	"log"

	"github.com/gitprauthor/viperconfig/pkg/config"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(c)
}
