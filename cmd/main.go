package main

import (
	"readSoftware/api"
	"readSoftware/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
