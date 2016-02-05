package coreos

import (
	"github.com/bernardolins/vandame/metadata"
)

func Config(name string, config metadata.Config) *CoreOs {
	coreos := new(CoreOs)

	coreos.Etcd.MachineName = name
	configureEtcd2(coreos, config)

	return coreos
}

func configureEtcd2(coreos *CoreOs, config metadata.Config) {
	coreos.Etcd.InitialClusterToken = config.GetClusterToken()
	coreos.Etcd.InitialClusterState = config.GetClusterState()
	setEtcd2InitialCluster(coreos, config.GetClusterNodes())
	setEtcd2Endpoints(coreos, config.GetClusterNodes())
}

func setEtcd2InitialCluster(config *CoreOs, nodes []metadata.Node) {
	initialClusterString := ""
	for _, node := range nodes {

		if initialClusterString != "" {
			initialClusterString = initialClusterString + ","
		}

		nodeString := node.GetNodeName() + "=http://" + node.GetNodeIp() + ":2380"
		initialClusterString = initialClusterString + nodeString
	}

	config.Etcd.InitialCluster = initialClusterString
}

func setEtcd2Endpoints(config *CoreOs, nodes []metadata.Node) {
	endpoints := ""
	for _, node := range nodes {

		if endpoints != "" {
			endpoints = endpoints + ","
		}

		nodeString := "=http://" + node.GetNodeIp() + ":2379"
		endpoints = endpoints + nodeString
	}

	config.Etcd.Endpoints = endpoints
}
