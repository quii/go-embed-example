package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
)

//go:embed embedded-data
var data embed.FS

//go:embed embedded-data/file1.txt
var contentsOfFile1 string

func main() {
	{
		subDir, _ := fs.Sub(data, "embedded-data")
		poop := findThePoop(subDir)
		fmt.Printf("the files containing poop in the embeded file system are %v\n", poop)
	}
	{
		folderPath := "normal-fs"
		normalFS := os.DirFS(folderPath)
		poop := findThePoop(normalFS)
		fmt.Printf("the files containing poop in %q are %v", folderPath, poop)
	}
}

func findThePoop(filesystem fs.FS) (poopyfiles []string) {
	dir, _ := fs.ReadDir(filesystem, ".")
	for _, file := range dir {
		if fileIsPoopy(filesystem, file.Name()) {
			poopyfiles = append(poopyfiles, file.Name())
		}
	}
	return
}

func fileIsPoopy(filesystem fs.FS, fileName string) bool {
	f, _ := filesystem.Open(fileName)
	defer f.Close()
	data, _ := io.ReadAll(f)
	isPoopy := strings.Contains(string(data), "poop")
	return isPoopy
}
