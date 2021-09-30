package server

import (
	"code_sim/config"
	"code_sim/pb_gen"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartServe() {
	appName := config.Conf.AppName
	addr := fmt.Sprintf("%s:%d", config.Conf.Host, config.Conf.Port)
	log.Printf("%s Dialing addr: %s", appName, addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb_gen.RegisterCodeSimServer(grpcServer, newCodeSimServer())
	log.Printf("%s ready to server at %s...", appName, addr)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer Serve failed, err=[%v]", err)
	}
}
