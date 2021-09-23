//go:build run
// +build run

package main

import (
	"fmt"
	"os"

	"github.com/zat-kaoru-hayama/go-msidb"
)

func mains() error {
	uninit := msidb.CoInit()
	defer uninit()

	for _, arg1 := range os.Args[1:] {
		db, err := msidb.Query(arg1)
		if err != nil {
			return fmt.Errorf("%s: %w", arg1, err)
		}
		for key, val := range db {
			fmt.Printf("%s=%s\n", key, val)
		}
	}
	return nil
}

func main() {
	if err := mains(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
