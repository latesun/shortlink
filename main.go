package main

import (
	"github.com/latesun/shortlink/controllers"
	"github.com/latesun/shortlink/db"
)

func main() {
	db.Init()

	a := controllers.App{}
	a.Initialize()
	a.Run(":8899")
}
