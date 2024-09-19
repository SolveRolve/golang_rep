package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type dataFlags struct {
	symLinks    *bool
	directories *bool
	files       *bool
	extension   *string
}

func deLinker(root, path string, b bool) {
	if !b {
		return
	}
	pathFromLink, ok := os.Readlink(root + path)

	if ok != nil {
		fmt.Println(path, "->", "[broken]")
	} else {
		fmt.Println(path, "->", pathFromLink)
	}
}
func flagInit() dataFlags {
	f := dataFlags{}
	f.symLinks = flag.Bool("sl", false, "print only symlinks")
	f.directories = flag.Bool("d", false, "print only directories")
	f.files = flag.Bool("f", false, "print only files")
	f.extension = flag.String("ext", "", "sort by extension of file")
	flag.Parse()

	return f
}
func printFiles(path, ext string, b bool) {
	if !b {
		return
	}
	if len(ext) == 0 {
		fmt.Println(path)
	} else {
		if strings.Contains(path, "."+ext) {
			fmt.Println(path)
		}
	}
}
func main() {

	on := flagInit()
	root := flag.Args()
	if len(root) == 0 {
		return
	}
	if !*on.files && !*on.symLinks && !*on.directories {
		*on.files, *on.symLinks, *on.directories = true, true, true
	}

	fileSystem := os.DirFS(root[0])
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {

		fi, _ := d.Info()
		strFi := fs.FormatFileInfo(fi)

		if fi.IsDir() && *on.directories {
			fmt.Println(path)
		} else {
			if !strings.Contains(strFi, "Lr") {
				printFiles(root[0]+path, *on.extension, *on.files)
			} else {
				deLinker(root[0], path, *on.symLinks)
			}
		}

		return nil
	})

}
