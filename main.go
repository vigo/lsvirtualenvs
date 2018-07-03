package main

import (
	"fmt"
	"os"

	"github.com/vigo/lsvirtualenvs/app"
)

func main() {
	cmd := app.LsVirtualenvsApp()

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
