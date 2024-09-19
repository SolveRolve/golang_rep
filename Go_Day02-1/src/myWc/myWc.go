package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

type dataFlags struct {
	counterChar *bool
	counterLine *bool
	CounterWord *bool
}

func flagInit() dataFlags {
	f := dataFlags{}
	f.CounterWord = flag.Bool("w", false, "counter word")
	f.counterLine = flag.Bool("l", false, "counter line")
	f.counterChar = flag.Bool("m", false, "counter character")
	flag.Parse()
	return f
}
func workWithFile(name string, f dataFlags, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanerFile := bufio.NewScanner(file)
	scanerFile.Split(bufio.ScanLines)

	countChar := 0
	countLine := 0
	CountWord := 0

	for scanerFile.Scan() {
		line := scanerFile.Text()
		if *f.counterChar {
			countChar += utf8.RuneCountInString(line)
		}
		if *f.counterLine {
			countLine++
		}
		if *f.CounterWord {
			CountWord += strings.Count(line, " ")
		}
	}
	if *f.counterLine {
		fmt.Printf("%d\t%s\n", countLine, name)
	}
	if *f.CounterWord {
		fmt.Printf("%d\t%s\n", CountWord, name)
	}
	if *f.counterChar {
		fmt.Printf("%d\t%s\n", countChar, name)
	}

}
func main() {
	//start := time.Now()
	mode := flagInit()
	var wg sync.WaitGroup

	if *mode.counterChar && *mode.counterLine || *mode.counterChar && *mode.CounterWord || *mode.counterLine && *mode.CounterWord {
		print("used to many flag, use only one")
		return
	}
	files := flag.Args()

	for _, fileName := range files {
		wg.Add(1)
		go workWithFile(fileName, mode, &wg)
	}
	wg.Wait()
	//duration := time.Since(start)
	//fmt.Println(duration)
}
