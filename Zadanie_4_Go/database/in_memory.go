package database

import (
	"go_app/models"
	"sync"
)

// UserStore simulates a database using an in-memory map
type UserStore struct {
	mu    sync.Mutex
	users map[string]*models.User // userID to User mapping
}

// NewUserStore creates and returns a new UserStore
func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]*models.User),
	}
}

// AddUser adds a new user to the store
func (s *UserStore) AddUser(user *models.User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[user.ID] = user
}

// GetUser retrieves a user by ID
func (s *UserStore) GetUser(id string) (*models.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, exists := s.users[id]
	return user, exists
}

// GetUserByEmail retrieves a user by email
func (s *UserStore) GetUserByEmail(email string) (*models.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, user := range s.users {
		if user.Email == email {
			return user, true
		}
	}
	return nil, false
}
