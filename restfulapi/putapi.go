package restfulapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Putapi(router *gin.Engine) {

	router.PUT("/put", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
		c.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusOK,
		})
	})

}
