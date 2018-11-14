package main

import (
	get "llgin/restfulapi/getapi"
	post "llgin/restfulapi/postapi"
	put "llgin/restfulapi/putapi"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	post.Postapi(router)
	get.Getapi(router)
	put.Putapi(router)

	router.Run(":8000")
}
