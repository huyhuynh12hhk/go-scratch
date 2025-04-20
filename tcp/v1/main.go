package main

import (
	"lab/goratch/tcp/v1/servers"
)

func main() {
	s := servers.NewServer(42231)
	s.Router.AddRoute("/hello", helloWorld)
	
	s.Run()
}


func helloWorld(ctx *servers.HttpContext) {
	// msg:= "Hello, world"

	ctx.WriteSimpleResponse()

}
