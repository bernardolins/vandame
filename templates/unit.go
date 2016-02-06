package templates

import (
	"github.com/bernardolins/vandame/cluster"
	"github.com/bernardolins/vandame/systemd"
)

type Units struct {
	Template string
}

func (u *Units) SetText(member *cluster.Member) *Units {
	u.setUnitsTemplateString(member.Units)

	return u
}

func (u *Units) setUnitsTemplateString(units []systemd.Unit) {
	if len(units) == 0 {
		u.Template = ""
	} else {
		u.Template = "units:\n"

		for _, unit := range units {
			u.setUnitName(unit.Name)
			u.setUnitCommand(unit.Command)
		}
	}
}

func (u *Units) setUnitName(name string) {
	u.Template = u.Template + "  - name: " + name + "\n"
}

func (u *Units) setUnitCommand(command string) {
	if command != "" {
		u.Template = u.Template + "  command: " + command + "\n"
	}
}
