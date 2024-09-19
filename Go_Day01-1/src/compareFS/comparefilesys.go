package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"sort"
	"strings"
)

func dumpFileSystem(s string) {
	root := s
	fileSystem := os.DirFS(root)
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {

		indexOfPrefix := strings.LastIndex(path, "/")
		indexOfPoint := strings.LastIndex(path, ".")
		if indexOfPrefix < indexOfPoint && indexOfPrefix >= 0 && indexOfPoint >= 0 {
			fmt.Println(path)
		}

		return nil
	})
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func compareSortedStr(old, new []string) []string {
	result := make([]string, 0, 10)
	findInOld := false
	for _, strOld := range old {
		findInOld = false
		for _, strNew := range new {
			if strOld == strNew {
				findInOld = true
				break
			}
		}
		if !findInOld && len(strOld) != 0 {
			result = append(result, strOld)
		}
	}
	return result
}
func sortFiles(s string) []string {
	sortedFiles := make([]string, 0, 10000)

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sortedFiles = append(sortedFiles, scanner.Text())
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}

	sort.Strings(sortedFiles)
	return sortedFiles
}

func main() {

	old := flag.String("old", "", "read file")
	new := flag.String("new", "", "read file")
	makeDump := flag.String("d", "", "make dump  for path")
	flag.Parse()

	if (len(*old) == 0 || len(*new) == 0) && len(*makeDump) == 0 {
		panic("not have all flags")
	}

	if len(*makeDump) != 0 {
		dumpFileSystem(*makeDump)
		return
	}

	oldFiles := sortFiles(*old)
	newFiles := sortFiles(*new)
	addedStr := compareSortedStr(oldFiles, newFiles)
	removedStr := compareSortedStr(newFiles, oldFiles)

	for _, str := range addedStr {
		fmt.Printf("ADDED %s\n", str)
	}
	for _, str := range removedStr {
		fmt.Printf("REMOVED %s\n", str)
	}

}
