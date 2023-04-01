package product

import (


	"codemen.org/inventory/usermgm/storage"
)



type ProductStore interface {
	CreateProduct(storage.Product) (*storage.Product, error)
	ListOfProduct(storage.ProductFilter) ([]storage.Product, error)
	GetProductByID(int) (*storage.Product, error)
	UpdateProduct(storage.Product) (*storage.Product, error)
	DeleteProductByID(int) error 
	
}
type CoreProduct struct { 
	store ProductStore
}

func NewCoreProduct(ps ProductStore) *CoreProduct {
	return &CoreProduct{
		store: ps,
	}
}

func (ps CoreProduct) CreateProduct(p storage.Product) (*storage.Product, error){
	product, err :=ps.store.CreateProduct(p)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps CoreProduct) ListOfProduct(pf storage.ProductFilter) ([]storage.Product, error){
	productlist, err:= ps.store.ListOfProduct(pf)
	if err != nil{
		return nil, err
	}
	return productlist, nil
}

func (ps CoreProduct) GetProductByID(id int) (*storage.Product, error){
	product, err := ps.store.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps CoreProduct) UpdateProduct(sp storage.Product) (*storage.Product, error){
	editProduct, err := ps.store.UpdateProduct(sp)
	if err != nil {
		return nil, err
	}

	return editProduct, nil
}

func (ps CoreProduct) DeleteProduct(id int)  error{
	err := ps.store.DeleteProductByID(id)
	if err != nil{
		return err
	}
	
	return nil
}
