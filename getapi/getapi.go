package getapi

import (
	sql "llgin/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Getapi 测试
func Getapi(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
		sql.Query1()
	})

	router.GET("/lilei/", func(c *gin.Context) {
		c.String(http.StatusOK, "this is a shuaib")
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
