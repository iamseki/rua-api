package eventusecases

import (
	"github.com/iamseki/rua-api/app/protocols"
	"github.com/iamseki/rua-api/domain"
)

// DbFindEvent implements the EventFinder business rule
type DbFindEvent struct {
	repository protocols.FindEventRepository
}

// NewDbFindEvent returns an instance of DbFindEvent
func NewDbFindEvent(repository protocols.FindEventRepository) *DbFindEvent {
	return &DbFindEvent{repository}
}

// Find method search a event with options criteria
func (db *DbFindEvent) Find(options domain.EventFinderOptions) (domain.Event, error) {
	event, err := db.repository.Find(options)
	if err != nil {
		return domain.Event{}, err
	}
	return event, nil
}
