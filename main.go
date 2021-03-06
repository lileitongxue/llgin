package main

import (
	"llgin/getlogs"
	rf "llgin/restfulapi"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/api/v1/getNodeInfo", rf.GetNodeInfo)
	r.GET("/api/v1/getPodInfo", rf.GetPodInfo)
	r.GET("/api/v1/getAllNodes", rf.GetAllNodes)
	r.GET("/api/v1/getDeploymentInfo", rf.GetDeploymentInfo)
	r.GET("/api/v1/getDeploymentList", rf.GetDeploymentList)
	r.GET("/api/v1/getPodLog", getlogs.GetInstanceLogs)
	r.Run(":8080")
}
