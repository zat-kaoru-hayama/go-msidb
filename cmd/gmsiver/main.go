package main

import (
	"os"
	"strings"

	"github.com/mattn/msgbox"
	"github.com/sqweek/dialog"
	"github.com/zat-kaoru-hayama/go-msidb/internal/msiver"
)

func gShow() string {
	target, err := dialog.File().Title("MSI,ZIP or folder").Load()
	if err != nil {
		return err.Error()
	}
	var buffer strings.Builder
	if err := msiver.Show([]string{target}, &buffer); err != nil {
		return err.Error()
	}
	return buffer.String()
}

func show(args []string) string {
	var buffer strings.Builder
	if err := msiver.Show(args, &buffer); err != nil {
		return err.Error()
	}
	return buffer.String()
}

func main() {
	var output string
	if len(os.Args) >= 2 {
		output = show(os.Args[1:])
	} else {
		output = gShow()
	}
	msgbox.Show(0, output, "gMsiver", msgbox.OK)
}
