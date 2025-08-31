package service

import (
	"context"

	setcmd "github.com/RyanDerr/EchoKV/pkg/set"
	pb "github.com/RyanDerr/EchoKV/proto-public/go"
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
		s.logger.Error("validation error", "error", err)
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %s", err.Error())
	}

	resp := setcmd.Set(s.store, req.GetKey(), req.GetValue())
	s.logger.Info("key set", "key", req.GetKey(), "value", resp)
	return &pb.SetResponse{Key: req.GetKey(), Value: resp}, nil
}
