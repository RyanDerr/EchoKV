package service

import (
	"context"

	getcmd "github.com/RyanDerr/EchoKV/pkg/get"
	pb "github.com/RyanDerr/EchoKV/proto-public/go"
	validation "github.com/go-ozzo/ozzo-validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateGetRequest(req *pb.GetRequest) error {
	return validation.ValidateStruct(req, validation.Field(&req.Key, validation.Required))
}

func (s *Service) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	err := validateGetRequest(req)
	if err != nil {
		s.logger.Error("validation error", "error", err)
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %s", err.Error())
	}

	resp, err := getcmd.Get(s.store, req.GetKey())
	if err != nil {
		switch err {
		case getcmd.ErrKeyNotFound:
			s.logger.Error("key not found", "key", req.GetKey())
			return nil, status.Errorf(codes.NotFound, "key not found: %s", req.GetKey())
		default:
			s.logger.Error("error getting key", "error", err)
			return nil, status.Errorf(codes.Internal, "error getting key: %s", err.Error())
		}
	}

	// Log the successful retrieval
	s.logger.Info("key retrieved", "key", req.GetKey(), "value", resp)
	return &pb.GetResponse{Key: req.GetKey(), Value: resp}, nil
}
