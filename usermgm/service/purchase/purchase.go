package purchase

import (
	"context"

	purchasepb "codemen.org/inventory/gunk/v1/purchase"
	"codemen.org/inventory/usermgm/storage"
)

type CorePurchase interface {
	CreatePurchase(storage.Purchase) (*storage.Purchase, error)
	ListOfPurchase(storage.PurchaseFilter) ([]storage.Purchase, error)
	GetPurchaseByID(int) (*storage.Purchase, error)
	UpdatePurchase(sp storage.Purchase) (*storage.Purchase, error)
	
}

type PurchaseSvc struct {
	purchasepb.UnimplementedPurchaseServiceServer
	core CorePurchase
}

func NewPurchaseSvc(cp CorePurchase) *PurchaseSvc {
	return &PurchaseSvc{
		core: cp,
	}
}

func (ps PurchaseSvc) CreatePurchase(ctx context.Context, r *purchasepb.CreatePurchaseRequest) (*purchasepb.CreatePurchaseResponse, error) {
	purchase := storage.Purchase{
		SupplierId: int(r.GetSupplierId()),
		ProductId:  int(r.GetProductId()),
		Quantity:   int(r.GetQuantity()),
	}
	if err := purchase.Validate(); err != nil {
		return nil, err 
	}

	p, err := ps.core.CreatePurchase(purchase)
	if err != nil {
		return nil, err
	}

	return &purchasepb.CreatePurchaseResponse{
		Purchase: &purchasepb.Purchase{
			ID:         int32(p.ID),
			SupplierId: int32(p.SupplierId),
			ProductId:  int32(p.ProductId),
			Quantity:   int32(p.Quantity),
			UnitPrice:  p.UnitPrice,
			TotalPrice: p.TotalPrice,
		},
	}, nil
}
func (ps PurchaseSvc) ListPurchase(ctx context.Context, r *purchasepb.ListPurchaseRequest) (*purchasepb.ListPurchaseResponse, error){
	Purchasefilts := storage.PurchaseFilter{
		SearchTerm: r.GetSearchTerm(),
		Offset:     int(r.GetOffset()),
		Limit:      int(r.GetLimit()),
	}
	sl, err := ps.core.ListOfPurchase(Purchasefilts)
	if err != nil {
		return nil, err
	}
	
	list := make([]*purchasepb.Purchase, len(sl))
	for i, pur := range sl {
		list[i] = &purchasepb.Purchase{
			ID:         int32(pur.ID),
			SupplierId: int32(pur.SupplierId),
			ProductId:  int32(pur.ProductId),
			Quantity:   int32(pur.Quantity),
			UnitPrice:  pur.UnitPrice,
			TotalPrice: pur.TotalPrice,
		}
	}
	return &purchasepb.ListPurchaseResponse{
		Purchases: list,
	}, nil
}

func (ps PurchaseSvc) GetPurchase(ctx context.Context, r *purchasepb.GetPurchaseRequest) (*purchasepb.GetPurchaseResponse, error){
	sId := int(r.GetID())
	
	p, err := ps.core.GetPurchaseByID(sId)
	if err != nil {
		return nil, err
	}

	return &purchasepb.GetPurchaseResponse{
		Purchase: &purchasepb.Purchase{
			ID:         int32(p.ID),
			SupplierId: int32(p.SupplierId),
			ProductId:  int32(p.ProductId),
			Quantity:   int32(p.Quantity),
			UnitPrice:  p.UnitPrice,
			TotalPrice: p.TotalPrice,
		},
	},nil
}

func (ps PurchaseSvc) UpdatePurchase(ctx context.Context, r *purchasepb.UpdatePurchaseRequest) (*purchasepb.UpdatePurchaseResponse, error){
	purchase := storage.Purchase{
		ID:         int(r.GetID()),
		SupplierId: int(r.GetSupplierId()),
		ProductId:  int(r.GetProductId()),
		Quantity:   int(r.GetQuantity()),
		
	}
	p, err := ps.core.UpdatePurchase(purchase)
	if err != nil {
		return nil, err
	}
	return &purchasepb.UpdatePurchaseResponse{
		Purchase: &purchasepb.Purchase{
			ID:         int32(p.ID),
			SupplierId: int32(p.SupplierId),
			ProductId:  int32(p.ProductId),
			Quantity:   int32(p.Quantity),
			UnitPrice:  p.UnitPrice,
			TotalPrice: p.TotalPrice,
		},
	},nil

}
