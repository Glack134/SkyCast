package service

type Skyfunction interface {
	Search()
}

type Service struct {
	Skyfunction
}
