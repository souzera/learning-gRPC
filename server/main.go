package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	//pb "github.com/souzera/learning-gRPC/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users  map[int32]*pb.User
	mu     sync.RWMutex
	nextID int32
}

func newServer() *server {
	return &server{
		users:  make(map[int32]*pb.User),
		nextID: 1,
	}
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[req.Id]
	if !exists {
		return nil, fmt.Errorf("user with ID %d not found", req.Id)
	}

	return user, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &pb.User{
		Id:    s.nextID,
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	s.users[s.nextID] = user
	s.nextID++

	return &pb.CreateUserResponse{
		User:    user,
		Message: "user created successfully",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listen.Close()

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, newServer())

	fmt.Println("Server is running on port :50051")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
