package main

import (
	"github.com/labstack/echo"
	"github.com/r4bbitz/SayHello/api/v1/promotion/di"
	"github.com/r4bbitz/SayHello/api/v1/promotion/gateway/route"
)

func main() {
	e := echo.New()

	// Routes
	route.NewSayHelloRoute(di.ProvideSayHelloHandler()).Initial(e)

	// Listener
	e.Logger.Fatal(e.Start(":1991"))
}
