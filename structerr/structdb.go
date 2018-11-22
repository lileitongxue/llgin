package structerr

import (
	"database/sql"
	"time"
)

type ClusterPool struct {
	ID           sql.NullInt64   `json:"id" form:"id"`
	Name         sql.NullString  `json:"name" form:"name"`
	Namespace    sql.NullString  `json:"namespace" form:"namespace"`
	ClusterIP    sql.NullString  `json:"cluster_ip" form:"cluster_ip"`
	ClusterPort  sql.NullFloat64 `json:"cluster_port" form:"cluster_port"`
	ClusterToken sql.NullString  `json:"cluster_token" form:"cluster_token"`
	ClusterInfo  sql.NullString  `json:"cluster_info" form:"cluster_info"`
	CreateTime   time.Time       `json:"create_time" form:"create_time"`
}

type ConfigmapsPool struct {
	ID            int       `json:"id" form:"id"`
	TmplName      string    `json:"tmpl_name" form:"tmpl_name"`
	Namespace     string    `json:"namespace" form:"namespace"`
	Labels        string    `json:"labels" form:"labels"`
	Data          string    `json:"data" form:"data"`
	URL           string    `json:"url" form:"url"`
	HostIPPorts   string    `json:"hostip_ports" form:"hostip_ports"`
	ConfigmapName string    `json:"configmap_name" form:"configmap_name"`
	ClusterID     int       `json:"cluster_id" form:"cluster_id"`
	CreateTime    time.Time `json:"create_time" form:"create_time"`
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

type DeploymentsPool struct {
	//ID         int       `json:"id" form:"id"`
	//DeployID   int       `json:"deploy_id" form:"deploy_id"`
	Name string `json:"name" form:"name"`
	//Namespace  string    `json:"namespace" form:"namespace"`
	Labels    string `json:"labels" form:"labels"`
	Version   string `json:"version" form:"version"`
	Selector  string `json:"selector" form:"selector"`
	Desired   int    `json:"desired" form:"desired"`
	Availabel int    `json:"availabel" form:"availabels"`
	//ClusterID  int       `json:"cluster_id" form:"cluster_id"`
	CreateTime time.Time `json:"create_time" form:"create_time"`
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
	//ID                  int       `json:"id" form:"id"`
	Labels              []string `json:"labels" form:"labels"`
	HostIP              string   `json:"host_ip" form:"host_ip"`
	HostName            string   `json:"hostname" form:"hostname"`
	HostStatus          string   `json:"hosts_status" form:"hosts_status"`
	OS                  string   `json:"os" form:"os"`
	KubeletVersion      string   `json:"kubelet_version" form:"kubelet_version"`
	KubeletProxyVersion string   `json:"kubelet_proxy_version" form:"kubelet_proxy_version"`
	HostNetwork         string   `json:"host_network" form:"host_network"`
	HostCPU             string   `json:"host_cpu" form:"host_cpu"`
	HostMemory          string   `json:"host_memory" form:"host_memory"`
	HostUsedCPU         string   `json:"host_used_cpu" form:"host_used_cpu"`
	HostUsedMemory      string   `json:"host_used_memory" form:"host_used_memory"`
	RunningPod          string   `json:"running_pod" form:"running_pod"`
	//ClusterID           int       `json:"" form:""`
	CreateTime time.Time `json:"create_time" form:"create_time"`
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
	//ID            int       `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	//Namespace     string    `json:"namespace" form:"namespace"`
	Version       string    `json:"version" form:"version"`
	HostIP        string    `json:"host_ip" form:"host_ip"`
	PodIP         string    `json:"pod_ip" form:"pod_ip"`
	ImageName     string    `json:"image_name" form:"image_name"`
	Ports         string    `json:"posts" form:"ports"`
	Resource      string    `json:"resource" form:"resource"`
	Environment   string    `json:"environment" form:"environment"`
	MountsVolumes string    `json:"mounts_volumes" form:"mounts_volumes"`
	StartTime     time.Time `json:"start_time" form:"start_time"`
	PodStatus     int       `json:"pod_status" form:"pod_status"`
	RestartCount  int       `json:"restart_count" form:"restart_count"`
	//DeployID      int       `json:"deploy_id" form:"deploy_id"`
	//ClusterID     int       `json:"cluster_id" form:"cluster_id"`
	CreateTime time.Time `json:"create_time" form:"create_time"`
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
