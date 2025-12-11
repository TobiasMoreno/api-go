package repositories

import (
	"errors"

	"helloworld/models"
)

var (
	ErrUserNotFound = errors.New("usuario no encontrado")
)

// UserRepository define la interfaz para el almacenamiento y recuperación de usuarios
type UserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(id string, user *models.User) error
	Delete(id string) error
}

