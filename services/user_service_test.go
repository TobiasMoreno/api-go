package services

import (
	"testing"

	"helloworld/models"
	"helloworld/repositories"
)

// mockRepository es un mock del repositorio para testing
type mockRepository struct {
	users map[string]*models.User
}

func newMockRepository() *mockRepository {
	return &mockRepository{
		users: make(map[string]*models.User),
	}
}

func (m *mockRepository) Create(user *models.User) error {
	m.users[user.ID] = user
	return nil
}

func (m *mockRepository) GetByID(id string) (*models.User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, repositories.ErrUserNotFound
	}
	return user, nil
}

func (m *mockRepository) GetAll() ([]*models.User, error) {
	users := make([]*models.User, 0, len(m.users))
	for _, user := range m.users {
		users = append(users, user)
	}
	return users, nil
}

func (m *mockRepository) Update(id string, user *models.User) error {
	if _, exists := m.users[id]; !exists {
		return repositories.ErrUserNotFound
	}
	m.users[id] = user
	return nil
}

func (m *mockRepository) Delete(id string) error {
	if _, exists := m.users[id]; !exists {
		return repositories.ErrUserNotFound
	}
	delete(m.users, id)
	return nil
}

func TestUserService_CreateUser(t *testing.T) {
	repo := newMockRepository()
	service := NewUserService(repo)

	tests := []struct {
		name    string
		req     models.CreateUserRequest
		wantErr bool
		errType error
	}{
		{
			name: "crear usuario válido",
			req: models.CreateUserRequest{
				Name:  "Juan Pérez",
				Email: "juan@example.com",
				Age:   30,
			},
			wantErr: false,
		},
		{
			name: "crear usuario sin nombre",
			req: models.CreateUserRequest{
				Name:  "",
				Email: "juan@example.com",
				Age:   30,
			},
			wantErr: true,
			errType: ErrInvalidName,
		},
		{
			name: "crear usuario con email inválido",
			req: models.CreateUserRequest{
				Name:  "Juan Pérez",
				Email: "email-invalido",
				Age:   30,
			},
			wantErr: true,
			errType: ErrInvalidEmail,
		},
		{
			name: "crear usuario con edad inválida",
			req: models.CreateUserRequest{
				Name:  "Juan Pérez",
				Email: "juan@example.com",
				Age:   -1,
			},
			wantErr: true,
			errType: ErrInvalidAge,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := service.CreateUser(tt.req)
			if tt.wantErr {
				if err == nil {
					t.Errorf("CreateUser() esperaba error, pero no obtuvo ninguno")
					return
				}
				if err != tt.errType {
					t.Errorf("CreateUser() error = %v, esperaba %v", err, tt.errType)
				}
			} else {
				if err != nil {
					t.Errorf("CreateUser() error = %v, no esperaba error", err)
					return
				}
				if user == nil {
					t.Errorf("CreateUser() retornó usuario nil")
					return
				}
				if user.Name != tt.req.Name {
					t.Errorf("CreateUser() nombre = %v, esperaba %v", user.Name, tt.req.Name)
				}
			}
		})
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	repo := newMockRepository()
	service := NewUserService(repo)

	// Crear un usuario de prueba
	req := models.CreateUserRequest{
		Name:  "Test User",
		Email: "test@example.com",
		Age:   25,
	}
	createdUser, err := service.CreateUser(req)
	if err != nil {
		t.Fatalf("Error al crear usuario de prueba: %v", err)
	}

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "obtener usuario existente",
			id:      createdUser.ID,
			wantErr: false,
		},
		{
			name:    "obtener usuario inexistente",
			id:      "id-inexistente",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := service.GetUserByID(tt.id)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetUserByID() no retornó error como se esperaba")
				}
			} else {
				if err != nil {
					t.Errorf("GetUserByID() error = %v, no esperaba error", err)
					return
				}
				if user == nil {
					t.Errorf("GetUserByID() retornó usuario nil")
				}
			}
		})
	}
}

func TestUserService_GetAllUsers(t *testing.T) {
	repo := newMockRepository()
	service := NewUserService(repo)

	// Crear algunos usuarios
	users := []models.CreateUserRequest{
		{Name: "User 1", Email: "user1@example.com", Age: 25},
		{Name: "User 2", Email: "user2@example.com", Age: 30},
	}

	for _, req := range users {
		_, err := service.CreateUser(req)
		if err != nil {
			t.Fatalf("Error al crear usuario: %v", err)
		}
	}

	allUsers, err := service.GetAllUsers()
	if err != nil {
		t.Errorf("GetAllUsers() error = %v, no esperaba error", err)
		return
	}

	if len(allUsers) != len(users) {
		t.Errorf("GetAllUsers() retornó %d usuarios, esperaba %d", len(allUsers), len(users))
	}
}

