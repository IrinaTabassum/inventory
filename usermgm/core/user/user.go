package user

import (
	"fmt"

	"codemen.org/inventory/usermgm/storage"
	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	Register(storage.User) (*storage.User, error)
	GetUserByUsername(string) (*storage.User, error)
	ListUser(storage.UserFilter) ([]storage.User, error)
	UpdateUser(storage.User) (*storage.User, error)
	DeleteUserByID(int) (storage.User, error)
	GetUserByID(int) (*storage.User, error)
	GetUserByEmail(string) (*storage.User, error)
}

type CoreUser struct {
	store UserStore
}

func NewCoreUser(us UserStore) *CoreUser {
	return &CoreUser{
		store: us,
	}
}

func (cu CoreUser) Register(u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u.Password = string(hashPass)
	ru, err := cu.store.Register(u)
	if err != nil {
		return nil, err
	}

	if ru == nil {
		return nil, fmt.Errorf("unable to register")
	}

	return ru, nil
}

func (cu CoreUser) Login(l storage.Login) (*storage.User, error) {
	u, err := cu.store.GetUserByUsername(l.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password)); err != nil {
		return nil, err
	}

	return u, nil
}

func (cu CoreUser) ListUser(uf storage.UserFilter) ([]storage.User, error) {
	userList, err := cu.store.ListUser(uf)
	if err != nil {
		return nil, err
	}
	return userList, nil
}
func (cu CoreUser) UpdateUser(u storage.User) (*storage.User, error) {
	user, err := cu.store.UpdateUser(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (cu CoreUser) DeleteUserByID(id int) (storage.User, error) {
	user, err := cu.store.DeleteUserByID(id)
	if err != nil {
		return storage.User{}, err
	}
	return user, nil
}

func (cu CoreUser) GetUserByID(id int) (*storage.User, error) {
	user, err := cu.store.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (cu CoreUser) GetUserByUsername(un string) (*storage.User, error){
	u, err := cu.store.GetUserByUsername(un)
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (cu CoreUser) GetUserByEmail(email string) (*storage.User, error){
	u, err := cu.store.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return u, nil
}
