package app

import (
	"context"
	"fmt"

	pb "github.com/sahilrana7582/orderX/pkg/generated/user"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	fmt.Printf("Register user: Name=%v, Email=%v, Password=%v\n", req.Name, req.Email, req.Password)

	return &pb.RegisterUserResponse{
		Id:      "mock-user-id",
		Message: "User registered successfully",
	}, nil
}

func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	fmt.Printf("Login user: Email=%v, Password=%v\n", req.Email, req.Password)

	return &pb.LoginUserResponse{
		Token: "mock-jwt-token",
		User: &pb.User{
			Id:     "mock-user-id",
			Name:   "Mock User",
			Email:  req.Email,
			Role:   pb.UserRole_USER_ROLE_USER,
			Status: pb.UserStatus_USER_STATUS_ACTIVE,
		},
	}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	fmt.Printf("Get user by ID: %v\n", req.Id)

	return &pb.GetUserByIDResponse{
		User: &pb.User{
			Id:     req.Id,
			Name:   "Mock User",
			Email:  "mock@example.com",
			Role:   pb.UserRole_USER_ROLE_USER,
			Status: pb.UserStatus_USER_STATUS_ACTIVE,
		},
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	fmt.Printf("Update user: ID=%v, Name=%v, Email=%v, Role=%v, Status=%v\n",
		req.Id, req.Name, req.Email, req.Role, req.Status)

	return &pb.UpdateUserResponse{
		User: &pb.User{
			Id:     req.Id,
			Name:   req.Name,
			Email:  req.Email,
			Role:   req.Role,
			Status: req.Status,
		},
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	fmt.Printf("Delete user: ID=%v\n", req.Id)

	return &pb.DeleteUserResponse{
		Message: fmt.Sprintf("User with ID %s deleted successfully", req.Id),
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	fmt.Printf("List users: Page=%v, Limit=%v\n", req.Page, req.Limit)

	users := []*pb.User{
		{
			Id:     "user1",
			Name:   "User One",
			Email:  "user1@example.com",
			Role:   pb.UserRole_USER_ROLE_USER,
			Status: pb.UserStatus_USER_STATUS_ACTIVE,
		},
		{
			Id:     "user2",
			Name:   "User Two",
			Email:  "user2@example.com",
			Role:   pb.UserRole_USER_ROLE_ADMIN,
			Status: pb.UserStatus_USER_STATUS_ACTIVE,
		},
	}

	return &pb.ListUsersResponse{
		Users: users,
		Total: int32(len(users)),
	}, nil
}
