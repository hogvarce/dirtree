package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	fileObj, err := os.Open(path)
	defer fileObj.Close()
	if err != nil {
		log.Fatalf("Could not open %s: %s", path, err.Error())
	}
	fileName := fileObj.Name()
	files, err := ioutil.ReadDir(fileName)
	if err != nil {
		log.Fatalf("Could not read dir names in %s: %s", path, err.Error())
	}
	for _, file := range files {
		if !printFiles && file.IsDir() {

			out.Write([]byte("├───" + file.Name() + "\n\t"))
			dirTree(out, path+string(os.PathSeparator)+file.Name(), printFiles)
		} else if printFiles {
			out.Write([]byte("├───" + file.Name() + "\n\t"))
			if file.IsDir() {
				dirTree(out, path+string(os.PathSeparator)+file.Name(), printFiles)
			}
		}
	}

	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
