package query

import (
	"database/sql"
	"fmt"
	mydb "llgin/database/init"
	"llgin/structerr"
	"log"
)

//GetAllNodes获取集群信息列表
func GetAllNodes() []string {

	db := mydb.InitDB()
	defer db.Close()
	rows, err := db.Query("SELECT host_name FROM hosts_pool")
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	var value []string
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("log:", err)
			panic(err.Error())
		}
		for _, col := range values {
			if col == nil {
				value = append(value, "NULL")
			} else {
				value = append(value, string(col))
			}
		}
	}
	return value
}

var NodeName = make([]structerr.HostsPool, 0)

func GetNodeName(nodename string) (NodeName []structerr.HostsPool) {

	var a structerr.HostsPool
	db := mydb.InitDB()
	defer db.Close()
	//start := time.Now()
	rows, err := db.Query("SELECT labels,host_ip,host_name,host_status,os,kubelet_version,kubelet_proxy_version,host_network,host_cpu,host_memory,host_used_cpu,host_used_memory,running_pods,create_time FROM hosts_pool where host_name = ? ", nodename)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			//&a.ID,
			&a.Labels,
			&a.HostIP,
			&a.HostName,
			&a.HostStatus,
			&a.OS,
			&a.KubeletVersion,
			&a.KubeletProxyVersion,
			&a.HostNetwork,
			&a.HostCPU,
			&a.HostMemory,
			&a.HostUsedCPU,
			&a.HostUsedMemory,
			&a.RunningPod,
			//&a.ClusterID,
			&a.CreateTime); err != nil {
			log.Fatal(err)
		}

		// fmt.Printf(" name: %s\n labels: %s\n version: %s\n selector: %s\n desired: %d\n availabels: %d\n create_time: %s\n",
		// 	//a.ID,
		// 	//a.DeployID,
		// 	a.Name,
		// 	//a.Namespace,
		// 	a.Labels,
		// 	a.Version,
		// 	a.Selector,
		// 	a.Desired,
		// 	a.Availabel,
		// 	//a.ClusterID,
		// 	a.CreateTime)
		NodeName = append(NodeName, a)
	}
	//end := time.Now()
	//fmt.Println("query total time:", end.Sub(start).Seconds())
	return
}
