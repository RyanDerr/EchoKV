package service

import (
	"context"

	getcmd "github.com/RyanDerr/GoKeyValueStore/pkg/get"
	pb "github.com/RyanDerr/GoKeyValueStore/proto-public/go"
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
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %s", err.Error())
	}

	resp, err := getcmd.Get(s.cache, req.GetKey())
	if err != nil {
		switch err {
		case getcmd.ErrKeyNotFound:
			return nil, status.Errorf(codes.NotFound, "key not found: %s", req.GetKey())
		default:
			return nil, status.Errorf(codes.Internal, "error getting key: %s", err.Error())
		}
	}

	return &pb.GetResponse{Key: req.GetKey(), Value: resp}, nil
}
