package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

func makerTar(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	nameTar := path[:strings.LastIndex(path, ".")] + "_" + timeStamp + ".tag.gz"
	cmd := exec.Command("tar", "czf", nameTar, "-P", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	var wg sync.WaitGroup
	args := os.Args[1:]

	for _, path := range args {
		wg.Add(1)
		go makerTar(path, &wg)
	}
	wg.Wait()
}
