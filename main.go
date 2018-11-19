package main

import (
	rf "llgin/restfulapi"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	//r.GET("/getClusterPoolInfo", rf.GetClusterPoolInfo)
	r.GET("/getPodInfo", rf.GetPodInfo)
	r.GET("/getPodPoolList", rf.GetPodPoolList)
	//r.GET("/getDeploymentInfo", rf.GetDeploymentInfo)
	r.GET("/getDeploymentList", rf.GetDeploymentList)
	r.Run(":8000")
}
