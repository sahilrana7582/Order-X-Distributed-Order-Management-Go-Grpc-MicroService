package app

import (
	"context"
	"fmt"

	pb "github.com/sahilrana7582/orderX/pkg/generated/user"
	"github.com/sahilrana7582/user-service/internal/domain"
	"github.com/sahilrana7582/user-service/internal/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	var user domain.User

	user = domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     "USER",
		Status:   "ACTIVE",
	}

	err := s.userRepo.Create(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to register user: %v", err)
	}

	return &pb.RegisterUserResponse{
		Id:      user.ID,
		Message: fmt.Sprintf("User %s registered successfully", user.Name),
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

	user, err := s.userRepo.GetByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}

	return &pb.GetUserByIDResponse{
		User: &pb.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      pb.UserRole_USER_ROLE_USER,
			Status:    pb.UserStatus_USER_STATUS_ACTIVE,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {

	user, err := s.userRepo.Update(ctx, req.Id, &domain.UpdateUserRequest{
		Name:   req.Name,
		Email:  req.Email,
		Role:   req.Role.String(),
		Status: req.Status.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	return &pb.UpdateUserResponse{
		User: &pb.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      pb.UserRole_USER_ROLE_USER,
			Status:    pb.UserStatus_USER_STATUS_ACTIVE,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {

	err := s.userRepo.Delete(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %v", err)
	}

	return &pb.DeleteUserResponse{
		Message: fmt.Sprintf("User with ID %s deleted successfully", req.Id),
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	fmt.Printf("List users: Page=%v, Limit=%v\n", req.Page, req.Limit)

	users, err := s.userRepo.List(ctx, int(req.Limit), int(req.Page)*int(req.Limit))
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %v", err)
	}
	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      pb.UserRole_USER_ROLE_USER,
			Status:    pb.UserStatus_USER_STATUS_ACTIVE,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		})
	}

	return &pb.ListUsersResponse{
		Users: pbUsers,
		Total: int32(len(users)),
	}, nil
}
