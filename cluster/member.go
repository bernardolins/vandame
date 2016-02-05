package cluster

import (
	"github.com/bernardolins/vandame/coreos"
	"github.com/bernardolins/vandame/metadata"
	"github.com/bernardolins/vandame/systemd"
)

type Member struct {
	CoreOS coreos.CoreOs
	Node   metadata.Node
	Units  []systemd.Unit
}

func MemberConfig(node metadata.Node, coreos coreos.CoreOs) *Member {
	member := new(Member)

	member.CoreOS = coreos
	member.Node = node

	return member
}

func (member *Member) AddUnit(unit systemd.Unit) {
	member.Units = append(member.Units, unit)
}
