package service

import (
	"context"

	deletecmd "github.com/RyanDerr/GoKeyValueStore/pkg/delete"
	pb "github.com/RyanDerr/GoKeyValueStore/proto-public/go"
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
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %s", err.Error())
	}

	// Perform the delete operation on the cache
	err = deletecmd.Delete(s.cache, req.GetKey())
	if err != nil {
		switch err {
		case deletecmd.ErrKeyNotFound:
			return nil, status.Errorf(codes.NotFound, "key not found: %s", req.GetKey())
		default:
			return nil, status.Errorf(codes.Internal, "error deleting key: %s", err.Error())
		}
	}

	return &pb.DeleteResponse{}, nil
}
