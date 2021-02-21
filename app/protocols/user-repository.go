package protocols

import "github.com/iamseki/rua-api/domain"

// SaveUserRepository is the interface implemented by database dependencie injected into DbSaveUser
type SaveUserRepository interface {
	Save(domain.User) error
}
