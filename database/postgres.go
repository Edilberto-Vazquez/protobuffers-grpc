package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/Edilberto-Vazquez/protobuffers-grpc/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db: db}, nil
}

func (p *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var student models.Student
	for rows.Next() {
		rows.Scan(&student.Id, &student.Name, &student.Age)
		return &student, err
	}
	return &student, nil
}

func (p *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := p.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)
	return err
}

func (p *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var test models.Test
	for rows.Next() {
		rows.Scan(&test.Id, &test.Name)
		return &test, err
	}
	return &test, nil
}

func (p *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := p.db.ExecContext(ctx, "INSERT INTO tests (id, name) VALUES ($1, $2)", test.Id, test.Name)
	return err
}
