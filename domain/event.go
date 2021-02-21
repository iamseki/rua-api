package domain

import "time"

// Event represents a event happening on the "Rua"
type Event struct {
	Title           string    `bson:"title" json:"title"`
	Date            time.Time `bson:"date" json:"date"`
	Place           string    `bson:"place" json:"place"`
	PlaceAnnotation string    `bson:"placeAnnotation" json:"placeAnnotation"`
	Category        string    `bson:"category" json:"category"`
	KeyWords        []string  `bson:"keyWords" json:"keyWords"`
}

// EventModel represents the input to pass to the EventSaver interface to save a event properly
type EventModel struct {
	Artist
	Event
}

// EventSaver is the interface needed to be implemented to save an event
type EventSaver interface {
	Save(EventModel) error
}
