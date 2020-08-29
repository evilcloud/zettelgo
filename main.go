package main

import (
	"fmt"
)

const (
	directory string = "/Users/kg/Dropbox/academic"
	numLength        = 12
)

func main() {
	fileNames := listFiles([]string{".md"})
	var fixed []string
	toFix := make(map[string]string)
	for _, fileName := range fileNames {
		fixedName := fixFilename(fileName)
		if fileName != fixedName {
			//fmt.Printf("%d. was: [%s]\n   now: [%s]\n\n", i, fileName, fixedName)
			toFix[fileName] = fixedName
		} else {
			//fmt.Printf("%d. %s is a valid file name", i, fileName)
			fixed = append(fixed, fileName)
		}
	}
	for key, value := range toFix {
		fmt.Printf("Key: %s\nVal: %s\n\n", key, value)
	}
	fmt.Printf("Fixed %d, to fix: %d\nTotal: %d\n", len(fixed), len(toFix), len(fixed)+len(toFix))
}
