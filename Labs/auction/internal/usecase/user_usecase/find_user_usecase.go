package user_usecase

import (
	"context"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/user_entity"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
)

func NewUserUseCase(userRepository user_entity.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{userRepository}
}

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

type UserOutputDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserUseCaseInterface interface {
	FindUserById(ctx context.Context, userId string) (*UserOutputDto, *internal_error.InternalError)
}

func (u *UserUseCase) FindUserById(ctx context.Context,
	userId string) (*UserOutputDto, *internal_error.InternalError) {
	userEntity, err := u.UserRepository.FindUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &UserOutputDto{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}, err
}
