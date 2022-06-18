package main

import (
	"strings"

	"github.com/mattn/msgbox"
	"github.com/sqweek/dialog"
	"github.com/zat-kaoru-hayama/go-msidb/internal/msiver"
)

func show() string {
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

func main() {
	msgbox.Show(0, show(), "gMsiver", msgbox.OK)
}
