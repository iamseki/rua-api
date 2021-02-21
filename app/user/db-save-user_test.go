package userusecases

import (
	"errors"
	"testing"

	"github.com/iamseki/rua-api/domain"
)

type fakeRepository struct {
	SaveFn func(u domain.User) error
}

func (fr *fakeRepository) Save(u domain.User) error {
	return fr.SaveFn(u)
}

func makeUser() domain.User {
	return domain.User{
		Name:     "any-name",
		LastName: "any-last-name",
		Email:    "any-email@mail.com",
		Password: "any-password",
	}
}

func TestSaveUserSucceed(t *testing.T) {
	u := makeUser()
	r := &fakeRepository{SaveFn: func(u domain.User) error {
		return nil
	}}

	sut := NewDbSaveUser(r)
	err := sut.Save(u)
	if err != nil {
		t.Errorf("Error on save user into repository: %v\n", err)
	}
}

func TestSaveUserError(t *testing.T) {
	u := makeUser()
	r := &fakeRepository{SaveFn: func(u domain.User) error {
		return errors.New("fake error")
	}}

	sut := NewDbSaveUser(r)
	err := sut.Save(u)
	if err == nil {
		t.Error("Expect to be an error with fake error message")
	}
}
