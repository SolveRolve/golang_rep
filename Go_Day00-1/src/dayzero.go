package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	newScaner := bufio.NewScanner(os.Stdin)
	record := NewData()

	for newScaner.Scan() {
		str := newScaner.Text()
		num, err := strconv.Atoi(str)

		if err == nil {
			record.add(num)
		} else {
			fmt.Println("is not a integer number!\nrange for input (-1e+6:1e+6)")
		}
	}

	printResult(initFlags(), *record)

}
