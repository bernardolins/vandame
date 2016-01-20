package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// VandameCommand is the command that represents the binary itselfs.
var VandameCommand = &cobra.Command{
	Use:   "vandame",
	Short: vandameShort(),
	Long:  vandameLong(),
	Run: func(cmd *cobra.Command, args []string) {
		printHelp()
	},
}

func vandameShort() string {
	return "vandameShort"
}

func vandameLong() string {
	return "vandameLong"
}

// Prints the help message for vandame command
func printHelp() {
	var helpmsg = `Vandame is a command line tool to help you handle cluster initialization and updates.
	
	commands:
		generate: Generates initial configuration. CoreOS cloud-config is supported.
`

	fmt.Println(helpmsg)
}
