package protocols

import "github.com/iamseki/rua-api/domain"

// SaveEventRepository is the interface implemented by database dependencie injected into DbSaveEvent
type SaveEventRepository interface {
	Save(domain.EventModel) error
}

// FindEventRepository is the interface implemented by database dependencie injected into DbFindEvent
type FindEventRepository interface {
	Find(domain.EventFinderOptions) (domain.Event, error)
}
