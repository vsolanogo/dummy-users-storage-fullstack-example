package store

import (
	"sync"
	"userproject/models"

	"github.com/google/uuid"
)

type UserStore struct {
	mu    sync.RWMutex
	users map[string]models.User
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]models.User),
	}
}

func (us *UserStore) AddUser(user models.User) models.User {
	us.mu.Lock()
	defer us.mu.Unlock()

	// Generate a UUID for the user
	user.ID = uuid.New().String()

	// Store the user in the map with the UUID as the key
	us.users[user.ID] = user

	return user
}

func (us *UserStore) GetUser(userID string) (models.User, bool) {
	us.mu.RLock()
	defer us.mu.RUnlock()

	user, exists := us.users[userID]
	return user, exists
}

func (us *UserStore) GetUserList() []models.User {
	us.mu.RLock()
	defer us.mu.RUnlock()

	// Create a slice to store the list of users
	userList := make([]models.User, 0, len(us.users))

	// Iterate over the map and append users to the slice
	for _, user := range us.users {
		userList = append(userList, user)
	}

	return userList
}

func (us *UserStore) RemoveUser(userID string) bool {
	us.mu.Lock()
	defer us.mu.Unlock()

	_, exists := us.users[userID]
	if !exists {
		return false // User not found
	}

	delete(us.users, userID)
	return true
}

func (us *UserStore) UpdateUser(updatedUser models.User) bool {
	us.mu.Lock()
	defer us.mu.Unlock()

	userToUpdate, exists := us.users[updatedUser.ID]
	if !exists {
		return false // User not found
	}

	if updatedUser.Name != "" {
		userToUpdate.Name = updatedUser.Name
	}

	if updatedUser.Email != "" {
		userToUpdate.Email = updatedUser.Email
	}

	us.users[updatedUser.ID] = userToUpdate

	return true
}
