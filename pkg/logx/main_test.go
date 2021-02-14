package logx

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_, _ = New(&Config{
		Debug:    true,
		AppName:  "gdk",
		Filename: "",
	})

	os.Exit(m.Run())
}
