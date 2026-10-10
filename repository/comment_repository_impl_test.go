package repository

import (
	Belajar_Database "Belajar-Database"
	"Belajar-Database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInser(t *testing.T) {
	commentRepository := NewCommentRepository(Belajar_Database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "contoh@test.com",
		Comment: "admin123",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(Belajar_Database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 21)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(Belajar_Database.GetConnection())
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
