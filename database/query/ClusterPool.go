package query

import (
	"fmt"
	mydb "llgin/database/init"
	"llgin/structerr"
	"log"
)

//InfoClusterPool 返回切片信息
var InfoClusterPool = make([]structerr.ClusterPool, 0)

//ListClusterPool 返回数据以切片的方式
var ListClusterPool = make([]structerr.ClusterPool, 0)

//GetClusterPoolInfo 获取一条数据的详细信息
func GetClusterPoolInfo(ns string, ip string) (InfoClusterPool []structerr.ClusterPool) {

	var a structerr.ClusterPool
	db := mydb.InitDB()
	defer db.Close()

	//start := time.Now()
	rows, err := db.Query("SELECT * FROM cluster_pool where namespace = ? and cluster_ip = ?", ns, ip)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&a.ID,
			&a.Name,
			&a.Namespace,
			&a.ClusterIP,
			&a.ClusterPort,
			&a.ClusterToken,
			&a.ClusterInfo,
			&a.CreateTime); err != nil {
			log.Fatal(err)
		}

		// fmt.Printf(" id: %d\n name: %s\n namespace: %s\n clusterip: %s\n clusterport: %d\n clustertoken: %s\n clusterinfo: %s\n createtime: %s\n",
		// 	a.ID,
		// 	a.Name,
		// 	a.Namespace,
		// 	a.ClusterIP,
		// 	a.ClusterPort,
		// 	a.ClusterToken,
		// 	a.ClusterInfo,
		// 	a.CreateTime)
		InfoClusterPool = append(InfoClusterPool, a)
	}
	//end := time.Now()
	//fmt.Println("query total time:", end.Sub(start).Seconds())
	return
}

//GetClusterPoolList 获取集群信息列表
func GetClusterPoolList() (ListClusterPool []structerr.ClusterPool) {

	var a structerr.ClusterPool
	db := mydb.InitDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM cluster_pool")
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&a.ID, &a.Name, &a.Namespace, &a.ClusterIP, &a.ClusterPort, &a.ClusterToken, &a.ClusterInfo, &a.CreateTime); err != nil {
			log.Fatal(err)
		}

		// fmt.Printf(" id: %d\n name: %s\n namespace: %s\n clusterip: %s\n clusterport: %d\n clustertoken: %s\n clusterinfo: %s\n createtime: %s\n",
		// 	a.ID,
		// 	a.Name,
		// 	a.Namespace,
		// 	a.ClusterIP,
		// 	a.ClusterPort,
		// 	a.ClusterToken,
		// 	a.ClusterInfo,
		// 	a.CreateTime)
		ListClusterPool = append(ListClusterPool, a)
	}
	return
}
