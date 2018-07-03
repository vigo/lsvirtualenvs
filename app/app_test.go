package app

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

var cmd *CmdApp

func init() {
	cmd = LsVirtualenvsApp()
}

func TestAppVersion(t *testing.T) {
	t.Run("App should have a version information", func(t *testing.T) {
		buff := new(bytes.Buffer)

		*cmdOptionVersionInfo = true
		cmd.out = buff
		cmd.Run()

		want := "2.0.0"
		got := fmt.Sprintf("%s", bytes.TrimSpace(buff.Bytes()))
		if want != got {
			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
		}
	})
}

func TestColorPrint(t *testing.T) {
	t.Parallel()
	t.Run("Text should be non-colored", func(t *testing.T) {
		want := "Hello"
		got := cmd.PrintColorf("Hello", colorRed)
		if want != got {
			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
		}
	})
	t.Run("Text should be red", func(t *testing.T) {
		*cmdOptionColorOutput = true

		want := "\x1b[31mHello\x1b[0m"
		got := cmd.PrintColorf("Hello", colorRed)

		if want != got {
			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
		}
	})
}

func TestGetEnvOrDie(t *testing.T) {
	t.Parallel()
	t.Run("FAKE_FOO_BAR env-var raise an error", func(t *testing.T) {
		want := "FAKE_FOO_BAR doesn't exists in your environment!"
		_, err := cmd.GetEnvOrDie("FAKE_FOO_BAR")
		got := fmt.Sprintf("%s", err)
		if want != got {
			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
		}
	})
	t.Run("FAKEISH_FOO_BAR env-var should return: BAZ", func(t *testing.T) {
		os.Setenv("FAKEISH_FOO_BAR", "BAZ")

		want := "BAZ"
		got, _ := cmd.GetEnvOrDie("FAKEISH_FOO_BAR")
		if want != got {
			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
		}
	})
}

func TestPluralize(t *testing.T) {
	t.Parallel()
	t.Run("1 file", func(t *testing.T) {
		want := "file"
		got := cmd.Pluralize("file", "s", 1)
		if want != got {
			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
		}
	})
	t.Run("2 files", func(t *testing.T) {
		want := "files"
		got := cmd.Pluralize("file", "s", 2)
		if want != got {
			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
		}
	})
	t.Run("person should become people", func(t *testing.T) {
		want := "people"
		got := cmd.Pluralize("person", "people", 10)
		if want != got {
			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
		}
	})
	t.Run("zero should not change to result", func(t *testing.T) {
		want := "burger"
		got := cmd.Pluralize("burger", "s", 0)
		if want != got {
			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
		}
	})
}

func TestRightPadding(t *testing.T) {
	t.Parallel()
	t.Run("foo should return with space and 5 dots", func(t *testing.T) {
		want := "foo ....."
		got := cmd.RightPaddingWithChar("foo", len("foo")+5, "", "%s %s")
		if want != got {
			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
		}
	})
	t.Run("foo should with 10 stars", func(t *testing.T) {
		want := "foo**********"
		got := cmd.RightPaddingWithChar("foo", len("foo")+10, "*", "%s%s")
		if want != got {
			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
		}
	})
}
