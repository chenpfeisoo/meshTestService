package main

import (
	"meshTestService/pkg/route"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	gateway()
}
//入口服务
func gateway(){
	router := gin.Default()
	route.Route(router)
	router.GET("/", func(c *gin.Context) {
	})
	http.ListenAndServe(":8000", router)
}


