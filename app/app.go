package app

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"
	"sync"

	"github.com/fatih/color"
)

const version = "0.1.3"

var (
	listEnvs              sync.Map
	sortedKeys            []string
	optVersionInformation *bool
	optColorEnabled       *bool

	colorTitle = color.New(color.Bold, color.FgYellow).SprintFunc()

	usage = `
usage: %[1]s [-flags]

  flags:

  -c, -color          enable colored output
      -version        display version information (%s)

`
)

// CLIApplication represents app structure
type CLIApplication struct {
	Out                  io.Writer
	WorkOnHomeEnvVarName string
}

// NewCLIApplication creates new CLIApplication instance
func NewCLIApplication() *CLIApplication {
	flag.Usage = func() {
		// w/o os.Stdout, you need to pipe out via
		// cmd &> /path/to/file
		fmt.Fprintf(os.Stdout, usage, os.Args[0], version)
		os.Exit(0)
	}

	optVersionInformation = flag.Bool("version", false, "")

	optColorEnabled = flag.Bool("color", false, "enable color")
	flag.BoolVar(optColorEnabled, "c", false, "")

	flag.Parse()

	if !*optColorEnabled {
		color.NoColor = true
	}

	return &CLIApplication{
		Out:                  os.Stdout,
		WorkOnHomeEnvVarName: "WORKON_HOME",
	}
}

// Run executes main application
func (c *CLIApplication) Run() error {
	if *optVersionInformation {
		fmt.Fprintln(c.Out, version)
		return nil
	}

	workonHome, ok := os.LookupEnv(c.WorkOnHomeEnvVarName)
	if !ok {
		return fmt.Errorf("%s environment variable doesn't exists in your environment", c.WorkOnHomeEnvVarName)
	}

	files, err := ioutil.ReadDir(workonHome)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, file := range files {
		if file.IsDir() {
			wg.Add(1)
			go func(dirName string) {
				defer wg.Done()
				pythonBin := workonHome + "/" + dirName + "/bin/python"
				cmd := pythonBin + " --version 2>&1"
				fmt.Printf("cmd: %v\n", cmd)

				pyVersion, err := exec.Command("bash", "-c", cmd).Output()
				if err == nil {
					pyVersion = bytes.TrimSpace(pyVersion)
					listEnvs.Store(dirName, strings.Split(string(pyVersion), " ")[1])
				}
			}(file.Name())
		}
	}
	wg.Wait()

	m := map[string]interface{}{}
	listEnvs.Range(func(key, value interface{}) bool {
		m[key.(string)] = value
		return true
	})

	longestKey := ""
	for key := range m {
		sortedKeys = append(sortedKeys, key)
		if len(key) > len(longestKey) {
			longestKey = key
		}
	}
	sort.Strings(sortedKeys)

	fmt.Fprintf(c.Out, colorTitle("You have %d %s available\n\n"), len(sortedKeys), "environment")

	for _, key := range sortedKeys {
		fmt.Fprintf(c.Out, "[%-*v] %v\n", len(longestKey), key, m[key])
	}

	return nil
}
