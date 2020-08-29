package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/djherbis/times.v1"
)

type fileNameInfo struct {
	pureName  string
	extension string
	id        string
	title     string
}

// Rename file
func renameFile(dir, oldFile, newFile string) (string, error) {
	n := newFile
	if isFile(dir, n) {
		fullName := trinitifyFileName(n)
		timestamp := timeID()
		fmt.Printf("File %s already exists, adding %s at the end \n", n, timestamp)
		n = fullName.pureName + "_" + timestamp + fullName.extension
	}

	err := os.Rename(path.Join(dir, oldFile), path.Join(dir, n))
	if err != nil {
		log.Println(err)
	}
	return n, err
}

// returns list of files from given directory with given list of extensions
func listFiles(extensions []string) []string {
	if !isDir(directory) {
		log.Panic("directory %s not found")
	}
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Panic(err)
	}

	// check for presence of given extensions
	extMatch := func(extensions []string, extension string) bool {
		for _, ext := range extensions {
			if ext == extension {
				return true
			}
		}
		return false
	}

	// put together the valid files and return the list
	var ret []string
	for _, fileName := range files {
		if extMatch(extensions, filepath.Ext(fileName.Name())) {
			ret = append(ret, fileName.Name())
		}
	}
	return ret
}

// Checking if file exists
func isFile(dir, name string) bool {
	_, err := os.Stat(path.Join(dir, name))
	return err == nil
}

// Checking if directory exists
func isDir(dir string) bool {
	_, err := os.Stat(directory)
	return err == nil
}

func fixFilename(filename string) string {
	f := trinitifyFileName(filename)
	fTitle := cleanText(f.title)
	newID := f.id
	if newID == "" {
		fTime, err := times.Stat(path.Join(directory, filename))
		if err != nil {
			log.Panic(err)
		}
		if fTime.HasBirthTime() {
			newID = timeID(fTime.BirthTime())
		} else {
			newID = timeID()
		}
	}
	return fmt.Sprintf(strings.TrimSpace(newID+" "+fTitle) + f.extension)
}
