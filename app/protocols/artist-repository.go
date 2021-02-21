package protocols

import "github.com/iamseki/rua-api/domain"

// SaveArtistRepository is the interface implemented by database dependencie injected into DbSaveArtist
type SaveArtistRepository interface {
	Save(domain.Artist) error
}
