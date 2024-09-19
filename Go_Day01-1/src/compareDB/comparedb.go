package main

import (
	"flag"
	"fmt"
	"src/readDB/bin"
)

func removedOrChange(old, new bin.Data) {
	for i, d := range old.Cake {
		count := 0
		countCake := 0
		for j, data := range new.Cake {
			if data.Name != "" {
				countCake++
			}
			if data.Name != d.Name && len(d.Name) != 0 && len(data.Name) != 0 {
				count++
			}
			if data.Name == d.Name && len(d.Name) != 0 && len(data.Name) != 0 {
				changed(old, new, i, j)
			}
		}
		if count == countCake {
			fmt.Printf("REMOVED cake \"%v\"\n", d.Name)
		}
	}
}
func added(old, new bin.Data) {
	for _, d := range new.Cake {
		count := 0
		countCake := 0
		for _, data := range old.Cake {
			if data.Name != "" {
				countCake++
			}
			if data.Name != d.Name && len(d.Name) != 0 && len(data.Name) != 0 {
				count++
			}
		}
		if count == countCake {
			fmt.Printf("ADDED cake \"%v\"\n", d.Name)
		}
	}
}
func changed(old, new bin.Data, i, j int) {

	if new.Cake[j].Stovetime != old.Cake[i].Stovetime {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", new.Cake[j].Name, new.Cake[j].Stovetime, old.Cake[i].Stovetime)
	}

	for _, ingrOld := range old.Cake[i].Ingredients.Item {
		countOfINgredient := 0
		counter := 0
		for _, ingrNew := range new.Cake[j].Ingredients.Item {
			if ingrNew.Name != "" {
				countOfINgredient++
			}
			ok := len(ingrNew.Name) != 0 && len(ingrOld.Name) != 0

			if ingrOld.Name != ingrNew.Name && ok {
				counter++
			}
			if ingrOld.Name == ingrNew.Name && ok {
				notEmpty := len(ingrNew.Unit) != 0 && len(ingrOld.Unit) != 0
				if ingrOld.Unit != ingrNew.Unit && notEmpty {
					fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", ingrOld.Name, old.Cake[i].Name, ingrNew.Unit, ingrOld.Unit)
				}
				okk := len(ingrNew.Count) != 0 && len(ingrOld.Count) != 0
				if ingrOld.Count != ingrNew.Count && okk {
					fmt.Printf("CHANGED unit count for ingredient  \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", ingrOld.Name, old.Cake[i].Name, ingrNew.Count, ingrOld.Count)
				}
				if ingrOld.Unit != ingrNew.Unit && len(ingrNew.Unit) == 0 {
					fmt.Printf("REMOVED unit \"%s\" for cake  \"%s\"\n", ingrOld.Unit, old.Cake[i].Name)
				}
			}
		}
		if counter == countOfINgredient {
			fmt.Printf("REMOVED ingredient \"%v\" for cake  \"%v\"\n", ingrOld.Name, new.Cake[j].Name)
		}
	}
}

func main() {
	old := flag.String("old", "", "read file")
	new := flag.String("new", "", "read file")

	flag.Parse()
	if len(*old) == 0 || len(*new) == 0 {
		panic("not have all flags")
	}

	fileNew := bin.InitReader(*new)
	fileOld := bin.InitReader(*old)

	fileOld.ParseFile(*old)
	fileNew.ParseFile(*new)

	n := fileNew.CopyToData()
	o := fileOld.CopyToData()

	removedOrChange(*o, *n)
	added(*o, *n)

}
