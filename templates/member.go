package templates

type Node struct {
	Units Units
}

func (m *Node) AddUnitTemplate(units Units) {
	m.Units = units
}
