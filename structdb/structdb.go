package structdb

import (
	"time"
)

type ClusterPool struct {
	ID           int
	Name         string
	Namespace    string
	ClusterIP    string
	ClusterPort  int
	ClusterToken string
	ClusterInfo  string
	CreateTime   time.Time
}

type ConfigmapsPool struct {
	ID            int
	TmplName      string
	Namespace     string
	Labels        string
	Data          string
	URL           string
	HostIPPorts   string
	ConfigmapName string
	ClusterID     int
	CreateTime    time.Time
}

type DaemonetsPool struct {
	ID          int
	Name        string
	Selector    string
	Labels      string
	Status      string
	PodTemplate string
	ClusterID   int
	CreateTime  time.Time
}

type EndpointsPool struct {
	ID         int
	Name       string
	Namespace  string
	Labels     string
	Subsets    string
	ClusterID  int
	CreateTime time.Time
}

type HostsPool struct {
	ID                  int
	Labels              string
	HostIP              string
	HostName            string
	HostStatus          string
	OS                  string
	KubeletVersion      string
	KubeletProxyVersion string
	HostNetwork         string
	HostCPU             string
	HostMemory          string
	HostUsedCPU         string
	HostUsedMemory      string
	RunningPod          string
	ClusterID           int
	CreateTime          time.Time
}

type NamespacesPool struct {
	ID         int
	Name       string
	Labels     string
	Status     string
	ClusterID  int
	CreateTime time.Time
}

type PodPool struct {
	ID            int
	Name          string
	Namespace     string
	Version       string
	HostIP        string
	PodIP         string
	ImageName     string
	Ports         string
	Resource      string
	Environment   string
	MountsVolumes string
	StartTime     time.Time
	PodStatus     int
	RestartCount  int
	ClusterID     int
	CreateTime    time.Time
}

type PoolController struct {
	ID                     int
	ClusterLastVersion     int
	ConfigmapsLastVersion  int
	DaemonsetsLastVersion  int
	DeploymentsLastVersion int
	EndpointsLastVersion   int
	HostsLastVersion       int
	NamespacesLastVersion  int
	PodLastVersion         int
	ReplicasetsLastVersion int
	SecretsLastVersion     int
	ServicesLastVersion    int
}

type ReplicasetsPool struct {
	ID           int
	Name         string
	Namespace    string
	Labels       string
	Annotations  string
	ControlledBy string
	PodStatus    string
	PodTemplate  string
	ClusterID    int
	CreateTime   time.Time
}

type SecretsPool struct {
	ID         int
	Name       string
	Namespace  string
	Labels     string
	Type       string
	Data       string
	ClusterID  int
	CreateTime time.Time
}

type ServicesPool struct {
	ID            int
	Name          string
	Namespace     string
	Selector      string
	Type          string
	ServiceIP     string
	EndpointsInfo string
	ClusterID     int
	CreateTime    time.Time
}
