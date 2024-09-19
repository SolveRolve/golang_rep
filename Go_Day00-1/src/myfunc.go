package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
)

type data struct {
	NumberMap map[int]int
	Numbers   []int
	maxRepit  int
}

type flags struct {
	Mean,
	Median,
	SD,
	Mode *bool
}

func initFlags() *flags {
	var f flags
	f.Mean = flag.Bool("mean", false, "for mean")
	f.Median = flag.Bool("median", false, "for median")
	f.SD = flag.Bool("sd", false, "for stand dispersion")
	f.Mode = flag.Bool("mode", false, "for mode")
	flag.Parse()
	return &f
}

func (d *data) add(number int) {
	if d == nil {
		panic("in func mode *data == nil")
	}
	d.Numbers = append(d.Numbers, number)
	d.NumberMap[number]++
}

func NewData() *data {
	var d data
	d.NumberMap = make(map[int]int)
	d.Numbers = make([]int, 0, 100)
	d.maxRepit = 0
	return &d
}

func (d *data) Mode() int {
	if d == nil {
		panic("in func mode *data == nil")
	}
	var repited []int
	if len(d.Numbers) == 0 {
		return 0
	}
	for _, value := range d.NumberMap {
		if d.maxRepit < value {
			d.maxRepit = value
		}
	}
	for key, value := range d.NumberMap {
		if value == d.maxRepit {
			repited = append(repited, key)
		}
	}

	sort.Ints(repited)
	return repited[0]
}

func (d *data) Mean() (result float64) {
	if d == nil {
		panic("in func mean *data == nil")
	}
	for _, value := range d.Numbers {
		result += float64(value)
	}
	if len(d.Numbers) != 0 {
		result /= float64(len(d.Numbers))
	} else {
		result = 0.0
	}

	return
}

func (d *data) Median() (result float64) {
	if d == nil {
		panic("in func median *data == nil")
	}
	sort.Ints(d.Numbers)
	if len(d.Numbers) == 0 {
		return 0.0
	} else if len(d.Numbers)%2 == 1 {
		result = float64(d.Numbers[len(d.Numbers)/2])
	} else {
		result = (float64(d.Numbers[len(d.Numbers)/2-1]) + float64(d.Numbers[len(d.Numbers)/2])) / 2
	}
	return
}

func (d *data) StandDispersion() (result float64) {
	if d == nil {
		panic("in func standDispersion *data == nil")
	}
	countNum := len(d.Numbers)
	if countNum == 0 || countNum == 1 {
		return 0.0
	}
	sumDispersion := 0.0
	mean := d.Mean()
	for _, val := range d.Numbers {
		sumDispersion += math.Pow((float64(val) - mean), 2)
	}
	sumDispersion /= float64(countNum - 1)

	return math.Sqrt(sumDispersion)
}

func printResult(f *flags, d data) {
	if &d == nil {
		panic("in func printResult *data == nil")
	} else if f == nil {
		panic("in func standDispersion *data == nil")
	}

	if !*f.Mean && !*f.Median && !*f.SD && !*f.Mode {
		fmt.Printf("Mode: = %v\n", d.Mode())
		fmt.Printf("Mean: = %.1f\n", d.Mean())
		fmt.Printf("Median: = %.1f\n", d.Median())

		fmt.Printf("SD: = %.1f\n", d.StandDispersion())
	} else {
		if *f.Mean {
			fmt.Printf("Mean: = %.1f\n", d.Mean())
		}
		if *f.Median {
			fmt.Printf("Median: = %.1f\n", d.Median())
		}
		if *f.SD {
			fmt.Printf("SD: = %.1f\n", d.StandDispersion())
		}
		if *f.Mode {
			fmt.Printf("Mode: = %v\n", d.Mode())
		}
	}
}
