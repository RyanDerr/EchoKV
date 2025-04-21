package service

import (
	"context"

	deletecmd "github.com/RyanDerr/EchoKV/pkg/delete"
	pb "github.com/RyanDerr/EchoKV/proto-public/go"
	validate "github.com/go-ozzo/ozzo-validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateDeleteRequest(req *pb.DeleteRequest) error {
	return validate.ValidateStruct(req, validate.Field(&req.Key, validate.Required))
}

func (s *Service) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := validateDeleteRequest(req)

	if err != nil {
		s.logger.Error("validation error", "error", err)
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %s", err.Error())
	}

	// Perform the delete operation on the cache
	err = deletecmd.Delete(s.store, req.GetKey())
	if err != nil {
		switch err {
		case deletecmd.ErrKeyNotFound:
			s.logger.Error("key not found", "key", req.GetKey())
			return nil, status.Errorf(codes.NotFound, "key not found: %s", req.GetKey())
		default:
			s.logger.Error("error deleting key", "error", err)
			return nil, status.Errorf(codes.Internal, "error deleting key: %s", err.Error())
		}
	}

	// Log the successful deletion
	s.logger.Info("key deleted", "key", req.GetKey())
	return &pb.DeleteResponse{}, nil
}
