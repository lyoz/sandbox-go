package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lyoz/sandbox-go/router"
)

func main() {
	e := router.Init()
	e.Logger.Fatal(e.Start(":3000"))
}
