package service

import "github.com/Glack134/SkyCast/pkg/repository"

type Skyfunction interface {
}

type Service struct {
	Skyfunction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
