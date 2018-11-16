package restfulapi

import (
	myquery "llgin/database/query"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Getapi
func Getapi(router *gin.Engine) {
	//api接口提供cluster_pool的信息，以id筛选
	router.GET("/getClusterPoolInfo", func(c *gin.Context) {
		ns := c.DefaultQuery("ns", "gaojihealth")
		//whichInt, _ := strconv.Atoi(which)
		ip := c.Query("ip")
		c.JSON(http.StatusOK, gin.H{
			"cluster_pool": myquery.GetClusterPoolInfo(ns, ip),
		})
	})
	//api接口提供cluster_pool的全量信息
	router.GET("/getClusterPoolList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"cluster_pool": myquery.GetClusterPoolList(),
		})
	})
	//api接口提供pod_pool的信息,以id筛选
	router.GET("/getPodInfo", func(c *gin.Context) {
		ns := c.DefaultQuery("ns", "gaojihealth")
		appname := c.Query("appname")
		//whichInt, _ := strconv.Atoi(which)
		c.JSON(http.StatusOK, gin.H{
			"pod_pool": myquery.GetPodPoolInfo(appname, ns),
		})
	})
	//api接口提供pod_pool的全量信息
	// router.GET("/getPodPoolList", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"pod_pool": myquery.GetPodPoolList(),
	// 	})
	// })
	//api接口提供deployments_pool的信息,以id筛选
	router.GET("/getDeploymentInfo", func(c *gin.Context) {
		ns := c.DefaultQuery("ns", "gaojihealth")
		//whichInt, _ := strconv.Atoi(which)
		appname := c.Query("appname")
		c.JSON(http.StatusOK, gin.H{
			"deployments_pool": myquery.GetDeployPoolInfo(appname, ns),
		})
	})
	//api接口提供deployments_pool的全量信息
	router.GET("/getDeploymentList", func(c *gin.Context) {
		ns := c.Query("gaojihealth")
		c.JSON(http.StatusOK, gin.H{
			"deployments_pool": myquery.GetDeployPoolList(ns),
		})
	})
	//api接口提供configmaps_pool的信息,以id筛选
	router.GET("/getConfigmapsPoolInfo", func(c *gin.Context) {
		which := c.DefaultQuery("which", "1")
		whichInt, _ := strconv.Atoi(which)
		c.JSON(http.StatusOK, gin.H{
			"configmaps_pool": myquery.GetConfigmapsPoolInfo(whichInt),
		})
	})
	//api接口提供configmaps_pool的全量信息
	router.GET("/getConfigmapsPoolList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"configmaps_pool": myquery.GetConfigmapsPoolList(),
		})
	})
}
