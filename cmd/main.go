package main

import (
	"fmt"
	"log"
	"securities_admin/db"
	"securities_admin/internal/controller"
	"securities_admin/internal/service"
	route2 "securities_admin/route"
)

func main() {
	LoadConfig()
	err := db.ConnectDB(Config.Database.DNS())
	if err != nil {
		fmt.Printf("Failed to connect to database %s", err)
		return
	}

	logService := service.NewRequestService()
	logController := controller.NewLogController(logService)

	server := route2.NewAdminServer(logController)

	err = server.Router.Run(fmt.Sprintf(":%s", Config.Server.HTTPPort))
	if err != nil {
		log.Fatalf("Fail to run route %v", err)
	}
	//LoadConfig()
	//err := db.ConnectDB(Config.Database.DNS())
	//if err != nil {
	//	fmt.Printf("Failed to connect to database %s", err)
	//	return
	//}
	//
	//lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", Config.Server.Host, Config.Server.GRPCPort))
	//
	//if err != nil {
	//	log.Fatal("Fail to create route")
	//}
	//
	//requestService := service.NewRequestService()
	//s := grpc.NewServer()
	//api.RegisterAdminServiceServer(s, rpcimpl.NewAdminServer(requestService))
	//log.Printf("Server is running on port %s", Config.Server.GRPCPort)
	//
	//err = s.Serve(lis)
	//if err != nil {
	//	log.Fatal("Fail to create route")
	//}
}
