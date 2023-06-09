package useCase

import (
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/repository"
	"context"
)

type LayerImplementation struct {
	repo repository.Repository
}

func NewRepository(repository repository.Repository) UseCase {
	return &LayerImplementation{repo: repository}
}

func (r *LayerImplementation) CreateContact(ctx context.Context, contact domain.Contact) (domain.Contact, error) {
	return domain.Contact{}, nil
}

func (r *LayerImplementation) ReadContact(ctx context.Context, contact domain.Contact) (domain.Contact, error) {
	return domain.Contact{}, nil
}
func (r *LayerImplementation) UpdateContact(ctx context.Context, contact domain.Contact) (updated domain.Contact, err error) {
	return domain.Contact{}, nil
}

func (r *LayerImplementation) DeleteContact(ctx context.Context, contact domain.Contact) (id int, err error) {
	return 0, nil
}
func (r *LayerImplementation) CreateGroup(ctx context.Context, g domain.Group) (domain.Group, error) {
	return domain.Group{}, nil
}

func (r *LayerImplementation) GetGroup(ctx context.Context, g domain.Group) (domain.Group, error) {
	return domain.Group{}, nil
}

func (r *LayerImplementation) AddContactToGroup(ctx context.Context, c domain.Contact, g domain.Group) (domain.Contact, domain.Group, error) {
	return domain.Contact{}, domain.Group{}, nil
}
