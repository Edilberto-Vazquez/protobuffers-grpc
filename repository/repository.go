package repository

import (
	"context"

	"github.com/Edilberto-Vazquez/protobuffers-grpc/models"
)

type StudentRepository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
}

var implementation StudentRepository

func SetRepository(repository StudentRepository) {
	implementation = repository
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}
