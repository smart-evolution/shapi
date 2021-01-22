package platform

import "github.com/smart-evolution/shapi/data/persistence"

// Repository - platform repository
type Repository struct {
	Persistence persistence.IPersistance
}

// New - creates new platform repository
func New(p persistence.IPersistance) *Repository {
	return &Repository{
		p,
	}
}

// Drop - drops whole database
func (r *Repository) Drop() error {
	return r.Persistence.DropDatabase()
}
