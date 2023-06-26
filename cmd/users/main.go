package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/msik-404/micro-appoint-users/internal/database"
	"github.com/msik-404/micro-appoint-users/internal/userspb"
)

func main() {
	mongoClient, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	_, errors := database.CreateDBIndexes(mongoClient)
	for _, err := range errors {
		if err != nil {
			panic(err)
		}
	}
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	userspb.RegisterApiServer(s, &userspb.Server{Client: mongoClient})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
