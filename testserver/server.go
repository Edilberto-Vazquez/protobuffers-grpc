package testserver

import (
	"context"

	"github.com/Edilberto-Vazquez/protobuffers-grpc/models"
	"github.com/Edilberto-Vazquez/protobuffers-grpc/repository"
	"github.com/Edilberto-Vazquez/protobuffers-grpc/testpb"
)

type TestServer struct {
	repo repository.StudentRepository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.StudentRepository) *TestServer {
	return &TestServer{repo: repo}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	test := models.Test{
		Id:   req.GetId(),
		Name: req.GetName(),
	}
	err := s.repo.SetTest(ctx, &test)
	if err != nil {
		return nil, err
	}
	return &testpb.SetTestResponse{Id: test.Id}, nil
}
