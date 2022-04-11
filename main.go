package main

import (
	"myapp/routes"
)

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
