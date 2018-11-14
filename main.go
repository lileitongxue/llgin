package main

import (
	restful "llgin/restfulapi"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	restful.Postapi(router)
	restful.Getapi(router)
	restful.Putapi(router)

	router.Run(":8000")
}
