package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Generate string with the current timestamp
func timeId(args ...time.Time) string {
	t := time.Now()
	if len(args) > 1 {
		t = args[0]
	} else {
		t = time.Now()
	}
	return fmt.Sprintf(t.Format("20060102150405"))
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
	if err == nil && len(possibleId[0]) >= numLength {
		ret.id = possibleId[0]
		ret.title = strings.TrimSpace(strings.TrimPrefix(ret.pureName, ret.id))
	} else {
		ret.id = ""
		ret.title = ret.pureName
	}
	return ret
}
