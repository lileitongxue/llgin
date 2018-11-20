package query

import (
	"database/sql"
	"fmt"
	mydb "llgin/database/init"
	"llgin/structerr"
	"log"
)

//InfoDeployPool 返回切片类型的数据
var InfoDeployPool = make([]structerr.DeploymentsPool, 0)

//var ListDeployPool = make([]structerr.DeploymentsPool, 0)

//GetDeployPoolInfo 获取一条数据的详细信息
func GetDeployPoolInfo(ns string) (InfoDeployPool []structerr.DeploymentsPool) {

	var a structerr.DeploymentsPool
	db := mydb.InitDB()
	defer db.Close()

	//start := time.Now()
	rows, err := db.Query("SELECT name,labels,version,selector,desired,available,create_time FROM deployments_pool where namespace = ? ", ns)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			//&a.ID,
			//&a.DeployID,
			&a.Name,
			//&a.Namespace,
			&a.Labels,
			&a.Version,
			&a.Selector,
			&a.Desired,
			&a.Availabel,
			//&a.ClusterID,
			&a.CreateTime); err != nil {
			log.Fatal(err)
		}

		fmt.Printf(" name: %s\n labels: %s\n version: %s\n selector: %s\n desired: %d\n availabels: %d\n create_time: %s\n",
			//a.ID,
			//a.DeployID,
			a.Name,
			//a.Namespace,
			a.Labels,
			a.Version,
			a.Selector,
			a.Desired,
			a.Availabel,
			//a.ClusterID,
			a.CreateTime)
		InfoDeployPool = append(InfoDeployPool, a)
	}
	//end := time.Now()
	//fmt.Println("query total time:", end.Sub(start).Seconds())
	return
}

//GetDeployPoolList 获取集群信息列表
func GetDeployPoolList(ns string) []string {

	//var a structerr.DeploymentsPool
	db := mydb.InitDB()
	defer db.Close()
	rows, err := db.Query("SELECT name FROM deployments_pool where namespace = ?", ns)
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println("columns:", columns)
	values := make([]sql.RawBytes, len(columns))
	//fmt.Println("values:", values)
	scanArgs := make([]interface{}, len(values))
	//fmt.Println("scanArgs:", scanArgs)
	for i := range values {
		scanArgs[i] = &values[i]
		//fmt.Println("scanArgs:", scanArgs)
		//fmt.Println("values:", values)
	}
	//fmt.Println("scanArgs:", scanArgs)
	var value []string
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("log:", err)
			panic(err.Error())
		}
		for _, col := range values {
			//fmt.Println("col:", col)
			if col == nil {
				value = append(value, "NULL")
			} else {
				value = append(value, string(col))
			}
		}
	}
	//fmt.Println("scanArgs:", scanArgs)
	return value
}
