package supplier

import (


	"codemen.org/inventory/usermgm/storage"
)



type SupplierStore interface {
	CreateSupplier(storage.Supplier) (*storage.Supplier, error)
	ListOfSupplier(storage.SupplierFilter) ([]storage.Supplier, error)
	GetSupplierByID(int) (*storage.Supplier, error)
	UpdateSupplier(storage.Supplier) (*storage.Supplier, error)
	DeleteSupplierByID(int) error 
	GetSupplierByName(string) (*storage.Supplier, error)
	
}
type CoreSupplier struct { 
	store SupplierStore
}

func NewCoreSupplier(ss SupplierStore) *CoreSupplier {
	return &CoreSupplier{
		store: ss,
	}
}

func (cs CoreSupplier) CreateSupplier(s storage.Supplier) (*storage.Supplier, error){
	rs, err :=cs.store.CreateSupplier(s)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cs CoreSupplier) ListOfSupplier(sf storage.SupplierFilter) ([]storage.Supplier, error){
	supplierlist, err:= cs.store.ListOfSupplier(sf)
	if err != nil{
		return nil, err
	}
	return supplierlist, nil
}

func (cs CoreSupplier) GetSupplierByID(id int) (*storage.Supplier, error){

	rs, err := cs.store.GetSupplierByID(id)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cs CoreSupplier) UpdateSupplier(sp storage.Supplier) (*storage.Supplier, error){

	rs, err := cs.store.UpdateSupplier(sp)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cs CoreSupplier) DeleteSupplier(id int)  error{

	err := cs.store.DeleteSupplierByID(id)
	if err != nil{
		return err
	}
	
	return nil
}
func (cs CoreSupplier) GetSupplierByName(name string) (*storage.Supplier, error) {
     supplier, err:= cs.store.GetSupplierByName(name)
	 if err != nil{
		return nil, err
	}
	return supplier, nil
}
