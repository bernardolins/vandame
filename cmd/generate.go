package cmd

import (
	"github.com/bernardolins/vandame/coreos"
	"github.com/bernardolins/vandame/file"
	"github.com/bernardolins/vandame/metadata"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

// This function is called to insert
func init() {
	generateCommand := NewGenerateCommand()
	VandameCommand.AddCommand(generateCommand.command)
}

// GenerateCommand object structure.
type GenerateCommand struct {
	command   *cobra.Command
	input     string
	templates string
	output    string
}

// The generate command constructor. Returns a GenerateCommand object containing flag values and a cobra.Command object.
func NewGenerateCommand() *GenerateCommand {
	generate := new(GenerateCommand)

	generate.command = &cobra.Command{
		Use: "generate",
		Run: func(cmd *cobra.Command, args []string) {
			generate.run()
		},
	}

	generate.command.Flags().StringVarP(&generate.input, "input", "f", "./", "Path to input file")
	generate.command.Flags().StringVarP(&generate.templates, "templates", "t", "./", "Path to templates directory")
	generate.command.Flags().StringVarP(&generate.output, "output", "o", "./", "Path to output cloud-config")

	return generate
}

// Runs the generate command
func (generate *GenerateCommand) run() {
	file := file.Load(generate.input)
	spec := metadata.Build(file)
	config := coreos.Config(spec)

	generate.loadTemplates(config)
}

// Helper file to load templates. May be extracted to another module.
func (generate *GenerateCommand) loadTemplates(config *coreos.CoreOs) {
	templateFiles := file.Ls(generate.templates)

	for _, f := range templateFiles {
		t := template.New(f)
		t, _ = t.ParseFiles(generate.templates + f)
		t.Execute(os.Stdout, config)
	}
}