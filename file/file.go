package file

import (
	"io/ioutil"
	"log"
	"os"
)

// Loads a file from path.
// Returns the byte array corresponding to file's contents.
func Load(path string) []byte {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}

	return file
}

// Gets the names of all files from directory passed as argument, like ls from Unix
// Returns a string array with these names
func Ls(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	var filenames []string

	if err != nil {
		log.Fatalf("-- error: %v", err)
		os.Exit(1)
	}

	for _, f := range files {
		filenames = append(filenames, f.Name())
	}

	return filenames
}

func CreateFileAndDir(dirname string, filename string) *os.File {
	err := os.Mkdir(dirname, os.ModePerm)
	if err != nil {
		log.Fatalf("---- error %v", err)
		os.Exit(1)
	}

	file, err := os.Create(dirname + "/" + filename)

	return file
}
