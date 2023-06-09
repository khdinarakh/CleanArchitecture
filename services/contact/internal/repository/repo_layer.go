package repository

import (
	"architecture_go/services/contact/internal/domain"
	"context"
)

type RepositoryLayerImplementation struct {
}

func NewRepository() Repository {
	return &RepositoryLayerImplementation{}
}

func (r *RepositoryLayerImplementation) CreateContact(ctx context.Context, contact domain.Contact) (domain.Contact, error) {
	return domain.Contact{}, nil
}

func (r *RepositoryLayerImplementation) ReadContact(ctx context.Context, contact domain.Contact) (domain.Contact, error) {
	return domain.Contact{}, nil
}
func (r *RepositoryLayerImplementation) UpdateContact(ctx context.Context, contact domain.Contact) (updated domain.Contact, err error) {
	return domain.Contact{}, nil
}

func (r *RepositoryLayerImplementation) DeleteContact(ctx context.Context, contact domain.Contact) (id int, err error) {
	return 0, nil
}
func (r *RepositoryLayerImplementation) CreateGroup(ctx context.Context, g domain.Group) (domain.Group, error) {
	return domain.Group{}, nil
}

func (r *RepositoryLayerImplementation) GetGroup(ctx context.Context, g domain.Group) (domain.Group, error) {
	return domain.Group{}, nil
}

func (r *RepositoryLayerImplementation) AddContactToGroup(ctx context.Context, c domain.Contact, g domain.Group) (domain.Contact, domain.Group, error) {
	return domain.Contact{}, domain.Group{}, nil
}
