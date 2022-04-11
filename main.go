package main

import (
	"github.com/gin-gonic/gin"
	"library/auth"
	"library/db"
	"library/router"
	"library/utils/crypto"
	"log"
)

func main() {
	r := gin.Default()
	db.InitCaller()
	auth.InitAuth()     //init auth
	crypto.InitCrypto() // init crypto
	router.Register(r)  // init route
	// listen to 0.0.0.0:5000
	log.Fatal(r.Run(":5000"))
}
