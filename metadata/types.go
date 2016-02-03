package metadata

type Config struct {
	Cluster Cluster
}

//
// This struct represents configs for a new Cluster
//
type Cluster struct {
	Kubernetes bool   `yaml: "kubernetes,omitempty"`
	Token      string `yaml: "token"`
	State      string `yaml: "state"`
	Nodes      []Node `yaml: "nodes"`
}

func (config *Config) GetClusterToken() string {
	return config.Cluster.Token
}

func (config *Config) GetClusterState() string {
	return config.Cluster.State
}

func (config *Config) GetUseKubernetes() bool {
	return config.Cluster.Kubernetes
}

func (config *Config) GetClusterNodes() []Node {
	return config.Cluster.Nodes
}

//
// Node is a machine in the cluster
//
type Node struct {
	Name      string `yaml: "name"`
	Ip        string `yaml: "ip"`
	Interface string `yaml: "interface"`
}

func (node *Node) GetNodeName() string {
	return node.Name
}

func (node *Node) GetNodeIp() string {
	return node.Ip
}

func (node *Node) GetNodeInterface() string {
	return node.Interface
}
