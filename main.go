package main

import (
	"go-api/app/bootstrap"
	"go-api/app/utils"
)

func main() {
	utils.InitLogger()
	bootstrap.RunServer()
}
