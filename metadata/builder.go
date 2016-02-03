package metadata

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func Build(bytes []byte) Config {
	cluster := serialize(bytes)

	return cluster
}

func serialize(metadata []byte) Config {
	cluster := Config{}

	err := yaml.Unmarshal(metadata, &cluster)

	if err != nil {
		log.Fatalf("-- error %v", err)
		os.Exit(1)
	}

	return cluster
}
