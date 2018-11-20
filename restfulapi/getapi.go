package restfulapi

import (
	myquery "llgin/database/query"
	sq "llgin/structerr"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//api接口提供cluster_pool的信息，以id筛选
func GetClusterPoolInfo(c *gin.Context) {

	ns := c.Query("ns")
	ip := c.Query("ip")
	if ns == "" {
		c.JSON(400, gin.H{
			"error": "ERROR: param is invaid,param shoud be ns",
		})
		return
	}
	if ip == "" {
		c.JSON(400, gin.H{
			"error": "ERROR: param is invaid,param shoud be ip ",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cluster_pool": myquery.GetClusterPoolInfo(ns, ip),
	})
}

//api接口提供cluster_pool的全量信息
func GetClusterPoolList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"cluster_pool": myquery.GetClusterPoolList(),
	})
}

type NsAppname struct {
	Ns      string `form:"ns" binding:"required"`
	AppName string `form:"appname" binding:"required"`
}

//api接口提供pod_pool的信息,以id筛选
func GetPodInfo(c *gin.Context) {
	var Na NsAppname
	if bindErr := c.Bind(&Na); bindErr != nil {
		c.JSON(400, gin.H{
			"error": sq.ParamBindErr{Err: bindErr.Error()}.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pod_pool": myquery.GetPodPoolInfo(Na.AppName, Na.Ns),
	})
}

//api接口提供pod_pool的全量信息
func GetPodPoolList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pod_pool": myquery.GetPodPoolList(),
	})
}

//api接口提供deployments_pool的信息,以id筛选
func GetDeploymentInfo(c *gin.Context) {
	var Na NsAppname
	if bindErr := c.Bind(&Na); bindErr != nil {
		c.JSON(400, gin.H{
			"error": sq.ParamBindErr{Err: bindErr.Error()}.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"deployments_pool": myquery.GetDeployPoolInfo(Na.AppName, Na.Ns),
	})
}

type Ns struct {
	Ns string `form:"ns" binding:"required"`
}

//api接口提供deployments_pool的全量信息
func GetDeploymentList(c *gin.Context) {
	var ns Ns
	if bindErr := c.Bind(&ns); bindErr != nil {
		c.JSON(400, gin.H{
			"error": sq.ParamBindErr{Err: bindErr.Error()}.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"deployments_pool": myquery.GetDeployPoolList(ns.Ns),
	})
}

//api接口提供configmaps_pool的信息,以id筛选
func GetConfigmapsPoolInfo(c *gin.Context) {
	which := c.DefaultQuery("which", "1")
	whichInt, _ := strconv.Atoi(which)
	c.JSON(http.StatusOK, gin.H{
		"configmaps_pool": myquery.GetConfigmapsPoolInfo(whichInt),
	})
}

//api接口提供configmaps_pool的全量信息
func GetConfigmapsPoolList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"configmaps_pool": myquery.GetConfigmapsPoolList(),
	})
}
