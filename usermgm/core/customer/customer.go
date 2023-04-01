package customer

import (


	"codemen.org/inventory/usermgm/storage"
)



type CustomerStore interface {
	CreateCustomer(storage.Customer) (*storage.Customer, error)
	ListOfCustomer(storage.CustomerFilter) ([]storage.Customer, error)
	GetCustomerByID(int) (*storage.Customer, error)
	UpdateCustomer(storage.Customer) (*storage.Customer, error)
	DeleteCustomerByID(int) error 
	GetCustomerByName(string) (*storage.Customer, error)
	
}
type CoreCustomer struct { 
	store CustomerStore
}

func NewCoreCustomer(ss CustomerStore) *CoreCustomer {
	return &CoreCustomer{
		store: ss,
	}
}

func (cs CoreCustomer) CreateCustomer(s storage.Customer) (*storage.Customer, error){
	rs, err :=cs.store.CreateCustomer(s)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cs CoreCustomer) ListOfCustomer(sf storage.CustomerFilter) ([]storage.Customer, error){
	customerlist, err:= cs.store.ListOfCustomer(sf)
	if err != nil{
		return nil, err
	}
	return customerlist, nil
}

func (cs CoreCustomer) GetCustomerByID(id int) (*storage.Customer, error){

	rs, err := cs.store.GetCustomerByID(id)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cs CoreCustomer) UpdateCustomer(sp storage.Customer) (*storage.Customer, error){

	rs, err := cs.store.UpdateCustomer(sp)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cs CoreCustomer) DeleteCustomer(id int)  error{

	err := cs.store.DeleteCustomerByID(id)
	if err != nil{
		return err
	}
	
	return nil
}
func (cs CoreCustomer) GetCustomerByNmae(name string) (*storage.Customer, error){

	customer, err := cs.store.GetCustomerByName(name)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
