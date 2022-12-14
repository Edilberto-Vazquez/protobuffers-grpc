package studentserver

import (
	"context"

	"github.com/Edilberto-Vazquez/protobuffers-grpc/models"
	"github.com/Edilberto-Vazquez/protobuffers-grpc/repository"
	"github.com/Edilberto-Vazquez/protobuffers-grpc/studentpb"
)

type StudentServer struct {
	repo repository.StudentRepository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.StudentRepository) *StudentServer {
	return &StudentServer{repo: repo}
}

func (s *StudentServer) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetStudent(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *StudentServer) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	student := models.Student{
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}
	err := s.repo.SetStudent(ctx, &student)
	if err != nil {
		return nil, err
	}
	return &studentpb.SetStudentResponse{Id: student.Id}, nil
}
