package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/zat-kaoru-hayama/go-msidb"
)

var flagAll = flag.Bool("a", false, "show all values")

func showFile(arg string, w io.Writer) error {
	dict, err := msidb.Query(arg)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "%s:\n", filepath.Base(arg))
	if *flagAll {
		for key, val := range dict {
			fmt.Printf("%s=%s\n", key, val)
		}
	} else {
		for _, key := range []string{"ProductName", "ProductVersion"} {
			fmt.Fprintf(w, "%s=%s\n", key, dict[key])
		}
	}
	return nil
}

func showZip(arg string, w io.Writer) error {
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
				showFile(tmpMsiPath, w)
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

func showDir(arg string, w io.Writer) error {
	recordSeparator := ""
	return filepath.Walk(arg, func(path string, fi fs.FileInfo, _ error) error {
		if strings.EqualFold(filepath.Ext(path), ".msi") && !fi.IsDir() {
			fmt.Fprint(w, recordSeparator)
			if err := showFile(path, w); err != nil {
				return err
			}
			recordSeparator = "\n"
		}
		return nil
	})
}

func showOne(arg string, w io.Writer) error {
	stat, err := os.Stat(arg)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return showDir(arg, w)
	} else if strings.EqualFold(filepath.Ext(arg), ".zip") {
		return showZip(arg, w)
	} else {
		return showFile(arg, w)
	}
}

func mains(args []string, w io.Writer) error {
	clean := msidb.CoInit()
	defer clean()

	recordSeparator := ""
	for _, arg := range args {
		fmt.Fprint(w, recordSeparator)
		if err := showOne(arg, w); err != nil {
			return err
		}
		recordSeparator = "\n"
	}
	return nil
}

func main() {
	flag.Parse()
	if err := mains(flag.Args(), os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
