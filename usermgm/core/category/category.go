package category

import (
	"fmt"

	"codemen.org/inventory/usermgm/storage"
)



type CategoryStore interface {
	CreateCategory(storage.Category) (*storage.Category, error)
	ListOfCategory(storage.CategoryFilter) ([]storage.Category, error)
	GetCategoryByID(int) (*storage.Category, error)
	UpdateCategory(storage.Category) (*storage.Category, error)
	DeleteCategoryByID(int) error 
	GetCategoryByName(string) (*storage.Category, error)
}
type CoreCategory struct { 
	store CategoryStore
}

func NewCoreCategory(ss CategoryStore) *CoreCategory {
	return &CoreCategory{
		store: ss,
	}
}

func (cs CoreCategory) CreateCategory(s storage.Category) (*storage.Category, error){
	rs, err :=cs.store.CreateCategory(s)
	if err != nil {
		return nil, err
	}
	if rs == nil {
		return nil, fmt.Errorf("unable to create category")
	}

	return rs, nil
}

func (cs CoreCategory) ListOfCategory(sf storage.CategoryFilter) ([]storage.Category, error){
	Categorylist, err:= cs.store.ListOfCategory(sf)
	if err != nil{
		return nil, err
	}
	return Categorylist, nil
}

func (cs CoreCategory) GetCategoryByID(id int) (*storage.Category, error){

	rs, err := cs.store.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cs CoreCategory) UpdateCategory(sp storage.Category) (*storage.Category, error){

	rs, err := cs.store.UpdateCategory(sp)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cs CoreCategory) DeleteCategory(id int)  error{

	err := cs.store.DeleteCategoryByID(id)
	if err != nil{
		return err
	}
	
	return nil
}

func (cs CoreCategory) GetCategoryByName(name string) (*storage.Category, error){

	category, err := cs.store.GetCategoryByName(name)
	if err != nil {
		return nil, err
	}

	return category, nil
}