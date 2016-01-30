package coreos

type CoreOs struct {
	Machine    CoreOsMachine
	EtcdConfig Etcd2
}

type CoreOsMachine struct {
	Ip        string
	Interface string
}

type Etcd2 struct {
	MachineName         string
	InitialCluster      string
	InitialClusterState string
	InitialClusterToken string
}
