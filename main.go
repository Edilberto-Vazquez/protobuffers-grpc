package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/Edilberto-Vazquez/protobuffers-grpc/database"
	"github.com/Edilberto-Vazquez/protobuffers-grpc/studentpb"
	"github.com/Edilberto-Vazquez/protobuffers-grpc/studentserver"
	"github.com/Edilberto-Vazquez/protobuffers-grpc/testpb"
	"github.com/Edilberto-Vazquez/protobuffers-grpc/testserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	var wg sync.WaitGroup
	wg.Add(2)
	// student service
	go func() {
		defer wg.Done()
		_, cancel := context.WithCancel(ctx)
		defer cancel()

		studentListener, err := net.Listen("tcp", ":5060")

		if err != nil {
			log.Fatal(err)
		}

		studentServer := studentserver.NewStudentServer(repo)

		studentGrpcServer := grpc.NewServer()
		studentpb.RegisterStudentServiceServer(studentGrpcServer, studentServer)

		reflection.Register(studentGrpcServer)

		if err != studentGrpcServer.Serve(studentListener) {
			log.Fatal(err)
			cancel()
		}
	}()

	//test service
	go func() {
		defer wg.Done()
		_, cancel := context.WithCancel(ctx)
		defer cancel()

		testListener, err := net.Listen("tcp", ":5070")

		if err != nil {
			log.Fatal(err)
		}

		testServer := testserver.NewTestServer(repo)

		testGrpcServer := grpc.NewServer()
		testpb.RegisterTestServiceServer(testGrpcServer, testServer)

		reflection.Register(testGrpcServer)

		if err != testGrpcServer.Serve(testListener) {
			log.Fatal(err)
			cancel()
		}
	}()

	wg.Wait()
}
