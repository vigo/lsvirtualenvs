/*

	Build with: go version go1.11.1 darwin/amd64
	Created by Uğur "vigo" Özyılmazel on 2018-07-01.

*/

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
