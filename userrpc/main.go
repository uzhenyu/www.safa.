package main

import (
	"flag"
	grpc2 "google.golang.org/grpc"
	"zg5/z304/framework/app"
	"zg5/z304/framework/grpc"
	"zg5/z304/userrpc/api"
	"zg5/z304/userrpc/model"
)

var (
	port = flag.Int("port", 8081, "wzy")
)

func main() {
	flag.Parse()
	err := app.Init("DEFAULT_GROUP", "mysql")
	if err != nil {
		panic(err)
	}

	err = app.Init("DEFAULT_GROUP", "redis")
	if err != nil {
		panic(err)
	}

	err = model.AutoTable()
	if err != nil {
		panic(err)
	}

	err = grpc.GetGrpc(*port, func(s *grpc2.Server) {
		api.Reg(s)
	})
	if err != nil {
		panic(err)
	}
}
