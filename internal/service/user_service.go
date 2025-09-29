package service

import (
	"context"
	"errors"
	"go-api/internal/config"
	"go-api/internal/model"
	"go-api/internal/repository"
	"go-api/pkg/md5"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (*model.User, error)
	Register(ctx context.Context, username string, password string) (*model.User, error)
	GetByID(ctx context.Context, id uint) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) Login(ctx context.Context, username string, password string) (*model.User, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("用户或密码错误")
	}

	if user.Password != md5.MD5(password + config.Get().App.Salt) {
		return nil, errors.New("用户或密码错误")
	}

	return user, nil
}

func (s *userService) Register(ctx context.Context, username string, password string) (*model.User, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, errors.New("获取用户失败")
	}

	if user.ID > 0 {
		return nil, errors.New("用户已存在")
	}

	user.Username = username
	user.Password = password
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, errors.New("创建用户失败")
	}
	return user, nil
}

func (s *userService) GetByID(ctx context.Context, id uint) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("获取用户失败")
	}

	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}

	user.Password = ""

	return user, nil
}


func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}