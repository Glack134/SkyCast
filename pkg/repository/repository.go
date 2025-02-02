package repository

type Skyfunction interface {
}

type Repository struct {
	Skyfunction
}

func NewRepository() *Repository {
	return &Repository{}
}
