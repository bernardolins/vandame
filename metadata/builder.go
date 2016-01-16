package metadata

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func Build(bytes []byte) Specification {
	spec := serialize(bytes)

	return spec
}

func serialize(metadata []byte) Specification {
	spec := Specification{}

	err := yaml.Unmarshal(metadata, &spec)

	if err != nil {
		log.Fatalf("-- error %v", err)
		os.Exit(1)
	}

	return spec
}
