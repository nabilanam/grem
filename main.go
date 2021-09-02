package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Print("\nGrem, a recursive extension wise file remover.")
		fmt.Print("\n\n\tUsage: \t grem [start-path] [.extension]\n\n")
	} else {
		files := find(os.Args[1], os.Args[2])
		for _, file := range files {
			os.Remove(file)
		}
	}
}

func find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}
