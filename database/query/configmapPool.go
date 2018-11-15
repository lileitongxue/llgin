package query

import (
	"fmt"
	mydb "llgin/database/init"
	"llgin/structdb"
	"log"
)

var InfoConfigmapsPool = make([]structdb.ConfigmapsPool, 0)
var ListConfigmapsPool = make([]structdb.ConfigmapsPool, 0)

//GetConfigmapsPoolInfo 获取一条数据的详细信息
func GetConfigmapsPoolInfo(which int) (InfoConfigmapsPool []structdb.ConfigmapsPool) {

	var a structdb.ConfigmapsPool
	db := mydb.InitDB()
	defer db.Close()

	//start := time.Now()
	rows, err := db.Query("SELECT * FROM configmaps_pool where id = ?", which)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&a.ID,
			&a.TmplName,
			&a.Namespace,
			&a.Labels,
			&a.Data,
			&a.URL,
			&a.HostIPPorts,
			&a.ConfigmapName,
			&a.ClusterID,
			&a.CreateTime); err != nil {
			log.Fatal(err)
		}

		fmt.Printf(" id: %d\n tmplname: %s\n namespace: %s\n labels: %s\n data: %s\n url: %s\n host_ip_ports: %s\n configmap_name: %s\n cluster_id: %d\n create_time: %s\n",
			a.ID,
			a.TmplName,
			a.Namespace,
			a.Labels,
			a.Data,
			a.URL,
			a.HostIPPorts,
			a.ConfigmapName,
			a.ClusterID,
			a.CreateTime)
		InfoConfigmapsPool = append(InfoConfigmapsPool, a)
	}
	//end := time.Now()
	//fmt.Println("query total time:", end.Sub(start).Seconds())
	return
}

//GetConfigmapsPoolList 获取集群信息列表
func GetConfigmapsPoolList() (ListConfigmapsPool []structdb.ConfigmapsPool) {

	var a structdb.ConfigmapsPool
	db := mydb.InitDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM configmaps_pool")
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&a.ID,
			&a.TmplName,
			&a.Namespace,
			&a.Labels,
			&a.Data,
			&a.URL,
			&a.HostIPPorts,
			&a.ConfigmapName,
			&a.ClusterID,
			&a.CreateTime); err != nil {
			log.Fatal(err)
		}

		fmt.Printf(" id: %d\n tmplname: %s\n namespace: %s\n labels: %s\n data: %s\n url: %s\n host_ip_ports: %s\n configmap_name: %s\n cluster_id: %d\n create_time: %s\n",
			a.ID,
			a.TmplName,
			a.Namespace,
			a.Labels,
			a.Data,
			a.URL,
			a.HostIPPorts,
			a.ConfigmapName,
			a.ClusterID,
			a.CreateTime)
		ListConfigmapsPool = append(ListConfigmapsPool, a)
	}
	return
}