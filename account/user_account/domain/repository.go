package domain

type UserAccountRepository interface {
	Insert(userAccount *RootEntity) error
	FindByID(id ID) (*RootEntity, error)
	Update(userAccount *RootEntity) error
	Delete(id ID) error
}
