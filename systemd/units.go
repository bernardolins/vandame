package systemd

import (
	"github.com/bernardolins/vandame/file"
)

type Unit struct {
	Content string
	Name    string
	Command string
}

func UnitConfig(name, command, path string) *Unit {
	unit := new(Unit)

	if path != "" {
		content := file.Load(path)
		unit.Content = string(content)
	}

	unit.Name = name
	unit.Command = command

	return unit
}
