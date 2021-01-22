package platform

import "github.com/smart-evolution/shapi/data/persistence"

// Repository -
type Repository struct {
	Persistence persistence.IPersistance
}

// New -
func New(p persistence.IPersistance) *Repository {
	return &Repository{
		p,
	}
}

// Drop -
func (r *Repository) Drop() error {
	return r.Persistence.DropDatabase()
}
