package main

import (
	"fmt"
	"os"

	"github.com/vigo/lsvirtualenvs/app"
)

func main() {
	cmd := app.NewCLIApplication()
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
