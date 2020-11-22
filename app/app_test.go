package app

import (
	"bytes"
	"testing"
)

func TestCLIApplication(t *testing.T) {
	cmd := NewCLIApplication()
	buff := new(bytes.Buffer)
	cmd.Out = buff

	t.Run("app should response to version information request", func(t *testing.T) {
		*optVersionInformation = true

		if err := cmd.Run(); err != nil {
			t.Errorf("want: nil, got: %v", err)
		}
		if versionInfo := string(bytes.TrimSpace(buff.Bytes())); versionInfo != version {
			t.Errorf("want: %s, got: %s", version, versionInfo)
		}
	})

	t.Run("fake lookup envname should raise an error", func(t *testing.T) {
		*optVersionInformation = false
		cmd.WorkOnHomeEnvVarName = "THISISFAKEWORKONHOMEENV"
		if err := cmd.Run(); err == nil {
			t.Errorf("want: error, got: %v", err)
		}
	})
}

// func TestColorPrint(t *testing.T) {
// 	t.Parallel()
// 	t.Run("Text should be non-colored", func(t *testing.T) {
// 		want := "Hello"
// 		got := cmd.PrintColorf("Hello", colorRed)
// 		if want != got {
// 			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
// 		}
// 	})
// 	t.Run("Text should be red", func(t *testing.T) {
// 		*cmdOptionColorOutput = true
//
// 		want := "\x1b[31mHello\x1b[0m"
// 		got := cmd.PrintColorf("Hello", colorRed)
//
// 		if want != got {
// 			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
// 		}
// 	})
// }
//
// func TestPluralize(t *testing.T) {
// 	t.Parallel()
// 	t.Run("1 file", func(t *testing.T) {
// 		want := "file"
// 		got := cmd.Pluralize("file", "s", 1)
// 		if want != got {
// 			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
// 		}
// 	})
// 	t.Run("2 files", func(t *testing.T) {
// 		want := "files"
// 		got := cmd.Pluralize("file", "s", 2)
// 		if want != got {
// 			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
// 		}
// 	})
// 	t.Run("person should become people", func(t *testing.T) {
// 		want := "people"
// 		got := cmd.Pluralize("person", "people", 10)
// 		if want != got {
// 			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
// 		}
// 	})
// 	t.Run("zero should not change to result", func(t *testing.T) {
// 		want := "burger"
// 		got := cmd.Pluralize("burger", "s", 0)
// 		if want != got {
// 			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
// 		}
// 	})
// }
//
// func TestRightPadding(t *testing.T) {
// 	t.Parallel()
// 	t.Run("foo should return with space and 5 dots", func(t *testing.T) {
// 		want := "foo ....."
// 		got := cmd.RightPaddingWithChar("foo", len("foo")+5, "", "%s %s")
// 		if want != got {
// 			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
// 		}
// 	})
// 	t.Run("foo should with 10 stars", func(t *testing.T) {
// 		want := "foo**********"
// 		got := cmd.RightPaddingWithChar("foo", len("foo")+10, "*", "%s%s")
// 		if want != got {
// 			t.Errorf("\nwant: %s\n got: %s\n\n", want, got)
// 		}
// 	})
// }
