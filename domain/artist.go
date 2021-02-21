package domain

// Artist represents a artist user of the Projeto Rua app
type Artist struct {
	User
	PIX string `bson:"pix" json:"pix"`
}

// ArtistSaver is the interface needed to be implemented to save a regular user
type ArtistSaver interface {
	Save(Artist) error
}
