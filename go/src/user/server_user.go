package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "app/user/pb"
	app_service "app/user/app_service"
)

const (
	port = ":9001"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedUserServiceServer
}

// GET User
func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var id = in.Id
	var user, err = app_service.GetUser(id)
	return &pb.GetUserResponse{User: &user}, err
}

func (s *server) GetUserList(ctx context.Context, in *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	var userList, err = app_service.GetUserList()
	return &pb.GetUserListResponse{UserList: userList}, err
}

// Create User
func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var request = *in
	var id, err = app_service.CreateUser(request)
	if err == nil{
		return &pb.CreateUserResponse{Id: id}, nil
	}else{
		return &pb.CreateUserResponse{}, err
	}
}

// Update User
func (s *server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var request = *in
	var id, err = app_service.UpdateUser(request)
	if err == nil{
		return &pb.UpdateUserResponse{Id: id}, nil
	}else{
		return &pb.UpdateUserResponse{}, err
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}