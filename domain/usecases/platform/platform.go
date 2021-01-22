package platform

// Usecase -
type Usecase struct {
	repository IRepository
}

// New -
func New(r IRepository) *Usecase {
	return &Usecase{
		r,
	}
}

// Drop -
func (u *Usecase) Drop() error {
	return u.repository.Drop()
}
