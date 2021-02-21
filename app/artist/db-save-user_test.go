package artistusecases

import (
	"errors"
	"testing"

	"github.com/iamseki/rua-api/domain"
)

type fakeRepository struct {
	SaveFn func(a domain.Artist) error
}

func (fr *fakeRepository) Save(a domain.Artist) error {
	return fr.SaveFn(a)
}

func makeArtist() domain.Artist {
	return domain.Artist{
		User: domain.User{
			Name:     "any-name",
			LastName: "any-last-name",
			Email:    "any-email@mail.com",
			Password: "any-password",
		},
		PIX: "12345678",
	}
}

func TestSaveArtistSucceed(t *testing.T) {
	a := makeArtist()
	r := &fakeRepository{SaveFn: func(a domain.Artist) error {
		return nil
	}}

	sut := NewDbSaveArtist(r)
	err := sut.Save(a)
	if err != nil {
		t.Errorf("Error on save artist into repository: %v\n", err)
	}
}

func TestSaveArtistError(t *testing.T) {
	a := makeArtist()
	r := &fakeRepository{SaveFn: func(a domain.Artist) error {
		return errors.New("fake error")
	}}

	sut := NewDbSaveArtist(r)
	err := sut.Save(a)
	if err == nil {
		t.Error("Expect to be an error with fake error message")
	}
}
