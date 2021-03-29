package event_test

import (
	"errors"
	"testing"
	"time"

	"github.com/iamseki/rua-api/app/event"
	"github.com/iamseki/rua-api/domain"
)

type fakeRepository struct {
	SaveFn func(m domain.EventModel) error
}

func (fr *fakeRepository) Save(m domain.EventModel) error {
	return fr.SaveFn(m)
}

func makeEventModel() domain.EventModel {
	return domain.EventModel{
		Artist: domain.Artist{
			User: domain.User{
				Name:     "any-name",
				LastName: "any-last-name",
				Email:    "any-email@mail.com",
				Password: "any-password",
			},
			PIX: "12345678",
		},
		Event: domain.Event{
			Title:           "any-event",
			Date:            time.Now(),
			Place:           "any-place",
			PlaceAnnotation: "near-any",
			Category:        "any",
			KeyWords:        []string{"any-1", "any-2"},
		},
	}
}

func TestSaveEventSucceed(t *testing.T) {
	e := makeEventModel()
	r := &fakeRepository{SaveFn: func(m domain.EventModel) error {
		return nil
	}}

	sut := event.NewDbSaveEvent(r)
	err := sut.Save(e)
	if err != nil {
		t.Errorf("Error on save event into repository: %v\n", err)
	}
}

func TestSaveEventError(t *testing.T) {
	e := makeEventModel()
	r := &fakeRepository{SaveFn: func(m domain.EventModel) error {
		return errors.New("fake error")
	}}

	sut := event.NewDbSaveEvent(r)
	err := sut.Save(e)
	if err == nil {
		t.Error("Expect to be an error with fake error message")
	}
}
