package infrastructure

import (
	"database/sql"
	"emopathy-api/account/user_account/domain"
)

type userAccountRepository struct {
	db *sql.DB
}

func NewUserAccountRepository(db *sql.DB) domain.UserAccountRepository {
	return &userAccountRepository{db: db}
}

func (r *userAccountRepository) Insert(userAccount *domain.RootEntity) error {
	// TODO: implement
	return nil
}

func (r *userAccountRepository) FindByID(id domain.ID) (*domain.RootEntity, error) {
	// TODO: implement
	return nil, nil
}

func (r *userAccountRepository) Update(userAccount *domain.RootEntity) error {
	// TODO: implement
	return nil
}

func (r *userAccountRepository) Delete(id domain.ID) error {
	// TODO: implement
	return nil
}
