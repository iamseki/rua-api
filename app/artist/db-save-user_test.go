package artistusecases

import (
	"errors"
	"testing"

	"github.com/iamseki/rua-api/domain"
)

type fakeRepository struct {
	SaveFn func(u domain.Artist) error
}

func (fr *fakeRepository) Save(u domain.Artist) error {
	return fr.SaveFn(u)
}

func makeUser() domain.Artist {
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
	u := makeUser()
	r := &fakeRepository{SaveFn: func(u domain.Artist) error {
		return nil
	}}

	sut := NewDbSaveArtist(r)
	err := sut.Save(u)
	if err != nil {
		t.Errorf("Error on save user into repository: %v\n", err)
	}
}

func TestSaveArtistError(t *testing.T) {
	u := makeUser()
	r := &fakeRepository{SaveFn: func(u domain.Artist) error {
		return errors.New("fake error")
	}}

	sut := NewDbSaveArtist(r)
	err := sut.Save(u)
	if err == nil {
		t.Error("Expect to be an error with fake error message")
	}
}
