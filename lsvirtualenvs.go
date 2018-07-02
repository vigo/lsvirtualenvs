/*

	Build with: go version go1.10.3 darwin/amd64
	Created by Uğur "vigo" Özyılmazel on 2018-07-01.

*/
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var (
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

	cmdOptionColorOutput  *bool
	cmdOptionSimpleOutput *bool
	cmdOptionIndexEnabled *bool
	cmdOptionVersionInfo  *bool
	maxIndexDigits        int

	outputHeader = `You have %s %s available

`
)

const (
	VERSION = "0.1.1"

	COLOR_RED     = "\x1b[31m"
	COLOR_GREEN   = "\x1b[32m"
	COLOR_YELLOW  = "\x1b[33m"
	COLOR_BLUE    = "\x1b[34m"
	COLOR_MAGENTA = "\x1b[35m"
	COLOR_CYAN    = "\x1b[36m"
	COLOR_WHITE   = "\x1b[37m"
	COLOR_GRAY    = "\x1b[38m"

	COLOR_RESET = "\x1b[0m"
)

// init implements commandline flag parsing
func init() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	cmdOptionColorOutput = flag.Bool("color", false, "")
	flag.BoolVar(cmdOptionColorOutput, "c", false, "")

	cmdOptionSimpleOutput = flag.Bool("simple", false, "")
	flag.BoolVar(cmdOptionSimpleOutput, "s", false, "")

	cmdOptionIndexEnabled = flag.Bool("index", false, "")
	flag.BoolVar(cmdOptionIndexEnabled, "i", false, "")

	cmdOptionVersionInfo = flag.Bool("version", false, "")
	flag.Parse()
}

// printColorf implements terminal friendly color output if
// color enable flag is set!
func printColorf(text string, color string) string {
	if *cmdOptionColorOutput {
		return fmt.Sprintf("%s%s%s", color, text, COLOR_RESET)
	} else {
		return text
	}

}

// checkEnvVar checks given environment variable if exists
// returns error unless exists
func checkEnvVar(envVar string) (string, error) {
	v := os.Getenv(envVar)
	if v == "" {
		return "", errors.New("$WORKON_HOME does't exists in your environment!")
	}
	return v, nil
}

// pluralize implements a quick and dirty pluralize process
// if you pass plural version as "s" only, output will be
// suffixed with "s" infront of singular version of text.
func pluralize(singular string, plural string, amount int) string {
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
func dynamicDigitPadding(number int) string {
	return fmt.Sprintf("%0*d", maxIndexDigits, number+1)
}

// rightPaddingWithChar implements padding for given length with given padding
// character.
func rightPaddingWithChar(text string, length int, padChar string, format string) string {
	if padChar == "" {
		padChar = "."
	}

	repeatingChars := strings.Repeat(padChar, length-len(text))
	return fmt.Sprintf(format, text, repeatingChars)
}

func main() {
	if *cmdOptionVersionInfo {
		fmt.Fprint(os.Stdout, fmt.Sprintf("%s\n", VERSION))
		os.Exit(0)
	}

	envWorkonHome, err := checkEnvVar("WORKON_HOME")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	// read all files under envWorkonHome
	filesList, err := ioutil.ReadDir(envWorkonHome)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	var lock = sync.RWMutex{}

	virtualEnvsList := make(map[string]string)

	for _, file := range filesList {
		if file.IsDir() == true {
			wg.Add(1)

			go func(dirName string) {
				defer wg.Done()
				lock.Lock()
				defer lock.Unlock()

				pythonBinPath := fmt.Sprintf("%s/%s/bin/python", envWorkonHome, dirName)
				bashCommand := fmt.Sprintf("%s --version 2>&1", pythonBinPath)
				pythonVersion, err := exec.Command("bash", "-c", bashCommand).Output()

				if err != nil {
					pythonVersion = []byte("???")
				} else {
					pythonVersion = bytes.TrimSpace(pythonVersion)
				}
				virtualEnvsList[dirName] = strings.Split(fmt.Sprintf("%s", pythonVersion), " ")[1]

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
		fmt.Println(strings.Join(keysOfVirtualEnvsList[:], "\n"))
		os.Exit(0)
	}

	fmt.Printf(outputHeader,
		printColorf(strconv.Itoa(lengthOfEnvironments), COLOR_YELLOW),
		fmt.Sprintf(
			"%s",
			pluralize("virtualenv", "s", lengthOfEnvironments)))

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
			indexString = fmt.Sprintf("%s ", printColorf(dynamicDigitPadding(index), COLOR_GREEN))
		}
		appendValue := fmt.Sprintf(
			"%s%s %s",
			indexString,
			rightPaddingWithChar(
				environmentName,
				len(longestKeylength)+5, "",
				fmt.Sprintf("[%s] %s", printColorf("%s", COLOR_YELLOW), printColorf("%s", COLOR_GRAY))),
			printColorf(virtualEnvsList[environmentName], COLOR_WHITE))
		computedResult = append(computedResult, appendValue)
	}

	fmt.Println(strings.Join(computedResult[:], "\n"))
	os.Exit(0)
}
