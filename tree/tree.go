package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

// сюда писать функцию DirTree

func dirTreeIteration(out io.Writer, path string, full bool, prefix string) error {
	// Get dirs and files in sorted list
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	// Remove files in not full mode
	var dirs []os.FileInfo
	if !full {
		for _, file := range files {
			if file.IsDir() {
				dirs = append(dirs, file)
			}
		}
		files = dirs
	}

	for i, file := range files {
		// Get prefixes
		currPrefix := prefix
		newPrefix := prefix
		if i < len(files) - 1 {
			currPrefix += "├───"
			newPrefix += "│\t"
		} else {
			currPrefix += "└───"
			newPrefix += "\t"
		}

		// Get postfixes
		postfix := "\n"
		if !file.IsDir() {
			if size := file.Size(); size == 0 {
				postfix = " (empty)" + postfix
			} else {
				postfix = " (" + strconv.FormatInt(int64(size), 10) + "b)" + postfix
			}
		}

		// Print
		_, err = fmt.Fprint(out, currPrefix + file.Name() + postfix)
		if err != nil {
			return  err
		}

		// Handle dirs
		if file.IsDir() {
			err = dirTreeIteration(out, path + string(os.PathSeparator) + file.Name(), full, newPrefix)
			if err != nil {
				return  err
			}
		}
	}
	return nil
}

func dirTree(out io.Writer, path string, full bool) error {
	return dirTreeIteration(out, path, full, "")
}