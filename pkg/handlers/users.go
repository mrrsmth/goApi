package handlers

import (
	"go_api_template/pkg/services"
	"go_api_template/pkg/utils"
	"net/http"
	"strconv"
)

// UserHandler contains the business logic for users.
type UserHandler struct {
	usersService services.IUserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(usersService services.IUserService) *UserHandler {
	return &UserHandler{usersService: usersService}
}

// GetUserByID returns a user by their ID.
func (h *UserHandler) GetUserByID(logger utils.ILogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		idInt, err := strconv.Atoi(id)

		if err != nil {
			logger.Error("Invalid ID")
			utils.RespondError(w, http.StatusBadRequest, "Invalid ID")
			return
		}

		user, err := h.usersService.GetUserByID(idInt)

		if err != nil {
			logger.Errorf("Something went wrong getting user by ID: %v", err)
			utils.RespondError(w, http.StatusInternalServerError, "Error getting user")
			return
		}

		utils.RespondJSON(w, http.StatusOK, user)
	}
}
