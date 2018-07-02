package main

import (
	"os"
	"testing"
)

func TestPrintColorf(t *testing.T) {
	t.Parallel()
	t.Run("Text should be non-colored", func(t *testing.T) {
		redText := printColorf("Hello", COLOR_RED)
		if redText != "Hello" {
			t.Fail()
		}
	})
	t.Run("Text should colored", func(t *testing.T) {
		*cmdOptionColorOutput = true
		redText := printColorf("Hello", COLOR_RED)
		if redText != "\x1b[31mHello\x1b[0m" {
			t.Fail()
		}
	})
}

func TestEnvironmentVariableChecker(t *testing.T) {
	t.Parallel()
	t.Run("Fake env-var VIGORULEZ should not be exists", func(t *testing.T) {
		_, err := checkEnvVar("VIGORULEZ")
		if err == nil {
			t.Fail()
		}
	})
	t.Run("Fake env-var VIGORULEZ should be equal to OK", func(t *testing.T) {
		os.Setenv("VIGORULEZ", "OK")
		val, _ := checkEnvVar("VIGORULEZ")
		if val != "OK" {
			t.Fail()
		}
	})
}
