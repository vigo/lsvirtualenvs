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
		*optColorEnabled = false

		if err := cmd.Run(); err != nil {
			t.Errorf("want: nil, got: %v", err)
		}
		if versionInfo := string(bytes.TrimSpace(buff.Bytes())); versionInfo != version {
			t.Errorf("want: %s, got: %s", version, versionInfo)
		}
	})
}
