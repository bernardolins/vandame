package templates

import (
	"github.com/bernardolins/vandame/cluster"
)

type Template interface {
	SetText(member *cluster.Member)
}
