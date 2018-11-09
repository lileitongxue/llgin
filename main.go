package main

import (
	get "llgin/getapi"
	post "llgin/postapi"
	put "llgin/putapi"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	post.Postapi(router)
	get.Getapi(router)
	put.Putapi(router)

	router.Run(":8000")
}
