// cmd/pulsedev/main.go
package main

import (
	"flag"
	"log"
	"os"

	"pulsedev/internal/pulsedev"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := pulsedev.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
