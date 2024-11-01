package repository

import (
	"clean/domain"

	"gorm.io/gorm"
)

// logika database
type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}
func (r *UserRepositoryImpl) GetByUsername(email string) (*domain.User, error) {
	var user domain.User
	err := r.DB.Where("email = ?", email).First(&user).Error // pastikan "email" sesuai kolom database
	return &user, err
}
