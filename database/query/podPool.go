package query

import (
	"fmt"
	mydb "llgin/database/init"
	"llgin/structerr"
	"log"
)

var InfoPodPool = make([]structerr.PodPool, 0)
var ListPodPool = make([]structerr.PodPool, 0)

//GetPodPoolInfo 获取一条数据的详细信息
func GetPodPoolInfo(appname string, ns string) (InfoPodPool []structerr.PodPool) {

	var a structerr.PodPool
	db := mydb.InitDB()
	defer db.Close()

	//start := time.Now()
	rows, _ := db.Query("SELECT * FROM pod_pool where  name = ? and namespace = ?", appname, ns)
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&a.ID,
			&a.Name,
			&a.Namespace,
			&a.Version,
			&a.HostIP,
			&a.PodIP,
			&a.ImageName,
			&a.Ports,
			&a.Resource,
			&a.Environment,
			&a.MountsVolumes,
			&a.StartTime,
			&a.PodStatus,
			&a.RestartCount,
			&a.ClusterID,
			&a.CreateTime); err != nil {
			log.Fatal(err)
		}

		fmt.Printf(" id: %d\n name: %s\n namespace: %s\n version: %s\n hostIP: %s\n podIP: %s\n imagename: %s\n ports: %s\n resource: %s\n environment: %s\n mountsVolumes: %s\n startTime: %s\n podStatus: %d\n restartCount %d\n clusterID: %d\n createTime: %s\n",
			a.ID,
			a.Name,
			a.Namespace,
			a.Version,
			a.HostIP,
			a.PodIP,
			a.ImageName,
			a.Ports,
			a.Resource,
			a.Environment,
			a.MountsVolumes,
			a.StartTime,
			a.PodStatus,
			a.RestartCount,
			a.ClusterID,
			a.CreateTime)
		InfoPodPool = append(InfoPodPool, a)
	}
	//end := time.Now()
	//fmt.Println("query total time:", end.Sub(start).Seconds())
	return
}

//GetPodPoolList 获取集群信息列表
func GetPodPoolList() (ListPodPool []structerr.PodPool) {

	var a structerr.PodPool
	db := mydb.InitDB()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM pod_pool")
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&a.ID,
			&a.Name,
			&a.Namespace,
			&a.Version,
			&a.HostIP,
			&a.PodIP,
			&a.ImageName,
			&a.Ports,
			&a.Resource,
			&a.Environment,
			&a.MountsVolumes,
			&a.StartTime,
			&a.PodStatus,
			&a.RestartCount,
			&a.ClusterID,
			&a.CreateTime); err != nil {
			log.Fatal(err)
		}

		fmt.Printf(" id: %d\n name: %s\n namespace: %s\n version: %s\n hostIP: %s\n podIP: %s\n imagename: %s\n ports: %s\n resource: %s\n environment: %s\n mountsVolumes: %s\n startTime: %s\n podStatus: %d\n restartCount %d\n clusterID: %d\n createTime: %s\n",
			a.ID,
			a.Name,
			a.Namespace,
			a.Version,
			a.HostIP,
			a.PodIP,
			a.ImageName,
			a.Ports,
			a.Resource,
			a.Environment,
			a.MountsVolumes,
			a.StartTime,
			a.PodStatus,
			a.RestartCount,
			a.ClusterID,
			a.CreateTime)
		ListPodPool = append(ListPodPool, a)
	}
	return
}
