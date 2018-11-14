package restfulapi

import (
	myquery "llgin/database/query"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Getapi
func Getapi(router *gin.Engine) {

	router.GET("/getClusterPoolInfo", func(c *gin.Context) {
		which := c.DefaultQuery("which", "1")
		whichInt, _ := strconv.Atoi(which)
		c.JSON(http.StatusOK, gin.H{
			"cluster_pool": myquery.GetClusterPoolInfo(whichInt),
		})
	})

	router.GET("/getClusterPoolList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"cluster_pool": myquery.GetClusterPoolList(),
		})
	})

	router.GET("/getPodPoolInfo", func(c *gin.Context) {
		which := c.DefaultQuery("which", "1")
		whichInt, _ := strconv.Atoi(which)
		c.JSON(http.StatusOK, gin.H{
			"cluster_pool": myquery.GetPodPoolInfo(whichInt),
		})
	})

	router.GET("/getPodPoolList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"cluster_pool": myquery.GetPodPoolList(),
		})
	})

}
