package getapi

import (
	sql "llgin/database/query"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Getapi 测试
func Getapi(router *gin.Engine) {

	router.GET("/getClusterPoolInfo", func(c *gin.Context) {
		which := c.DefaultQuery("which", "1")
		whichInt, _ := strconv.Atoi(which)
		c.JSON(http.StatusOK, gin.H{
			"cluster_pool": sql.GetClusterPoolInfo(whichInt),
		})
	})

	router.GET("/getClusterPoolList", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"cluster_pool": sql.GetClusterPoolList(),
		})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s", name)
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

}
