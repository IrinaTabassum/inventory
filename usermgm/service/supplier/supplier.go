package supplier

import (
	"context"
	"fmt"

	supplierpb "codemen.org/inventory/gunk/v1/supplier"
	"codemen.org/inventory/usermgm/storage"
)

type CoreSupplier interface {
	CreateSupplier(storage.Supplier) (*storage.Supplier, error)
	ListOfSupplier(storage.SupplierFilter) ([]storage.Supplier, error)
	GetSupplierByID(int) (*storage.Supplier, error)
	UpdateSupplier(storage.Supplier) (*storage.Supplier, error)
	DeleteSupplier(int) error
	GetSupplierByName(string) (*storage.Supplier, error)
	
}

type SupplierSvc struct {
	supplierpb.UnimplementedSupplierServiceServer
	core CoreSupplier
}

func NewSupplierSvc(cs CoreSupplier) *SupplierSvc {
	return &SupplierSvc{
		core: cs,
	}
}

func (ss SupplierSvc) CreateSupplier(ctx context.Context, r *supplierpb.CreateSupplierRequest) (*supplierpb.CreateSupplierResponse, error) {
	supplier := storage.Supplier{
		Name: r.GetName(),
	}
	sup,_ := ss.core.GetSupplierByName(supplier.Name)
	if sup != nil {
		return nil, fmt.Errorf("supplier name must be unique")
	}

	if err := supplier.Validate(); err != nil {
		fmt.Println(err)
		return nil, err 
	}

	s, err := ss.core.CreateSupplier(supplier)
	if err != nil {
		return nil, err
	}

	return &supplierpb.CreateSupplierResponse{
		Supplier: &supplierpb.Supplier{
			ID:   int32(s.ID),
			Name: s.Name,
		},
	}, nil
}
func (ss SupplierSvc) ListSupplier(ctx context.Context, r *supplierpb.ListSupplierRequest) (*supplierpb.ListSupplierResponse, error){
	supplierfilts := storage.SupplierFilter{
		SearchTerm: r.GetSearchTerm(),
		Offset:     int(r.GetOffset()),
		Limit:      int(r.GetLimit()),
	}
	sl, err := ss.core.ListOfSupplier(supplierfilts)
	if err != nil {
		return nil, err
	}
	
	list := make([]*supplierpb.Supplier, len(sl))
	for i, sup := range sl {
		list[i] = &supplierpb.Supplier{
			ID:   int32(sup.ID),
			Name: sup.Name,
		}
	}
	return &supplierpb.ListSupplierResponse{
		Suppliers: list,
	}, nil
}

func (ss SupplierSvc) GetSupplier(ctx context.Context, r *supplierpb.GetSupplierRequest) (*supplierpb.GetSupplierResponse, error){
	sId := int(r.GetID())
	

	s, err := ss.core.GetSupplierByID(sId)
	if err != nil {
		return nil, err
	}

	return &supplierpb.GetSupplierResponse{
		Supplier: &supplierpb.Supplier{
			ID:   int32(s.ID),
			Name: s.Name,
		},
	},nil
}

func (ss SupplierSvc) UpdateSupplie(ctx context.Context, r *supplierpb.UpdateSupplierRequest) (*supplierpb.UpdateSupplierResponse, error){
	supplier := storage.Supplier{
		ID:        int(r.GetSupplier().ID),
		Name:      r.GetSupplier().Name,
	}
	s, err := ss.core.UpdateSupplier(supplier)
	if err != nil {
		return nil, err
	}
	return &supplierpb.UpdateSupplierResponse{
		Supplier: &supplierpb.Supplier{
			ID:   int32(s.ID),
			Name: s.Name,
		},
	},nil

}

func (ss SupplierSvc) DeleteSupplier(ctx context.Context, r *supplierpb.DeleteSupplierRequest) (*supplierpb.DeleteSupplierResponse, error){
	sId := int(r.GetID())

	err := ss.core.DeleteSupplier(sId)
	if err != nil {
		return nil,err
	}

	return &supplierpb.DeleteSupplierResponse{},nil
}