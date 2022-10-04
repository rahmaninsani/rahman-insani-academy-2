package usecase

import (
	"go-echo-clean-arch/entity"
	"go-echo-clean-arch/repository"
)

type IUserUseCase interface {
	CreateUser(entity.UserRequest) (entity.UserResponse, error)
	GetAllUser() ([]entity.UserResponse, error)
	GetUserById(int) (entity.UserResponse, error)
	UpdateUser(entity.UserRequest, int) (entity.UserResponse, error)
	DeleteUser(int) error
}

type UserUseCase struct {
	userRepository repository.IUserRepository
}

func NewUserUseCase(userRepository repository.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepository}
}

func (useCase UserUseCase) CreateUser(user entity.UserRequest) (entity.UserResponse, error) {
	u := entity.User{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	createdUser, err := useCase.userRepository.Store(u)

	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
		Phone: createdUser.Phone,
	}

	return userRes, nil
}

func (useCase UserUseCase) GetAllUser() ([]entity.UserResponse, error) {
	users, err := useCase.userRepository.FindAll()

	if err != nil {
		return []entity.UserResponse{}, err
	}

	var userRes []entity.UserResponse

	for _, user := range users {
		userRes = append(userRes, entity.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		})
	}

	return userRes, nil
}

func (useCase UserUseCase) GetUserById(id int) (entity.UserResponse, error) {
	user, err := useCase.userRepository.FindById(id)

	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	return userRes, nil
}

func (useCase UserUseCase) UpdateUser(userRequest entity.UserRequest, id int) (entity.UserResponse, error) {
	user, err := useCase.userRepository.FindById(id)

	if err != nil {
		return entity.UserResponse{}, err
	}

	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.Phone = userRequest.Phone

	updatedUser, err := useCase.userRepository.Update(user)
	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		ID:    updatedUser.ID,
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
		Phone: updatedUser.Phone,
	}

	return userRes, nil
}

func (useCase UserUseCase) DeleteUser(id int) error {
	err := useCase.userRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
