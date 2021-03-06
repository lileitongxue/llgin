package restfulapi

import (
	myquery "llgin/database/query"
	sq "llgin/structerr"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NodeName struct {
	NodeName string `form:"nodename" binding:"required"`
}

//GetNodeInfo 接口提供hosts_pool的信息，以id筛选
func GetNodeInfo(c *gin.Context) {

	var Nn NodeName
	if bindErr := c.Bind(&Nn); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": sq.ParamBindErr{Err: bindErr.Error()}.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"node_info": myquery.GetNodeName(Nn.NodeName),
	})
}

//GetAllNodes 接口提供hosts_pool的全量name信息
func GetAllNodes(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"node_list": myquery.GetAllNodes(),
	})
}

//NsAppname 匹配required字段
type NsAppname struct {
	Ns      string `form:"ns" binding:"required"`
	AppName string `form:"appname" binding:"required"`
}

//GetPodInfo api接口提供pod_pool的信息,以id筛选
func GetPodInfo(c *gin.Context) {

	var Na NsAppname
	if bindErr := c.Bind(&Na); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": sq.ParamBindErr{Err: bindErr.Error()}.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pod_info": myquery.GetPodPoolInfo(Na.AppName, Na.Ns),
	})
}

//GetPodPoolList api接口提供pod_pool的全量信息
func GetPodPoolList(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"pod_pool": myquery.GetPodPoolList(),
	})
}

//GetDeploymentInfo api接口提供deployments_pool的信息,以id筛选
func GetDeploymentInfo(c *gin.Context) {

	var Ns Ns
	if bindErr := c.Bind(&Ns); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": sq.ParamBindErr{Err: bindErr.Error()}.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"deployments_info": myquery.GetDeployPoolInfo(Ns.Ns),
	})
}

//Ns 匹配required字段ns
type Ns struct {
	Ns string `form:"ns" binding:"required"`
}

//GetDeploymentList api接口提供deployments_pool的全量信息
func GetDeploymentList(c *gin.Context) {

	var ns Ns
	if bindErr := c.Bind(&ns); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": sq.ParamBindErr{Err: bindErr.Error()}.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"deployments_list": myquery.GetDeployPoolList(ns.Ns),
	})
}

//GetConfigmapsPoolInfo api接口提供configmaps_pool的信息,以id筛选
func GetConfigmapsPoolInfo(c *gin.Context) {

	which := c.DefaultQuery("which", "1")
	whichInt, _ := strconv.Atoi(which)
	c.JSON(http.StatusOK, gin.H{
		"configmaps_pool": myquery.GetConfigmapsPoolInfo(whichInt),
	})
}

//GetConfigmapsPoolList api接口提供configmaps_pool的全量信息
func GetConfigmapsPoolList(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"configmaps_pool": myquery.GetConfigmapsPoolList(),
	})
}
