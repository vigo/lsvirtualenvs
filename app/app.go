package app

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const version = "0.1.0"

var (
	errWorkonHomeEnvNotExists = errors.New("you need to set WORKON_HOME environment variable")

	listEnvs              map[string]string
	sortedKeys            []string
	optVersionInformation *bool
	optColorEnabled       *bool
	optSimpleOutput       *bool
	optIndexEnabled       *bool

	colorTitle        = color.New(color.Bold, color.FgYellow).SprintFunc()
	colorDots         = color.New(color.Faint).SprintFunc()
	colorEnvName      = color.New(color.Underline, color.FgGreen).SprintFunc()
	colorPyVersion    = color.New(color.FgWhite).SprintFunc()
	colorIndexNumbers = color.New(color.FgHiBlue).SprintFunc()

	usage = `
usage: %[1]s [-flags]

lists existing virtualenvs which are created via "mkvirtualenv" command.

  flags:

  -c, -color          enable colored output
  -s, -simple         just list environment names, overrides -c, -i
  -i, -index          add index number to output
      -version        display version information (%s)

`
)

// CLIApplication represents app structure.
type CLIApplication struct {
	Out io.Writer
}

// NewCLIApplication creates new CLIApplication instance.
func NewCLIApplication() *CLIApplication {
	flag.Usage = func() {
		// w/o os.Stdout, you need to pipe out via
		// cmd &> /path/to/file
		fmt.Fprintf(os.Stdout, usage, os.Args[0], version)
		os.Exit(0)
	}

	optVersionInformation = flag.Bool("version", false, "")

	optColorEnabled = flag.Bool("color", false, "enable colored output")
	flag.BoolVar(optColorEnabled, "c", false, "")

	optSimpleOutput = flag.Bool("simple", false, "just list environment names, overrides -c")
	flag.BoolVar(optSimpleOutput, "s", false, "")

	optIndexEnabled = flag.Bool("index", false, "add index number to output")
	flag.BoolVar(optIndexEnabled, "i", false, "")

	flag.Parse()

	if !*optColorEnabled {
		color.NoColor = true
	}

	if !(os.Getenv("LSVIRTUALENVS_COLOR_ALWAYS") == "") {
		color.NoColor = false
	}

	return &CLIApplication{
		Out: os.Stdout,
	}
}

// Run executes main application.
func (c *CLIApplication) Run() error {
	if *optVersionInformation {
		fmt.Fprintln(c.Out, version)
		return nil
	}

	workonHome, ok := os.LookupEnv("WORKON_HOME")
	if !ok {
		return errWorkonHomeEnvNotExists
	}

	files, err := ioutil.ReadDir(workonHome)
	if err != nil {
		return fmt.Errorf("read dir error %w", err)
	}

	listEnvs = make(map[string]string)
	for _, file := range files {
		if file.IsDir() {
			c := make(chan []string)
			go func(dirName string, c chan []string) {
				pythonBin := workonHome + "/" + dirName + "/bin/python"
				cmd := pythonBin + " --version 2>&1"

				pyVersion, err := exec.Command("bash", "-c", cmd).Output()
				if err == nil {
					pyVersion = bytes.TrimSpace(pyVersion)
					c <- []string{dirName, strings.Split(string(pyVersion), " ")[1]}
				}
			}(file.Name(), c)

			result := <-c
			listEnvs[result[0]] = result[1]
		}
	}

	longestKey := ""
	for key := range listEnvs {
		sortedKeys = append(sortedKeys, key)
		if len(key) > len(longestKey) {
			longestKey = key
		}
	}
	sort.Strings(sortedKeys)

	_ = message.Set(
		language.English,
		"you have %d environment available",
		plural.Selectf(1, "%d",
			"=1", "you have one environment available",
			"=2", "you have %[1]d environments available",
			"other", "you have %[1]d environments available",
		))
	p := message.NewPrinter(language.English)
	titleMessage := p.Sprintf("you have %d environment available", len(sortedKeys))

	if *optSimpleOutput {
		fmt.Fprintf(c.Out, "%s", strings.Join(sortedKeys, "\n"))
		fmt.Println()
		return nil
	}

	fmt.Fprintf(c.Out, "%s\n\n", colorTitle(titleMessage))
	for index, key := range sortedKeys {
		strIndex := ""
		if *optIndexEnabled {
			strIndex = colorIndexNumbers(string(fmt.Sprintf("[%04d] ", index+1)))
		}
		fmt.Fprintf(
			c.Out,
			"%s%s%s %v\n",
			strIndex,
			colorEnvName(key),
			colorDots(strings.Repeat(".", (len(longestKey)+5)-len(key))),
			colorPyVersion(listEnvs[key]),
		)
	}
	fmt.Println()
	return nil
}
