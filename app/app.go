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
	"strconv"
	"strings"
	"sync"
)

const (
	appVersion   = "2.1.0"
	colorRed     = "\x1b[31m"
	colorGreen   = "\x1b[32m"
	colorYellow  = "\x1b[33m"
	colorBlue    = "\x1b[34m"
	colorMagenta = "\x1b[35m"
	colorCyan    = "\x1b[36m"
	colorWhite   = "\x1b[37m"
	colorGray    = "\x1b[38m"
	colorReset   = "\x1b[0m"
)

var (
	cmdOptionVersionInfo  *bool
	cmdOptionSimpleOutput *bool
	cmdOptionColorOutput  *bool
	cmdOptionIndexEnabled *bool
	maxIndexDigits        int
	outputHeader          = `You have %s %s available

`
	usage = `
Usage: lsvirtualenvs [options...]

List available virtualenvironments created by virtualenvwrapper.

Options:

  -h, --help      Display help! :)
  -c, --color     Enable color output
  -s, --simple    Just list environment names, overrides -c, -i
  -i, --index     Add index number to output
      --version   Version information

Examples:

  lsvirtualenvs -h
  lsvirtualenvs --version
  lsvirtualenvs -c
  lsvirtualenvs --color
  lsvirtualenvs -c -i
  lsvirtualenvs --color --index

  lsvirtualenvs -s
  lsvirtualenvs -simple

`
)

type CmdApp struct {
	out io.Writer
}

// LsVirtualenvsApp is the core application
func LsVirtualenvsApp() *CmdApp {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}
	cmdOptionVersionInfo = flag.Bool("version", false, "Current version information")

	cmdOptionSimpleOutput = flag.Bool("simple", false, "Just list environment names, overrides -c, -i")
	flag.BoolVar(cmdOptionSimpleOutput, "s", false, "")

	cmdOptionColorOutput = flag.Bool("color", false, "Enable color output")
	flag.BoolVar(cmdOptionColorOutput, "c", false, "")

	cmdOptionIndexEnabled = flag.Bool("index", false, "Add index number to output")
	flag.BoolVar(cmdOptionIndexEnabled, "i", false, "")

	flag.Parse()

	return &CmdApp{
		out: os.Stdout,
	}
}

// Version returns the current version of LsVirtualenvsApp
func (m *CmdApp) Version() error {
	fmt.Fprintf(m.out, "%s\n", appVersion)
	return nil
}

// Run implements LsVirtualenvsApp to run!
func (m *CmdApp) Run() error {
	if *cmdOptionVersionInfo {
		return m.Version()
	}
	return m.GetVirtualenvs()
}

// printColorf implements terminal friendly color output if
// color enable flag is set!
func (m *CmdApp) PrintColorf(text string, color string) string {
	if *cmdOptionColorOutput {
		return fmt.Sprintf("%s%s%s", color, text, colorReset)
	} else {
		return text
	}
}

// pluralize implements a quick and dirty pluralize process
// if you pass plural version as "s" only, output will be
// suffixed with "s" infront of singular version of text.
func (m *CmdApp) Pluralize(singular string, plural string, amount int) string {
	out := singular
	if amount > 1 {
		if plural == "s" {
			out = fmt.Sprintf("%ss", singular)
		} else {
			out = plural
		}
	}
	return out
}

// dynamicDigitPadding implements dynamic string format with
// digit padding
func (m *CmdApp) DynamicDigitPadding(number int) string {
	return fmt.Sprintf("%0*d", maxIndexDigits, number+1)
}

// rightPaddingWithChar implements padding for given length with given padding
// character. If padChar is not provided, function uses "." as default.
func (m *CmdApp) RightPaddingWithChar(text string, length int, padChar string, format string) string {
	if padChar == "" {
		padChar = "."
	}
	repeatingChars := strings.Repeat(padChar, length-len(text))
	return fmt.Sprintf(format, text, repeatingChars)
}

// GetVirtualenvs implements fetching virtualenv names and related
// Python versions. Checks "WORKON_HOME" environment variable first,
// the checks if the required folder exists.
func (m *CmdApp) GetVirtualenvs() error {
	lookup := "WORKON_HOME"
	currentWorkingDir, envExists := os.LookupEnv(lookup)
	if !envExists {
		return errors.New(fmt.Sprintf("%s doesn't exists in your environment!", lookup))
	}

	filesList, err := ioutil.ReadDir(currentWorkingDir)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	var lock = sync.RWMutex{}

	virtualEnvsList := make(map[string]string)

	for _, file := range filesList {
		if file.IsDir() == true {
			wg.Add(1)

			go func(dirName string) {
				defer wg.Done()

				pythonBinPath := fmt.Sprintf("%s/%s/bin/python", currentWorkingDir, dirName)
				bashCommand := fmt.Sprintf("%s --version 2>&1", pythonBinPath)

				pythonVersion, err := exec.Command("bash", "-c", bashCommand).Output()

				if err != nil {
					pythonVersion = []byte("???")
				} else {
					pythonVersion = bytes.TrimSpace(pythonVersion)
				}

				lock.Lock()
				virtualEnvsList[dirName] = strings.Split(fmt.Sprintf("%s", pythonVersion), " ")[1]
				defer lock.Unlock()
			}(file.Name())
		}
	}
	wg.Wait()

	var keysOfVirtualEnvsList []string

	longestKeylength := ""
	for key := range virtualEnvsList {
		keysOfVirtualEnvsList = append(keysOfVirtualEnvsList, key)
		if len(key) > len(longestKeylength) {
			longestKeylength = key
		}
	}
	sort.Strings(keysOfVirtualEnvsList)
	lengthOfEnvironments := len(keysOfVirtualEnvsList)

	if *cmdOptionSimpleOutput {
		fmt.Fprintf(m.out, "%s", strings.Join(keysOfVirtualEnvsList[:], "\n"))
		return nil
	}

	fmt.Fprintf(
		m.out,
		outputHeader,
		m.PrintColorf(strconv.Itoa(lengthOfEnvironments), colorYellow),
		fmt.Sprintf("%s", m.Pluralize("virtualenv", "s", lengthOfEnvironments)))

	var computedResult []string

	if *cmdOptionIndexEnabled {
		maxIndexDigits = func(number int) int {
			if number < 9 {
				return 1
			} else if number > 9 {
				return 2
			} else if number > 99 {
				return 3
			} else {
				return 6
			}
		}(lengthOfEnvironments)
	}

	for index, environmentName := range keysOfVirtualEnvsList {
		indexString := ""
		if *cmdOptionIndexEnabled {
			indexString = fmt.Sprintf("%s ", m.PrintColorf(m.DynamicDigitPadding(index), colorGreen))
		}
		appendValue := fmt.Sprintf(
			"%s%s %s",
			indexString,
			m.RightPaddingWithChar(
				environmentName,
				len(longestKeylength)+5, "",
				fmt.Sprintf(
					"[%s] %s",
					m.PrintColorf("%s", colorYellow),
					m.PrintColorf("%s", colorGray))),
			m.PrintColorf(virtualEnvsList[environmentName], colorWhite))
		computedResult = append(computedResult, appendValue)
	}

	fmt.Fprintf(m.out, fmt.Sprintf("%s\n", strings.Join(computedResult[:], "\n")))
	return nil
}
