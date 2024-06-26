package service

import (
	"context"
	pb "forum-service/forum-protos/genprotos"
	st "forum-service/storage"

	"github.com/google/uuid"
)

type CommentService struct {
	storage st.Storage
	pb.UnimplementedCommentServiceServer
}

func NewCommentService(storage *st.Storage) *CommentService {
	return &CommentService{storage: *storage}
}

func (s *CommentService) Create(ctx context.Context, comment *pb.CommentCReqOrCResOrGResOrURes) (*pb.CommentCReqOrCResOrGResOrURes, error) {
	comment.CommentId = uuid.NewString()
	resp, err := s.storage.CommentS.Create(comment)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CommentService) GetByID(ctx context.Context, idReq *pb.CommentGReqOrDReq) (*pb.CommentCReqOrCResOrGResOrURes, error) {
	resp, err := s.storage.CommentS.GetByID(idReq)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CommentService) GetAll(ctx context.Context, allComments *pb.CommentGAReq) (*pb.CommentGARes, error) {
	comments, err := s.storage.CommentS.GetAll(allComments)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *CommentService) Update(ctx context.Context, comment *pb.CommentUReq) (*pb.CommentCReqOrCResOrGResOrURes, error) {
	resp, err := s.storage.CommentS.Update(comment)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *CommentService) Delete(ctx context.Context, idReq *pb.CommentGReqOrDReq) (*pb.Void, error) {
	tx, err := s.storage.Db.Begin()
	if err != nil {
		return nil, err
	}
	_, err = s.storage.CommentS.Delete(tx, idReq)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &pb.Void{}, nil
}
