package main

import (
	"flag"
	"src/readDB/bin"
)

func main() {
	ok := flag.Bool("f", false, "read file")
	flag.Parse()
	if !*ok {
		panic("no flag -f")
	}
	fileName := flag.Args()

	inter := bin.InitReader(fileName[0])

	inter.ParseFile(fileName[0])

	inter.PrettyPrint()
}
