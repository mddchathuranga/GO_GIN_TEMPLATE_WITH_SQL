package main

import (
	"github.com/user/test_template/db"
	"github.com/user/test_template/routers"
)

func main() {
	db.InitialDbConnection()
	routers.RunServer()
}
