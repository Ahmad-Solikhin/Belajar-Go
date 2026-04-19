package repository

import (
	"belajar_go_database/entity"
	"context"
	"database/sql"
	"errors"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepositoryImpl(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "insert into comment(comment) values($1) returning id"
	err := repository.DB.QueryRowContext(ctx, script, comment.Comment).Scan(&comment.Id)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "select * from comment where id = $1"
	rows, err := repository.DB.QueryContext(ctx, script, id)

	comment := entity.Comment{}

	if err != nil {
		return comment, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "select * from comment"
	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
