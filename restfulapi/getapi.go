package restfulapi

import (
	myquery "llgin/database/query"
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

//api接口提供pod_pool的信息,以id筛选
func GetPodInfo(c *gin.Context) {
	ns := c.DefaultQuery("ns", "gaojihealth")
	appname := c.Query("appname")
	if ns == "" {
		c.JSON(400, gin.H{
			"error": "ERROR: param is invaid,param shoud be ns",
		})
		return
	}
	if appname == "" {
		c.JSON(400, gin.H{
			"error": "ERROR: param is invaid,param shoud be appname ",
		})
		return
	}
	//whichInt, _ := strconv.Atoi(which)
	c.JSON(http.StatusOK, gin.H{
		"pod_pool": myquery.GetPodPoolInfo(appname, ns),
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
	ns := c.DefaultQuery("ns", "gaojihealth")
	//whichInt, _ := strconv.Atoi(which)
	appname := c.Query("appname")
	if ns == "" {
		c.JSON(400, gin.H{
			"error": "ERROR: param is invaid,param shoud be ns",
		})
		return
	}
	if appname == "" {
		c.JSON(400, gin.H{
			"error": "ERROR: param is invaid,param shoud be appname ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"deployments_pool": myquery.GetDeployPoolInfo(appname, ns),
	})
}

//api接口提供deployments_pool的全量信息
func GetDeploymentList(c *gin.Context) {
	ns := c.Query("ns")
	if ns == "" {
		c.JSON(400, gin.H{
			"error": "ERROR: param is invaid,param shoud be ns",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"deployments_pool": myquery.GetDeployPoolList(ns),
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
