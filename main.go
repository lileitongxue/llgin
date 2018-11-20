package main

import (
	rf "llgin/restfulapi"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	//r.GET("/getClusterPoolInfo", rf.GetClusterPoolInfo)
	r.GET("/api/v1/getPodInfo", rf.GetPodInfo)
	r.GET("/api/v1/getAllNodes", rf.GetAllNodes)
	r.GET("/api/v1/getDeploymentInfo", rf.GetDeploymentInfo)
	r.GET("/api/v1/getDeploymentList", rf.GetDeploymentList)
	r.Run(":8000")
}
