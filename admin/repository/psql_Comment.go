package repository

import (
	"database/sql"
	"errors"

	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// CommentRepositoryImpl implements the menu.CommentRepository interface
type CommentRepositoryImpl struct {
	conn *sql.DB
}

// NewCommentRepositoryImpl will create an object of PsqlCommentRepository
func NewCommentRepositoryImpl(Conn *sql.DB) *CommentRepositoryImpl {
	return &CommentRepositoryImpl{conn: Conn}
}

// Comments returns all Comments from the database
func (cri *CommentRepositoryImpl) Comments() ([]entity.Comment, error) {

	rows, err := cri.conn.Query("SELECT * FROM comments;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Comment{}

	for rows.Next() {
		Comment := entity.Comment{}

		err = rows.Scan(&Comment.ID, &Comment.UserName, &Comment.Email, &Comment.Message, &Comment.PlacedAt)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, Comment)
	}

	return ctgs, nil
}

// Comment returns a Comment with a given id
func (cri *CommentRepositoryImpl) Comment(id int) (entity.Comment, error) {

	row := cri.conn.QueryRow("SELECT * FROM comments WHERE id = $1", id)

	Comment := entity.Comment{}

	err := row.Scan(&Comment.ID, &Comment.UserName, &Comment.Email, &Comment.Message, &Comment.PlacedAt)
	if err != nil {
		return Comment, err
	}

	return Comment, nil
}

// UpdateComment updates a given object with a new data
func (cri *CommentRepositoryImpl) UpdateComment(c entity.Comment) error {

	_, err := cri.conn.Exec("UPDATE comments SET username=$1,email=$2,messages=$3,placedat=$4 where id=$5", c.UserName, c.Email, c.Message, c.PlacedAt, c.ID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteComment removes a Comment from a database by its id
func (cri *CommentRepositoryImpl) DeleteComment(id int) error {

	_, err := cri.conn.Exec("DELETE FROM comments WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreComment stores new Comment information to database
func (cri *CommentRepositoryImpl) StoreComment(c entity.Comment) error {

	_, err := cri.conn.Exec("INSERT INTO comments (username,email,messages,placedat) values($1, $2, $3,$4)", c.UserName, c.Email, c.Message, c.PlacedAt)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
