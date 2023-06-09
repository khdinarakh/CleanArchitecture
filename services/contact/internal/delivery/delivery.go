package Delivery

import (
	"architecture_go/services/contact/internal/repository"
	"architecture_go/services/contact/internal/useCase"
	_ "architecture_go/services/contact/internal/useCase"
)

type Delivery struct {
	useCase    useCase.UseCase
	repository repository.Repository
}

func NewDelivery(useCase useCase.UseCase, repository repository.Repository) *Delivery {
	return &Delivery{
		useCase:    useCase,
		repository: repository,
	}
}
