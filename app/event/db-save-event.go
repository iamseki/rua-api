package event

import (
	"github.com/iamseki/rua-api/app/protocols"
	"github.com/iamseki/rua-api/domain"
)

// DbSaveEvent implements domain.EventSaver interface
type DbSaveEvent struct {
	repository protocols.SaveEventRepository
}

// NewDbSaveEvent returns an instace of DbSaveEvent
func NewDbSaveEvent(repository protocols.SaveEventRepository) *DbSaveEvent {
	return &DbSaveEvent{repository}
}

// Save saves an event into event repository
func (db *DbSaveEvent) Save(m domain.EventModel) error {
	err := db.repository.Save(m)
	if err != nil {
		return err
	}
	return nil
}
