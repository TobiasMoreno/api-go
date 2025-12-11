package repositories

import (
	"errors"
	"sync"

	"helloworld/models"
)

var (
	ErrUserNotFound = errors.New("usuario no encontrado")
)

// UserRepository maneja el almacenamiento y recuperación de usuarios
type UserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(id string, user *models.User) error
	Delete(id string) error
}

// InMemoryUserRepository implementa UserRepository usando memoria
type InMemoryUserRepository struct {
	users map[string]*models.User
	mu    sync.RWMutex
}

// NewInMemoryUserRepository crea una nueva instancia del repositorio en memoria
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*models.User),
	}
}

// Create guarda un nuevo usuario
func (r *InMemoryUserRepository) Create(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}

// GetByID obtiene un usuario por su ID
func (r *InMemoryUserRepository) GetByID(id string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, exists := r.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}

// GetAll obtiene todos los usuarios
func (r *InMemoryUserRepository) GetAll() ([]*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	users := make([]*models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

// Update actualiza un usuario existente
func (r *InMemoryUserRepository) Update(id string, user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.users[id]; !exists {
		return ErrUserNotFound
	}
	r.users[id] = user
	return nil
}

// Delete elimina un usuario
func (r *InMemoryUserRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.users[id]; !exists {
		return ErrUserNotFound
	}
	delete(r.users, id)
	return nil
}

