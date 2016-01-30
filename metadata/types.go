package metadata

type Specification struct {
	Machine Machine `yaml: "machine"`
	Cluster Cluster `yaml: "cluster"`
}

type Machine struct {
	Name      string `yaml: "name"`
	Ip        string `yaml: "ip"`
	Interface string `yaml: "interface"`
	//Interfaces []Interface `yaml: "interfaces,omitempty"`
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
	Name string `yaml: "name"`
	Ip   string `yaml: "ip"`
}
