package coreos

type CoreOs struct {
	Etcd Etcd2
}

type Etcd2 struct {
	MachineName         string
	InitialCluster      string
	InitialClusterState string
	InitialClusterToken string
	Endpoints           string
}
