package file

import (
	"io/ioutil"
	"log"
	"os"
)

func Load(path string) []byte {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}

	return file
}
