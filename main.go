package main

import (
	"github.com/maximp14/golangreact/db"
	"github.com/maximp14/golangreact/handerls"
	"log"
)

func main() {
	if db.CheckConnection() == false{
		log.Fatal("Data base connection fail")
	}
	handerls.Handlers()
}
