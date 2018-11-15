package query

import (
	"fmt"
	mydb "llgin/database/init"
	"llgin/structdb"
	"log"
)

var InfoDeployPool = make([]structdb.DeploymentsPool, 0)
var ListDeployPool = make([]structdb.DeploymentsPool, 0)

//GetClusterPoolInfo 获取一条数据的详细信息
func GetDeployPoolInfo(which int) (InfoDeployPool []structdb.DeploymentsPool) {

	var a structdb.DeploymentsPool
	db := mydb.InitDB()
	defer db.Close()

	//start := time.Now()
	rows, err := db.Query("SELECT * FROM deployments_pool where id = ?", which)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&a.ID,
			&a.Name,
			&a.Namespace,
			&a.Labels,
			&a.Version,
			&a.Selector,
			&a.Desired,
			&a.Availabel,
			&a.CreateTime); err != nil {
			log.Fatal(err)
		}

		fmt.Printf(" id: %d\n name: %s\n namespace: %s\n labels: %s\n version: %d\n selector: %s\n desired: %s\n availabels: %s\n create_time: %s\n",
			a.ID,
			a.Name,
			a.Namespace,
			a.Labels,
			a.Version,
			a.Selector,
			a.Desired,
			a.Availabel,
			a.CreateTime)
		InfoDeployPool = append(InfoDeployPool, a)
	}
	//end := time.Now()
	//fmt.Println("query total time:", end.Sub(start).Seconds())
	return
}

//GetClusterPoolList 获取集群信息列表
func GetDeployPoolList() (ListDeployPool []structdb.DeploymentsPool) {

	var a structdb.DeploymentsPool
	db := mydb.InitDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM deployments_pool")
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&a.ID,
			&a.ID,
			&a.Name,
			&a.Namespace,
			&a.Labels,
			&a.Version,
			&a.Selector,
			&a.Desired,
			&a.Availabel,
			&a.CreateTime); err != nil {
			log.Fatal(err)
		}

		fmt.Printf(" id: %d\n name: %s\n namespace: %s\n labels: %s\n version: %d\n selector: %d\n desired: %d\n availabels: %d\n create_time: %s\n",
			&a.ID,
			&a.Name,
			&a.Namespace,
			&a.Labels,
			&a.Version,
			&a.Selector,
			&a.Desired,
			&a.Availabel,
			a.CreateTime)
		ListDeployPool = append(ListDeployPool, a)
	}
	return
}
