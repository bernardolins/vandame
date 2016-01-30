package coreos

import (
	"github.com/bernardolins/vandame/metadata"
)

func Config(spec metadata.Specification) *CoreOs {
	config := new(CoreOs)

	configureEtcd2(config, spec)
	configureMachine(config, spec)

	return config
}

func configureEtcd2(config *CoreOs, spec metadata.Specification) {
	config.EtcdConfig.MachineName = spec.Machine.Name
	config.EtcdConfig.InitialClusterToken = spec.Cluster.Name
	setEtcd2InitialCluster(config, spec.Cluster.Machines)
}

func setEtcd2InitialCluster(config *CoreOs, machines []metadata.ClusterMachine) {
	initialClusterString := ""
	for _, machine := range machines {

		if initialClusterString != "" {
			initialClusterString = initialClusterString + ","
		}

		machineString := machine.Name + "=http://" + machine.Ip + ":2380"
		initialClusterString = initialClusterString + machineString
	}

	config.EtcdConfig.InitialCluster = initialClusterString
}

func configureMachine(config *CoreOs, spec metadata.Specification) {
	config.Machine.Ip = spec.Machine.Ip
	config.Machine.Interface = spec.Machine.Interface
}
