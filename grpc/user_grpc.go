package server

import (
	"context"
	"golang-crud/models"
	"golang-crud/repository"
	pb "path/to/generated/user"
)

type UserGRPCServer struct {
    pb.UnimplementedUserServiceServer
    repo *repository.UserRepository
}

func (s *UserGRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    user := models.Users{
        Name: req.Name,
        Email: req.Email,
        Password: req.Password,
        Address: req.Address,
    }
    id, err := s.repo.Insert(user)
    if err != nil {
        return nil, err
    }
    return &pb.CreateUserResponse{Id: id}, nil
}

func (s *UserGRPCServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
    user, err := s.repo.GetOneUser(int(req.Id))
    if err != nil {
        return nil, err
    }
    return &pb.GetUserResponse{
        User: &pb.User{
            Id: user.Id,
            Name: user.Name,
            Email: user.Email,
            Password: user.Password,
            Address: user.Address,
        },
    }, nil
}

func (s *UserGRPCServer) GetAllUsers(ctx context.Context, req *pb.Empty) (*pb.GetAllUsersResponse, error) {
    users, err := s.repo.GetAll()
    if err != nil {
        return nil, err
    }
    var respUsers []*pb.User
    for _, u := range users {
        respUsers = append(respUsers, &pb.User{
            Id: u.Id,
            Name: u.Name,
            Email: u.Email,
            Password: u.Password,
            Address: u.Address,
        })
    }
    return &pb.GetAllUsersResponse{Users: respUsers}, nil
}

// UpdateUser এবং DeleteUser একইভাবে implement করতে হবে
