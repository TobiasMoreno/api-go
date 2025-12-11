package handlers

import (
	"encoding/json"
	"net/http"

	"helloworld/models"
	"helloworld/services"

	"github.com/gorilla/mux"
)

// UserHandler maneja las peticiones HTTP relacionadas con usuarios
type UserHandler struct {
	service services.UserService
}

// NewUserHandler crea una nueva instancia del handler de usuarios
func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// CreateUser maneja la creación de un nuevo usuario
// @Summary      Crear un nuevo usuario
// @Description  Crea un nuevo usuario en el sistema
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        user  body      models.CreateUserRequest  true  "Datos del usuario"
// @Success      201   {object}  models.User
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "error al decodificar el cuerpo de la petición")
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == services.ErrInvalidEmail || err == services.ErrInvalidAge || err == services.ErrInvalidName {
			statusCode = http.StatusBadRequest
		}
		respondWithError(w, statusCode, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

// GetUser maneja la obtención de un usuario por ID
// @Summary      Obtener un usuario por ID
// @Description  Obtiene la información de un usuario específico
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID del usuario"
// @Success      200  {object}  models.User
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /users/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.service.GetUserByID(id)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "error al obtener usuario: usuario no encontrado" {
			statusCode = http.StatusNotFound
		}
		respondWithError(w, statusCode, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

// GetAllUsers maneja la obtención de todos los usuarios
// @Summary      Obtener todos los usuarios
// @Description  Obtiene una lista de todos los usuarios registrados
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.User
// @Failure      500  {object}  map[string]string
// @Router       /users [get]
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

// UpdateUser maneja la actualización de un usuario
// @Summary      Actualizar un usuario
// @Description  Actualiza la información de un usuario existente (actualización parcial)
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        id    path      string                  true  "ID del usuario"
// @Param        user  body      models.UpdateUserRequest  true  "Datos a actualizar"
// @Success      200   {object}  models.User
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users/{id} [put]
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "error al decodificar el cuerpo de la petición")
		return
	}

	user, err := h.service.UpdateUser(id, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "error al obtener usuario: usuario no encontrado" {
			statusCode = http.StatusNotFound
		} else if err == services.ErrInvalidEmail || err == services.ErrInvalidAge || err == services.ErrInvalidName {
			statusCode = http.StatusBadRequest
		}
		respondWithError(w, statusCode, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

// DeleteUser maneja la eliminación de un usuario
// @Summary      Eliminar un usuario
// @Description  Elimina un usuario del sistema
// @Tags         usuarios
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID del usuario"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /users/{id} [delete]
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.service.DeleteUser(id); err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "error al eliminar usuario: usuario no encontrado" {
			statusCode = http.StatusNotFound
		}
		respondWithError(w, statusCode, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "usuario eliminado correctamente"})
}

// respondWithJSON envía una respuesta JSON
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

// respondWithError envía una respuesta de error JSON
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, map[string]string{"error": message})
}
