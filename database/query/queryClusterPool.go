package database

import (
	"fmt"
	sql "llgin/database/init"
	"llgin/structdb"
	"log"
)

var ResurtInfo = make([]structdb.ClusterPool, 0)
var ResurtList = make([]structdb.ClusterPool, 0)

//GetClusterPoolInfo 获取一条数据的详细信息
func GetClusterPoolInfo() (ResurtInfo []structdb.ClusterPool) {

	var a structdb.ClusterPool
	db := sql.InitDB()
	defer db.Close()

	//start := time.Now()

	rows, _ := db.Query("SELECT * FROM cluster_pool")
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&a.ID, &a.Name, &a.Namespace, &a.ClusterIP, &a.ClusterPort, &a.ClusterToken, &a.ClusterInfo, &a.CreateTime); err != nil {
			log.Fatal(err)
		}

		fmt.Printf(" id: %d\n name: %s\n namespace: %s\n clusterip: %s\n clusterport: %d\n clustertoken: %s\n clusterinfo: %s\n createtime: %s\n",
			a.ID,
			a.Name,
			a.Namespace,
			a.ClusterIP,
			a.ClusterPort,
			a.ClusterToken,
			a.ClusterInfo,
			a.CreateTime)
		ResurtInfo = append(ResurtInfo, a)
	}
	//end := time.Now()
	//fmt.Println("query total time:", end.Sub(start).Seconds())
	return
}

//GetClusterPoolList 获取集群信息列表
func GetClusterPoolList() (ResurtList []structdb.ClusterPool) {

	var a structdb.ClusterPool
	db := sql.InitDB()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM cluster_pool")
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&a.ID, &a.Name, &a.Namespace, &a.ClusterIP, &a.ClusterPort, &a.ClusterToken, &a.ClusterInfo, &a.CreateTime); err != nil {
			log.Fatal(err)
		}

		fmt.Printf(" id: %d\n name: %s\n namespace: %s\n clusterip: %s\n clusterport: %d\n clustertoken: %s\n clusterinfo: %s\n createtime: %s\n",
			a.ID,
			a.Name,
			a.Namespace,
			a.ClusterIP,
			a.ClusterPort,
			a.ClusterToken,
			a.ClusterInfo,
			a.CreateTime)
		ResurtList = append(ResurtList, a)
	}
	return
}
