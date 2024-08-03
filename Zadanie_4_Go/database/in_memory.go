package database

import (
	"go_app/models"
	"sync"
)

type UserStore struct {
    mu    sync.RWMutex
    Users map[string]*models.User
}

func NewUserStore() *UserStore {
    return &UserStore{
        Users: make(map[string]*models.User),
    }
}

func (s *UserStore) AddUser(user *models.User) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.Users[user.ID] = user
}

func (s *UserStore) GetUser(id string) (*models.User, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    user, exists := s.Users[id]
    return user, exists
}