package main

import (
	rf "llgin/restfulapi"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	rf.Postapi(r)
	rf.Getapi(r)
	rf.Putapi(r)

	r.Run(":8000")
}
