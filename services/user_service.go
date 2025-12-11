package services

import (
	"errors"
	"fmt"

	"helloworld/models"
	"helloworld/repositories"
)

var (
	ErrInvalidEmail = errors.New("email inválido")
	ErrInvalidAge   = errors.New("la edad debe ser mayor a 0")
	ErrInvalidName  = errors.New("el nombre no puede estar vacío")
)

// UserService maneja la lógica de negocio relacionada con usuarios
type UserService interface {
	CreateUser(req models.CreateUserRequest) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(id string, req models.UpdateUserRequest) (*models.User, error)
	DeleteUser(id string) error
}

type userService struct {
	repo repositories.UserRepository
}

// NewUserService crea una nueva instancia del servicio de usuarios
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

// CreateUser crea un nuevo usuario con validación
func (s *userService) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	if err := s.validateCreateRequest(req); err != nil {
		return nil, err
	}

	user := models.NewUser(req)
	if err := s.repo.Create(user); err != nil {
		return nil, fmt.Errorf("error al crear usuario: %w", err)
	}

	return user, nil
}

// GetUserByID obtiene un usuario por su ID
func (s *userService) GetUserByID(id string) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuario: %w", err)
	}
	return user, nil
}

// GetAllUsers obtiene todos los usuarios
func (s *userService) GetAllUsers() ([]*models.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %w", err)
	}
	return users, nil
}

// UpdateUser actualiza un usuario existente
func (s *userService) UpdateUser(id string, req models.UpdateUserRequest) (*models.User, error) {
	existingUser, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuario: %w", err)
	}

	// Aplicar actualizaciones parciales
	if req.Name != nil {
		if *req.Name == "" {
			return nil, ErrInvalidName
		}
		existingUser.Name = *req.Name
	}
	if req.Email != nil {
		if !s.isValidEmail(*req.Email) {
			return nil, ErrInvalidEmail
		}
		existingUser.Email = *req.Email
	}
	if req.Age != nil {
		if *req.Age <= 0 {
			return nil, ErrInvalidAge
		}
		existingUser.Age = *req.Age
	}

	if err := s.repo.Update(id, existingUser); err != nil {
		return nil, fmt.Errorf("error al actualizar usuario: %w", err)
	}

	return existingUser, nil
}

// DeleteUser elimina un usuario
func (s *userService) DeleteUser(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("error al eliminar usuario: %w", err)
	}
	return nil
}

// validateCreateRequest valida los datos de creación de usuario
func (s *userService) validateCreateRequest(req models.CreateUserRequest) error {
	if req.Name == "" {
		return ErrInvalidName
	}
	if !s.isValidEmail(req.Email) {
		return ErrInvalidEmail
	}
	if req.Age <= 0 {
		return ErrInvalidAge
	}
	return nil
}

// isValidEmail realiza una validación básica de email
func (s *userService) isValidEmail(email string) bool {
	return email != "" && contains(email, "@")
}

// contains verifica si un string contiene un substring
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
