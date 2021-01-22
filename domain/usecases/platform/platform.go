package platform

// Usecase - platform usecases
type Usecase struct {
	repository IRepository
}

// New - creates platform usecases instance
func New(r IRepository) *Usecase {
	return &Usecase{
		r,
	}
}

// Drop - drops database
func (u *Usecase) Drop() error {
	return u.repository.Drop()
}
