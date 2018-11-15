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
		which := c.DefaultQuery("which", "1")
		whichInt, _ := strconv.Atoi(which)
		c.JSON(http.StatusOK, gin.H{
			"cluster_pool": myquery.GetClusterPoolInfo(whichInt),
		})
	})
	//api接口提供cluster_pool的全量信息
	router.GET("/getClusterPoolList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"cluster_pool": myquery.GetClusterPoolList(),
		})
	})
	//api接口提供pod_pool的信息,以id筛选
	router.GET("/getPodPoolInfo", func(c *gin.Context) {
		which := c.DefaultQuery("which", "1")
		whichInt, _ := strconv.Atoi(which)
		c.JSON(http.StatusOK, gin.H{
			"cluster_pool": myquery.GetPodPoolInfo(whichInt),
		})
	})
	//api接口提供pod_pool的全量信息
	router.GET("/getPodPoolList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"cluster_pool": myquery.GetPodPoolList(),
		})
	})
	//api接口提供deployments_pool的信息,以id筛选
	router.GET("/getDeployPoolInfo", func(c *gin.Context) {
		which := c.DefaultQuery("which", "1")
		whichInt, _ := strconv.Atoi(which)
		c.JSON(http.StatusOK, gin.H{
			"deployments_pool": myquery.GetDeployPoolInfo(whichInt),
		})
	})
	//api接口提供deployments_pool的全量信息
	router.GET("/getDeployPoolList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"deployments_pool": myquery.GetDeployPoolList(),
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
