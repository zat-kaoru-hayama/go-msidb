package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/zat-kaoru-hayama/go-msidb/internal/msiver"
)

var version string

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		fmt.Fprintf(os.Stderr, "%s %s for %s/%s by %s\n",
			os.Args[0], version, runtime.GOOS, runtime.GOARCH, runtime.Version())
		flag.PrintDefaults()
		return
	}
	if err := msiver.Show(args, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
