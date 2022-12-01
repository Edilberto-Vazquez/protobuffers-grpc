package main

import (
	"log"
	"net"

	"github.com/Edilberto-Vazquez/protobuffers-grpc/database"
	"github.com/Edilberto-Vazquez/protobuffers-grpc/server"
	"github.com/Edilberto-Vazquez/protobuffers-grpc/studentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5060")

	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	server := server.NewStudentServer(repo)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)

	if err != s.Serve(list) {
		log.Fatal(err)
	}

}
