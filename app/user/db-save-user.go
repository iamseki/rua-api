package userusecases

import (
	"github.com/iamseki/rua-api/app/protocols"
	"github.com/iamseki/rua-api/domain"
)

// DbSaveUser implements domain.UserSaver interface
type DbSaveUser struct {
	repository protocols.SaveUserRepository
}

// NewDbSaveUser returns an instance of DbSaveUser object
func NewDbSaveUser(repository protocols.SaveUserRepository) *DbSaveUser {
	return &DbSaveUser{repository}
}

// Save saves a regular user into user repository
func (db *DbSaveUser) Save(u domain.User) error {
	err := db.repository.Save(u)
	if err != nil {
		return err
	}
	return nil
}
