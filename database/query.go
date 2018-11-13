package database

import (
	"fmt"
	"llgin/structdb"
	"log"
)

func Query1() {
	var a structdb.ClusterPool
	db := InitDB()
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
	}
	//end := time.Now()
	//fmt.Println("query total time:", end.Sub(start).Seconds())
}
