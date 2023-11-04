package main

import (
	"os"

	"github.com/fernandoocampo/clu/internal/application"
)

func main() {
	err := application.Run()
	if err != nil {
		os.Exit(1)
	}
}
