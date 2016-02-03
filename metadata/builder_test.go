package metadata

import "testing"

func TestBuilder(t *testing.T) {
	var input = `
cluster:
  kubernetes: true
  token: "rio"
  state: "new"
  nodes:
  - name: "macae"
    configIp: "192.168.0.3"
  - name: "campos"
    configIp: "192.168.0.4"
`

	config := Build([]byte(input))

	if config.Cluster.Token != "rio" {
		t.Fatalf("expected rio got %s", config.Cluster.Token)
	}

	if config.Cluster.State != "new" {
		t.Fatalf("expected new got %s", config.Cluster.State)
	}
}
