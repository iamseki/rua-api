package event_test

import (
	"errors"
	"testing"

	"github.com/iamseki/rua-api/app/event"
	"github.com/iamseki/rua-api/domain"
)

type fakeFindRepository struct {
	FindFn func(o domain.EventFinderOptions) (domain.Event, error)
}

func (fr *fakeFindRepository) Find(o domain.EventFinderOptions) (domain.Event, error) {
	return fr.FindFn(o)
}

func makeEventFinderOptions() domain.EventFinderOptions {
	return domain.EventFinderOptions{
		Title: "any-event",
	}
}

func TestFindEventSucceed(t *testing.T) {
	r := &fakeFindRepository{FindFn: func(o domain.EventFinderOptions) (domain.Event, error) {
		return domain.Event{}, nil
	}}
	sut := event.NewDbFindEvent(r)
	_, err := sut.Find(makeEventFinderOptions())
	if err != nil {
		t.Fatalf("Error on find event : %v\n", err)
	}
}

func TestFindEventError(t *testing.T) {
	r := &fakeFindRepository{FindFn: func(o domain.EventFinderOptions) (domain.Event, error) {
		return domain.Event{}, errors.New("fake error")
	}}
	sut := event.NewDbFindEvent(r)
	_, err := sut.Find(makeEventFinderOptions())
	if err == nil {
		t.Fatal("Expected a fake error")
	}
}
