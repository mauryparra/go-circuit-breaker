package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mauryparra/go-circuit-breaker/src/go-circuit-breaker/controllers/cbmiddle"
	"github.com/mauryparra/go-circuit-breaker/src/go-circuit-breaker/controllers/ping"
)

const (
	port = ":8181"
)

var (
	router = gin.Default()
)

func main() {

	router.GET("/ping", ping.Ping)
	router.GET("/cbmiddle", cbmiddle.Cb)
	router.Run(port)

}
