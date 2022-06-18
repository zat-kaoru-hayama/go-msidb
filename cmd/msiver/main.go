package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zat-kaoru-hayama/go-msidb/internal/msiver"
)

func main() {
	flag.Parse()
	if err := msiver.Show(flag.Args(), os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
