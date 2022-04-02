package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"library/db"
	"library/router"
	"log"
)

func main() {
	r := gin.Default()
	db.InitCaller()
	router.Register(r) // init route
	logrus.WithField("port", "5000").Error("server start")

	// listen to 0.0.0.0:5000
	log.Fatal(r.Run(":5000"))
}
