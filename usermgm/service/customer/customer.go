package customer

import (
	"context"
	"fmt"

	customerpb "codemen.org/inventory/gunk/v1/customer"
	"codemen.org/inventory/usermgm/storage"
)

type CoreCustomer interface {
	CreateCustomer(storage.Customer) (*storage.Customer, error)
	ListOfCustomer(storage.CustomerFilter) ([]storage.Customer, error)
	GetCustomerByID(int) (*storage.Customer, error)
	UpdateCustomer(sp storage.Customer) (*storage.Customer, error)
	DeleteCustomer(id int) error
	GetCustomerByNmae(string) (*storage.Customer, error)
	
}

type CustomerSvc struct {
	customerpb.UnimplementedCustomerServiceServer
	core CoreCustomer
}

func NewCustomerSvc(cc CoreCustomer) *CustomerSvc {
	return &CustomerSvc{
		core: cc,
	}
}

func (cs CustomerSvc) CreateCustomer(ctx context.Context, r *customerpb.CreateCustomerRequest) (*customerpb.CreateCustomerResponse, error) {
	customer := storage.Customer{
		Name: r.GetName(),
	}
	cus,_ := cs.core.GetCustomerByNmae(customer.Name)
	if cus != nil{
		return nil, fmt.Errorf("customer name must be unique")
	}
	if err := customer.Validate(); err != nil {
		fmt.Println(err)
		return nil, err 
	}

	s, err := cs.core.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}

	return &customerpb.CreateCustomerResponse{
		Customer: &customerpb.Customer{
			ID:   int32(s.ID),
			Name: s.Name,
		},
	}, nil
}
func (cs CustomerSvc) ListCustomer(ctx context.Context, r *customerpb.ListCustomerRequest) (*customerpb.ListCustomerResponse, error){
	Customerfilts := storage.CustomerFilter{
		SearchTerm: r.GetSearchTerm(),
		Offset:     int(r.GetOffset()),
		Limit:      int(r.GetLimit()),
	}
	sl, err := cs.core.ListOfCustomer(Customerfilts)
	if err != nil {
		return nil, err
	}
	
	list := make([]*customerpb.Customer, len(sl))
	for i, sup := range sl {
		list[i] = &customerpb.Customer{
			ID:   int32(sup.ID),
			Name: sup.Name,
		}
	}
	return &customerpb.ListCustomerResponse{
		Customers: list,
	}, nil
}

func (cs CustomerSvc) GetCustomer(ctx context.Context, r *customerpb.GetCustomerRequest) (*customerpb.GetCustomerResponse, error){
	sId := int(r.GetID())
	

	s, err := cs.core.GetCustomerByID(sId)
	if err != nil {
		return nil, err
	}

	return &customerpb.GetCustomerResponse{
		Customer: &customerpb.Customer{
			ID:   int32(s.ID),
			Name: s.Name,
		},
	},nil
}

func (cs CustomerSvc) UpdateCustomer(ctx context.Context, r *customerpb.UpdateCustomerRequest) (*customerpb.UpdateCustomerResponse, error){
	customer := storage.Customer{
		ID:        int(r.GetCustomer().ID),
		Name:      r.GetCustomer().Name,
	}
	s, err := cs.core.UpdateCustomer(customer)
	if err != nil {
		return nil, err
	}
	return &customerpb.UpdateCustomerResponse{
		Customer: &customerpb.Customer{
			ID:   int32(s.ID),
			Name: s.Name,
		},
	},nil

}

func (cs CustomerSvc) DeleteCustomer(ctx context.Context, r *customerpb.DeleteCustomerRequest) (*customerpb.DeleteCustomerResponse, error){
	sId := int(r.GetID())

	err := cs.core.DeleteCustomer(sId)
	if err != nil {
		return nil,err
	}

	return &customerpb.DeleteCustomerResponse{},nil
}