package domain

// User represents a regular user of the Projeto Rua app
type User struct {
	Name     string `bson:"name" json:"name"`
	LastName string `bson:"lastName" json:"lastName"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

// UserSaver is the interface needed to be implemented to save a regular user
type UserSaver interface {
	Save(User) error
}
