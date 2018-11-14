package postapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Postapi 测试
func Postapi(router *gin.Engine) {

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "lilei")

		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"nick":    nick,
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
		})
	})

}

