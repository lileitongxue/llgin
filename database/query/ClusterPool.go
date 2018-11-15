package query

import (
	"fmt"
	mydb "llgin/database/init"
	"llgin/structdb"
	"log"
)

var InfoClusterPool = make([]structdb.ClusterPool, 0)
var ListClusterPool = make([]structdb.ClusterPool, 0)

//GetClusterPoolInfo 获取一条数据的详细信息
func GetClusterPoolInfo(which int) (InfoClusterPool []structdb.ClusterPool) {

	var a structdb.ClusterPool
	db := mydb.InitDB()
	defer db.Close()

	//start := time.Now()
	rows, err := db.Query("SELECT id,name FROM cluster_pool where id = ?", which)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&a.ID,
			&a.Name,
			&a.Namespace,
			// &a.ClusterIP,
			// &a.ClusterPort,
			// &a.ClusterToken,
			// &a.ClusterInfo,
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
func GetClusterPoolList() (ListClusterPool []structdb.ClusterPool) {

	var a structdb.ClusterPool
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

		fmt.Printf(" id: %d\n name: %s\n namespace: %s\n clusterip: %s\n clusterport: %d\n clustertoken: %s\n clusterinfo: %s\n createtime: %s\n",
			a.ID,
			a.Name,
			a.Namespace,
			a.ClusterIP,
			a.ClusterPort,
			a.ClusterToken,
			a.ClusterInfo,
			a.CreateTime)
		ListClusterPool = append(ListClusterPool, a)
	}
	return
}
