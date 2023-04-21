package application

import (
	"cleancodego/internal/domain/repository"
	"cleancodego/internal/domain/services"
	"fmt"
)

type userService struct {

	// This user service will have dependecy of our repository
	UserRepo repository.UserRepository
}

func NewUserService(UserRepo repository.UserRepository) services.UserService {
	return &userService{
		UserRepo,
	}
}

func (u *userService) SaveUser() string {
	fmt.Println(u.UserRepo.Save()) // it will invoke the repository method as well
	return "from user service save user called"
}
