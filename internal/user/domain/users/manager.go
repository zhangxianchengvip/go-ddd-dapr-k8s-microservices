package users

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserManager struct {
	repository *gorm.DB
}

func NewManager(repository *gorm.DB) (*UserManager, error) {

	if repository == nil {
		return nil, ErrRepositoryNil
	}

	return &UserManager{
		repository: repository,
	}, nil
}

func (m *UserManager) Create(loginname, password string) (*User, error) {

	var (
		user User
	)

	if m.repository.Where("loginname = ?", loginname).First(&user).Error == nil && user.ID != uuid.Nil {
		return nil, ErrUserAlreadyExists
	}

	newUser, err := NewUser(uuid.New(), loginname, password, Address{})

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
