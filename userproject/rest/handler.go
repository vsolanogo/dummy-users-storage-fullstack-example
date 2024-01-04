package rest

import (
	"net/http"
	"userproject/models"
	"userproject/store"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userStore *store.UserStore
}

func NewHandler() (*Handler, error) {
	userStore := store.NewUserStore()

	handler := &Handler{
		userStore: userStore,
	}

	return handler, nil
}

func (h *Handler) GetUsers(c *gin.Context) {
	users := h.userStore.GetUserList()

	c.JSON(http.StatusOK, users)
}

func (h *Handler) AddUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = h.userStore.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) RemoveUser(c *gin.Context) {
	userID := c.Param("id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return
	}

	removed := h.userStore.RemoveUser(userID)
	if !removed {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User removed successfully"})
}

func (h *Handler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return
	}

	user, exists := h.userStore.GetUser(userID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) PatchUser(c *gin.Context) {
	userID := c.Param("id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return
	}

	var updatedUser models.User
	err := c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, exists := h.userStore.GetUser(userID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update only the non-empty fields
	if updatedUser.Name != "" {
		existingUser.Name = updatedUser.Name
	}

	if updatedUser.Email != "" {
		existingUser.Email = updatedUser.Email
	}

	// Update the user in the store
	h.userStore.UpdateUser(existingUser)

	c.JSON(http.StatusOK, existingUser)
}
