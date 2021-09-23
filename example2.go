//go:build run
// +build run

package main

import (
	"fmt"
	"os"

	"github.com/zat-kaoru-hayama/go-msidb"
)

func main() {
	for _, s := range os.Args[1:] {
		installed := msidb.IsInstalled(s)
		if installed {
			fmt.Print("[OK]")
		} else {
			fmt.Print("[NG]")
		}
		fmt.Printf(" %s\n", s)
	}
}
