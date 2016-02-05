package systemd

import (
	"github.com/bernardolins/vandame/file"
)

type Unit struct {
	Content []byte
	Name    string
	Command string
}

func UnitConfig(name, command, path string) *Unit {
	unit := new(Unit)

	content := file.Load(path)
	unit.Content = content

	unit.Name = name
	unit.Command = command

	return unit
}
