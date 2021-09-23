package main

import (
	"archive/zip"
	"fmt"
	"io"
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
	fmt.Printf("%s:\n", filepath.Base(arg))
	for _, key := range []string{"ProductName", "ProductVersion"} {
		fmt.Printf("%s=%s\n", key, dict[key])
	}
	return nil
}

func showZip(arg string) error {
	fd, err := os.Open(arg)
	if err != nil {
		return err
	}
	defer fd.Close()

	stat, err := fd.Stat()
	if err != nil {
		return err
	}

	zipReader, err := zip.NewReader(fd, stat.Size())
	if err != nil {
		return err
	}
	recordSeperator := ""
	for _, file := range zipReader.File {
		if strings.EqualFold(filepath.Ext(file.Name), ".msi") {
			fileReader, err := file.Open()
			if err != nil {
				return err
			}
			defer fileReader.Close()

			tmpMsiPath := filepath.Join(os.TempDir(), filepath.Base(file.Name))
			tmpMsiWriter, err := os.Create(tmpMsiPath)
			if err != nil {
				return err
			}
			io.Copy(tmpMsiWriter, fileReader)
			err = tmpMsiWriter.Close()
			fmt.Print(recordSeperator)
			if err == nil {
				showFile(tmpMsiPath)
			}
			os.Remove(tmpMsiPath)
			if err != nil {
				return err
			}
			recordSeperator = "\n"
		}
	}
	return nil
}

func showDir(arg string) error {
	recordSeparator := ""
	return filepath.Walk(arg, func(path string, fi fs.FileInfo, _ error) error {
		if strings.EqualFold(filepath.Ext(path), ".msi") && !fi.IsDir() {
			fmt.Print(recordSeparator)
			if err := showFile(path); err != nil {
				return err
			}
			recordSeparator = "\n"
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
	} else if strings.EqualFold(filepath.Ext(arg), ".zip") {
		return showZip(arg)
	} else {
		return showFile(arg)
	}
}

func mains(args []string) error {
	clean := msidb.CoInit()
	defer clean()

	recordSeparator := ""
	for _, arg := range args {
		fmt.Print(recordSeparator)
		if err := showOne(arg); err != nil {
			return err
		}
		recordSeparator = "\n"
	}
	return nil
}

func main() {
	if err := mains(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
