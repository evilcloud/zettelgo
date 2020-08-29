package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mozillazg/go-unidecode"
	"gopkg.in/djherbis/times.v1"
)

const (
	directory string = "/Users/kg/Dropbox/academic"
	numLength = 12
)

type fileNameInfo struct {
	pureName  string
	extension string
	id        string
	title     string
}

// Generate string with the current timestamp
func timeId() string {
	t := time.Now()
	return fmt.Sprintf(t.Format("20060102150405"))
}

func preTimeId(t time.Time) string {
	return fmt.Sprintf(t.Format("20060102150405"))
}

// Rename file
func renameFile(dir, oldFile, newFile string) (error, string) {
	n := newFile
	if isFile(dir, n){
		fullName := trinitifyFileName(n)
		timestamp := timeId()
		fmt.Printf("File %s already exists, adding %s at the end \n", n, timestamp)
		n = fullName.pureName + "_" + timestamp + fullName.extension
	}

	err := os.Rename(path.Join(dir, oldFile), path.Join(dir, n))
	if err != nil {
		log.Println(err)
	}
	return err, n
}

// returns list of files from given directory with given list of extensions
func ListFiles(extensions []string) []string {
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

func isDir(dir string) bool {
	_, err := os.Stat(directory)
	return err == nil
	}


// Split a filename into components and return struct
func trinitifyFileName(fileName string) fileNameInfo {
	var (
		possibleId []string
		ret        fileNameInfo
	)

	ret.extension = filepath.Ext(fileName)
	ret.pureName = strings.TrimSuffix(fileName, ret.extension)

	possibleId = strings.Split(ret.pureName, " ")
	_, err := strconv.Atoi(possibleId[0])
	if err == nil && len(possibleId[0]) >= numLength{
		ret.id = possibleId[0]
		ret.title = strings.TrimSpace(strings.TrimPrefix(ret.pureName, ret.id))
	} else {
		ret.id = ""
		ret.title = ret.pureName
	}
	return ret
}

func fixFilename(filename string) string {
	f := trinitifyFileName(filename)
	fTitle := unidecode.Unidecode(f.title)
	newId := f.id
	if newId == "" {
		fTime, err := times.Stat(path.Join(directory, filename))
		if err != nil {
			log.Panic(err)
		}
		if fTime.HasBirthTime() {
			newId = preTimeId(fTime.BirthTime())
		} else {
			newId = timeId()
		}
	}
	return fmt.Sprintf(newId + " " + fTitle + f.extension)
}

func main() {
	fileNames := ListFiles([]string{".md"})

	//var toFix map[string]string
	for i, fileName := range fileNames {
		fixedName := fixFilename(fileName)
		if fileName != fixedName {
			fmt.Printf("%d. was: [%s]\n   now: [%s]\n\n", i, fileName, fixedName)
			//toFix[fileName] = fixedName
		} else {
			//fmt.Printf("%d. %s is a valid file name", i, fileName)
			_ = true
		}
	}
}
