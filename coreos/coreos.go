package coreos

type CoreOs struct {
	EtcdConfig Etcd2
}

type Etcd2 struct {
	MachineName         string
	InitialCluster      string
	InitialClusterState string
	InitialClusterToken string
}
