package service

import (
	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// CommentServiceImpl implements menu.CommentService interface
type CommentServiceImpl struct {
	CommentRepo admin.CommentRepository
}

// NewCommentServiceImpl will create new CommentService object
func NewCommentServiceImpl(CatRepo admin.CommentRepository) *CommentServiceImpl {
	return &CommentServiceImpl{CommentRepo: CatRepo}
}

// Comments ..() returns list of Comments
func (cs *CommentServiceImpl) Comments() ([]entity.Comment, error) {

	comments, err := cs.CommentRepo.Comments()

	if err != nil {
		return nil, err
	}

	return comments, nil
}

// StoreComment persists new Comment information
func (cs *CommentServiceImpl) StoreComment(Comment entity.Comment) error {

	err := cs.CommentRepo.StoreComment(Comment)

	if err != nil {
		return err
	}

	return nil
}

// Comment returns a Comment object with a given id
func (cs *CommentServiceImpl) Comment(id int) (entity.Comment, error) {

	c, err := cs.CommentRepo.Comment(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateComment updates a cateogory with new data
func (cs *CommentServiceImpl) UpdateComment(Comment entity.Comment) error {

	err := cs.CommentRepo.UpdateComment(Comment)

	if err != nil {
		return err
	}

	return nil
}

// DeleteComment delete a Comment by its id
func (cs *CommentServiceImpl) DeleteComment(id int) error {

	err := cs.CommentRepo.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil
}
