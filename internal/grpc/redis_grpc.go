package grpc

import (
	"context"
	"io"

	"github.com/PyMarcus/gRPC-redis/internal/repository"
	pb "github.com/PyMarcus/gRPC-redis/proto"
)

type GRPCServer struct {
	pb.UnimplementedKVStoreServer
	redis *repository.RedisRepository
}

func NewGRPCServer(redis *repository.RedisRepository) *GRPCServer{
	return &GRPCServer{redis: redis}
}

func (s *GRPCServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error){
	val, err := s.redis.Get(ctx, req.Key)

	if err != nil{
		return &pb.GetResponse{Value: ""}, err 
	}

	return &pb.GetResponse{Value: val}, nil
}

func (s *GRPCServer) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	err := s.redis.Set(ctx, req.Key, req.Value)
	return &pb.SetResponse{Success: err == nil}, err
}

func (s *GRPCServer) Del(ctx context.Context, req *pb.DelRequest) (*pb.DelResponse, error) {
	err := s.redis.Del(ctx, req.Key)
	if err != nil {
		return &pb.DelResponse{Value: ""}, err
	}
	return &pb.DelResponse{Value: req.Key}, nil
}

// StreamSet: permite múltiplas gravações 
func (s *GRPCServer) StreamSet(stream pb.KVStore_StreamSetServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil 
		}
		if err != nil {
			return err
		}

		err = s.redis.Set(stream.Context(), req.Key, req.Value)

		res := &pb.SetResponse{Success: err == nil}
		if sendErr := stream.Send(res); sendErr != nil {
			return sendErr
		}
	}
}

// StreamGet: permite múltiplas leituras
func (s *GRPCServer) StreamGet(stream pb.KVStore_StreamGetServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		val, err := s.redis.Get(stream.Context(), req.Key)
		res := &pb.GetResponse{Key: req.Key}
		if err != nil {
			res.Error = err.Error()
		} else {
			res.Value = val
		}

		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
