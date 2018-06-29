package main

import (
	"bytes"
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

const (
	VERSION      = "0.1.0"
	COLOR_GREEN  = "\x1b[32m"
	COLOR_YELLOW = "\x1b[33m"
	COLOR_WHITE  = "\x1b[37m"
	COLOR_GRAY   = "\x1b[38m"

	COLOR_RESET = "\x1b[0m"
)

var (
	usage = `
Usage: lsvirtualenvs [options...]

List available virtualenvironments created by virtualenvwrapper.

Options:

  -h, --help      Display help! :)
  -c, --color     Enable color output
  -s, --simple    Just list environment names
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
	optionsColorOutput  *bool
	optionsSimpleOutput *bool
	optionsIndexEnabled *bool
	optionsVersion      *bool
	indexDigits         int
)

func colorPrint(color string, text string) string {
	return fmt.Sprintf("%s%s%s", color, text, COLOR_RESET)
}

func rightPadding(text string, length int) string {
	strRepeat := strings.Repeat(".", length-len(text))
	strOut := fmt.Sprintf("[%s] %s", text, strRepeat)

	if *optionsColorOutput {
		strOut = fmt.Sprintf("[%s] %s", colorPrint(COLOR_YELLOW, text), colorPrint(COLOR_GRAY, strRepeat))
	}
	return strOut
}

func calculateIndex(number int) string {
	return fmt.Sprintf("%0*d", indexDigits, number+1)
}

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	optionsColorOutput = flag.Bool("color", false, "")
	flag.BoolVar(optionsColorOutput, "c", false, "")

	optionsSimpleOutput = flag.Bool("simple", false, "")
	flag.BoolVar(optionsSimpleOutput, "s", false, "")

	optionsIndexEnabled = flag.Bool("index", false, "")
	flag.BoolVar(optionsIndexEnabled, "i", false, "")

	optionsVersion = flag.Bool("version", false, "")

	flag.Parse()

	if *optionsVersion {
		fmt.Fprint(os.Stdout, fmt.Sprintf("%s\n", VERSION))
		os.Exit(0)
	}

	workonHome := os.Getenv("WORKON_HOME")
	if workonHome == "" {
		fmt.Fprint(os.Stderr, "$WORKON_HOME does't exists in your environment!")
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(workonHome)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	var lock = sync.RWMutex{}

	result := make(map[string]string)

	for _, file := range files {
		if file.IsDir() == true {
			wg.Add(1)

			go func(dirName string) {
				defer wg.Done()
				lock.Lock()
				defer lock.Unlock()

				pythonBin := fmt.Sprintf("%s/%s/bin/python", workonHome, dirName)
				bashCmd := fmt.Sprintf("%s --version 2>&1", pythonBin)

				pythonVersion, err := exec.Command("bash", "-c", bashCmd).Output()
				if err != nil {
					pythonVersion = []byte("???")
				} else {
					pythonVersion = bytes.TrimSpace(pythonVersion)
				}

				pythonVersionStr := strings.Split(fmt.Sprintf("%s", pythonVersion), " ")[1]
				result[dirName] = pythonVersionStr
			}(file.Name())

		}
	}

	wg.Wait()

	var keys []string
	longestKey := ""
	for k := range result {
		keys = append(keys, k)
		if len(k) > len(longestKey) {
			longestKey = k
		}
	}
	sort.Strings(keys)

	if *optionsSimpleOutput {
		fmt.Println(strings.Join(keys[:], "\n"))
		os.Exit(0)
	}

	envLength := len(keys)
	envAmountStr := "environment"
	if envLength > 1 {
		envAmountStr += "s"
	}

	padOffset := len(longestKey) + 5
	indexDigits = 1
	if envLength > 9 {
		indexDigits = 2
	} else if indexDigits > 99 {
		indexDigits = 3
	} else {
		indexDigits = 8
	}

	if !*optionsSimpleOutput {
		outHeader := "\nYou have %s %s available!\n\n"

		if *optionsColorOutput {
			fmt.Printf(outHeader, colorPrint(COLOR_YELLOW, strconv.Itoa(envLength)), envAmountStr)
		} else {
			fmt.Printf(outHeader, strconv.Itoa(envLength), envAmountStr)
		}
	}

	for index, envName := range keys {
		pythonVersionStr := result[envName]

		if *optionsColorOutput {
			pythonVersionStr = colorPrint(COLOR_WHITE, result[envName])
		}

		if *optionsIndexEnabled {
			needIndex := calculateIndex(index)
			fmt.Printf("%s %v %v\n", needIndex, rightPadding(envName, padOffset), pythonVersionStr)
		} else {
			fmt.Printf("%v %v\n", rightPadding(envName, padOffset), pythonVersionStr)
		}
	}
}
