package service

import (
	"context"

	setcmd "github.com/RyanDerr/GoKeyValueStore/pkg/set"
	pb "github.com/RyanDerr/GoKeyValueStore/proto-public/go"
	validation "github.com/go-ozzo/ozzo-validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateSetRequest(req *pb.SetRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Key, validation.Required),
		validation.Field(&req.Value, validation.Required),
	)
}

func (s *Service) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	err := validateSetRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %s", err.Error())
	}

	resp := setcmd.Set(s.cache, req.GetKey(), req.GetValue())

	return &pb.SetResponse{Key: req.GetKey(), Value: resp}, nil
}
