package metadata

import "testing"

func TestBuilder(t *testing.T) {
	var input = `
		machine:
		  name: "caxias"
		  networkInterfaces:
		  - name: "enps3"
		    ip: "192.168.0.2"
	    - name: "enp4s"
	      ip: "10.0.2.3"

		cluster:
		  name: "rio"
		  state: "new"
		  machines:
  	  - name: "macae"
		    clusterIp: "192.168.0.3"
	    - name: "campos"
		    clusterIp: "192.168.0.4"
	`

	spec := Build([]byte(input))

	if spec.machine.Name != "caxias" {
		t.Fatalf("expected caxias got %s", spec.machine.name)
	}
}
