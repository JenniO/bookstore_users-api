package services

import (
	"github.com/JenniO/bookstore_users-api/domain/users"
	"github.com/JenniO/bookstore_users-api/utils/crypto_utils"
	"github.com/JenniO/bookstore_users-api/utils/date_utils"
	"github.com/JenniO/bookstore_users-api/utils/errors"
)

type (
	usersService          struct{}
	usersServiceInterface interface {
		GetUser(userId int64) (*users.User, *errors.RestErr)
		CreateUser(user users.User) (*users.User, *errors.RestErr)
		UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestErr)
		DeleteUser(userId int64) *errors.RestErr
		SearchUsers(status string) (users.Users, *errors.RestErr)
		LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr)
	}
)

var (
	UsersService usersServiceInterface = &usersService{}
)

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}

	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}

	return user.Delete()
}

func (s *usersService) SearchUsers(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
