package metadata

type Machine struct {
	Name       string      `yaml: "name"`
	Interfaces []Interface `yaml: "interfaces,omitempty"`
}

type Interface struct {
	Name string `yaml: "name"`
	Ip   string `yaml: "ip"`
}

type Cluster struct {
	Name     string           `yaml: "name"`
	State    string           `yaml: "state"`
	Machines []ClusterMachine `yaml: "machines"`
}

type ClusterMachine struct {
	Name      string `yaml: "name"`
	clusterIp string `yaml: "name"`
}
