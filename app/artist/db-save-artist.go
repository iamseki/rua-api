package artistusecases

import (
	"github.com/iamseki/rua-api/app/protocols"
	"github.com/iamseki/rua-api/domain"
)

// DbSaveArtist implements domain.ArtistSaver interface
type DbSaveArtist struct {
	repository protocols.SaveArtistRepository
}

// NewDbSaveArtist returns an instance of DbSaveArtist object
func NewDbSaveArtist(repository protocols.SaveArtistRepository) *DbSaveArtist {
	return &DbSaveArtist{repository}
}

// Save saves a artist user into user repository
func (db *DbSaveArtist) Save(u domain.Artist) error {
	err := db.repository.Save(u)
	if err != nil {
		return err
	}
	return nil
}
