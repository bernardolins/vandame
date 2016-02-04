package cluster

import (
	"github.com/bernardolins/vandame/coreos"
	"github.com/bernardolins/vandame/metadata"
)

type Member struct {
	CoreOS coreos.CoreOs
	Node   metadata.Node
}

func MemberConfig(node metadata.Node, coreos coreos.CoreOs) *Member {
	member := new(Member)

	member.CoreOS = coreos
	member.Node = node

	return member
}
