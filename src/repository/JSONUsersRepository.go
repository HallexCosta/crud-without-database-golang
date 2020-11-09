package repository

import (
	"crud-without-database-golang/src/config"
	"crud-without-database-golang/src/entity"
	"crud-without-database-golang/src/helpers"
	"encoding/json"
)

// PersistenceFile ...
type PersistenceFile = helpers.PersistenceFile

// Persistence ...
type Persistence = helpers.Persistence

// UserRepositoryInterface ...
type UserRepositoryInterface interface {
	Save(user *entity.User)
	FindAll() []*entity.User
}

// UserRepository ...
type UserRepository struct {
}

// Save ...
func (userRepository *UserRepository) Save(user *entity.User) {
	var persist PersistenceFile = &Persistence{
		Name: config.ImportAppConfig().Persistence.Name,
	}

	var users []*entity.User = userRepository.FindAll()

	users = append(users, user)

	updatedUsers, _ := json.MarshalIndent(users, "", "\t")

	persist.WriteFile(updatedUsers)
}

// FindAll ...
func (userRepository *UserRepository) FindAll() []*entity.User {
	var users []*entity.User

	var persist PersistenceFile = &Persistence{
		Name: config.ImportAppConfig().Persistence.Name,
	}

	getUsers := persist.ReadFile()

	_ = json.Unmarshal(getUsers, &users)

	return users
}
