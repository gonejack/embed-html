package main

import (
	"log"

	"github.com/gonejack/embed-html/embedhtml"
)

func main() {
	cmd := embedhtml.EmbedHTML{
		Options: embedhtml.MustParseOption(),
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
