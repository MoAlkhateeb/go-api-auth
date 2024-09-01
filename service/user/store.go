package user

import (
	"github.com/MoAlkhateeb/go-api-auth/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	err := s.db.Where("email = ?", email).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	var user types.User
	err := s.db.First(&user, id)
	if err != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (s *Store) CreateUser(user types.User) error {
	result := s.db.Create(&user)
	return result.Error
}
