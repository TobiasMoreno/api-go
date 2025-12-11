package repositories

import (
	"database/sql"
	"fmt"

	"helloworld/models"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLUserRepository implementa UserRepository usando MySQL
type MySQLUserRepository struct {
	db *sql.DB
}

// NewMySQLUserRepository crea una nueva instancia del repositorio MySQL
func NewMySQLUserRepository(dsn string) (*MySQLUserRepository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al abrir conexión a MySQL: %w", err)
	}

	// Verificar la conexión
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al conectar con MySQL: %w", err)
	}

	// Crear la tabla si no existe
	if err := createTableIfNotExists(db); err != nil {
		return nil, fmt.Errorf("error al crear tabla: %w", err)
	}

	return &MySQLUserRepository{
		db: db,
	}, nil
}

// createTableIfNotExists crea la tabla de usuarios si no existe
func createTableIfNotExists(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(36) PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		age INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		INDEX idx_email (email)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
	`

	_, err := db.Exec(query)
	return err
}

// Close cierra la conexión a la base de datos
func (r *MySQLUserRepository) Close() error {
	return r.db.Close()
}

// Create guarda un nuevo usuario
func (r *MySQLUserRepository) Create(user *models.User) error {
	query := "INSERT INTO users (id, name, email, age) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Age)
	if err != nil {
		return fmt.Errorf("error al crear usuario: %w", err)
	}
	return nil
}

// GetByID obtiene un usuario por su ID
func (r *MySQLUserRepository) GetByID(id string) (*models.User, error) {
	query := "SELECT id, name, email, age FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("error al obtener usuario: %w", err)
	}

	return &user, nil
}

// GetAll obtiene todos los usuarios
func (r *MySQLUserRepository) GetAll() ([]*models.User, error) {
	query := "SELECT id, name, email, age FROM users ORDER BY created_at DESC"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %w", err)
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age); err != nil {
			return nil, fmt.Errorf("error al escanear usuario: %w", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar usuarios: %w", err)
	}

	return users, nil
}

// Update actualiza un usuario existente
func (r *MySQLUserRepository) Update(id string, user *models.User) error {
	query := "UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?"
	result, err := r.db.Exec(query, user.Name, user.Email, user.Age, id)
	if err != nil {
		return fmt.Errorf("error al actualizar usuario: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}

// Delete elimina un usuario
func (r *MySQLUserRepository) Delete(id string) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
