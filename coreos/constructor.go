package coreos

import (
	"github.com/bernardolins/vandame/metadata"
)

func Config(name string, config metadata.Config) *CoreOs {
	coreos := new(CoreOs)

	coreos.EtcdConfig.MachineName = name
	configureEtcd2(coreos, config)

	return coreos
}

func configureEtcd2(coreos *CoreOs, config metadata.Config) {
	coreos.EtcdConfig.InitialClusterToken = config.GetClusterToken()
	coreos.EtcdConfig.InitialClusterState = config.GetClusterState()
	setEtcd2InitialCluster(coreos, config.GetClusterNodes())
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

	config.EtcdConfig.InitialCluster = initialClusterString
}
