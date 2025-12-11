package models

import "github.com/google/uuid"

// User representa un usuario en el sistema
// @Description Usuario del sistema
type User struct {
	ID    string `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"` // ID único del usuario
	Name  string `json:"name" example:"Juan Pérez"`                          // Nombre del usuario
	Email string `json:"email" example:"juan@example.com"`                  // Email del usuario
	Age   int    `json:"age" example:"30"`                                   // Edad del usuario
}

// CreateUserRequest representa la solicitud para crear un usuario
// @Description Datos requeridos para crear un nuevo usuario
type CreateUserRequest struct {
	Name  string `json:"name" example:"Juan Pérez" binding:"required"`       // Nombre del usuario
	Email string `json:"email" example:"juan@example.com" binding:"required"` // Email del usuario
	Age   int    `json:"age" example:"30" binding:"required"`                 // Edad del usuario
}

// UpdateUserRequest representa la solicitud para actualizar un usuario
// @Description Datos opcionales para actualizar un usuario (actualización parcial)
type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty" example:"Juan Pérez"`       // Nuevo nombre (opcional)
	Email *string `json:"email,omitempty" example:"juan@example.com"` // Nuevo email (opcional)
	Age   *int    `json:"age,omitempty" example:"31"`                // Nueva edad (opcional)
}

// NewUser crea una nueva instancia de User con un ID generado
func NewUser(req CreateUserRequest) *User {
	return &User{
		ID:    uuid.New().String(),
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}
}

