package main

import (
	"fmt"
	"github.com/bernardolins/vandame/cmd"
	"os"
)

func main() {
	if err := cmd.VandameCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
