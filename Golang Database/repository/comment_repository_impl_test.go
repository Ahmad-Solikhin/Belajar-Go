package repository

import (
	Goalng_Database "belajar_go_database"
	"belajar_go_database/entity"
	"context"
	"fmt"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(Goalng_Database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{Comment: "Test dari repository"}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(Goalng_Database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepositoryImpl(Goalng_Database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
