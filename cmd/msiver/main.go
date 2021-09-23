package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/zat-kaoru-hayama/go-msidb"
)

func showFile(arg string) error {
	dict, err := msidb.Query(arg)
	if err != nil {
		return err
	}
	fmt.Println(arg)
	for _, key := range []string{"ProductName", "ProductVersion"} {
		fmt.Printf("%s=%s\n", key, dict[key])
	}
	return nil
}

func showDir(arg string) error {
	return filepath.Walk(arg, func(path string, fi fs.FileInfo, _ error) error {
		if strings.EqualFold(filepath.Ext(path), ".msi") && !fi.IsDir() {
			if err := showFile(path); err != nil {
				return err
			}
		}
		return nil
	})
}

func showOne(arg string) error {
	stat, err := os.Stat(arg)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return showDir(arg)
	} else {
		return showFile(arg)
	}
}

func mains(args []string) error {
	clean := msidb.CoInit()
	defer clean()

	for _, arg := range args {
		if err := showOne(arg); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := mains(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
