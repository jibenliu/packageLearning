// Code generated by sqlc-grpc (https://github.com/walterwanderley/sqlc-grpc). DO NOT EDIT.

package books

import (
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"

	pb "booktest/api/books/v1"
	"booktest/internal/validation"
)

type Service struct {
	pb.UnimplementedBooksServiceServer
	logger  *zap.Logger
	querier *Queries
}

func (s *Service) BooksByTags(ctx context.Context, req *pb.BooksByTagsRequest) (*pb.BooksByTagsResponse, error) {
	dollar_1 := req.GetDollar_1()

	result, err := s.querier.BooksByTags(ctx, dollar_1)
	if err != nil {
		s.logger.Error("BooksByTags sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.BooksByTagsResponse)
	for _, r := range result {
		res.List = append(res.List, toBooksByTagsRow(r))
	}
	return res, nil
}

func (s *Service) BooksByTitleYear(ctx context.Context, req *pb.BooksByTitleYearRequest) (*pb.BooksByTitleYearResponse, error) {
	var arg BooksByTitleYearParams
	arg.Title = req.GetTitle()
	arg.Year = req.GetYear()

	result, err := s.querier.BooksByTitleYear(ctx, arg)
	if err != nil {
		s.logger.Error("BooksByTitleYear sql call failed", zap.Error(err))
		return nil, err
	}
	res := new(pb.BooksByTitleYearResponse)
	for _, r := range result {
		res.List = append(res.List, toBook(r))
	}
	return res, nil
}

func (s *Service) CreateAuthor(ctx context.Context, req *pb.CreateAuthorRequest) (*pb.CreateAuthorResponse, error) {
	name := req.GetName()

	result, err := s.querier.CreateAuthor(ctx, name)
	if err != nil {
		s.logger.Error("CreateAuthor sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreateAuthorResponse{Author: toAuthor(result)}, nil
}

func (s *Service) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	var arg CreateBookParams
	arg.AuthorID = req.GetAuthorId()
	arg.Isbn = req.GetIsbn()
	arg.BookType = BookType(req.GetBookType())
	arg.Title = req.GetTitle()
	arg.Year = req.GetYear()
	if v := req.GetAvailable(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid Available: %s%w", err.Error(), validation.ErrUserInput)
			return nil, err
		}
		arg.Available = v.AsTime()
	} else {
		err := fmt.Errorf("field Available is required%w", validation.ErrUserInput)
		return nil, err
	}
	arg.Tags = req.GetTags()

	result, err := s.querier.CreateBook(ctx, arg)
	if err != nil {
		s.logger.Error("CreateBook sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.CreateBookResponse{Book: toBook(result)}, nil
}

func (s *Service) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	bookID := req.GetBookId()

	err := s.querier.DeleteBook(ctx, bookID)
	if err != nil {
		s.logger.Error("DeleteBook sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.DeleteBookResponse{}, nil
}

func (s *Service) GetAuthor(ctx context.Context, req *pb.GetAuthorRequest) (*pb.GetAuthorResponse, error) {
	authorID := req.GetAuthorId()

	result, err := s.querier.GetAuthor(ctx, authorID)
	if err != nil {
		s.logger.Error("GetAuthor sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetAuthorResponse{Author: toAuthor(result)}, nil
}

func (s *Service) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	bookID := req.GetBookId()

	result, err := s.querier.GetBook(ctx, bookID)
	if err != nil {
		s.logger.Error("GetBook sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.GetBookResponse{Book: toBook(result)}, nil
}

func (s *Service) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	var arg UpdateBookParams
	arg.Title = req.GetTitle()
	arg.Tags = req.GetTags()
	arg.BookType = BookType(req.GetBookType())
	arg.BookID = req.GetBookId()

	err := s.querier.UpdateBook(ctx, arg)
	if err != nil {
		s.logger.Error("UpdateBook sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdateBookResponse{}, nil
}

func (s *Service) UpdateBookISBN(ctx context.Context, req *pb.UpdateBookISBNRequest) (*pb.UpdateBookISBNResponse, error) {
	var arg UpdateBookISBNParams
	arg.Title = req.GetTitle()
	arg.Tags = req.GetTags()
	arg.BookID = req.GetBookId()
	arg.Isbn = req.GetIsbn()

	err := s.querier.UpdateBookISBN(ctx, arg)
	if err != nil {
		s.logger.Error("UpdateBookISBN sql call failed", zap.Error(err))
		return nil, err
	}
	return &pb.UpdateBookISBNResponse{}, nil
}

func (s *Service) WithTx(tx *sql.Tx) *Service {
	return &Service{
		logger:  s.logger,
		querier: s.querier.WithTx(tx),
	}
}
