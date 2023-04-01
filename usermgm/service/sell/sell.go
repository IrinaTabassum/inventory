package sell

import (
	"context"
	"fmt"

	sellpb "codemen.org/inventory/gunk/v1/sell"
	"codemen.org/inventory/usermgm/storage"
)

type CoreSell interface {
	CreateSell(storage.Sell) (storage.Sell, error)
	ListOfSell(storage.SellFilter) ([]storage.Sell, error)
	GetSellByID(int) (*storage.Sell, error)
	UpdateSell(storage.Sell) (*storage.Sell, error)
}

type CoreStock interface {
	CreateOrUpdateStock(storage.Stock) (*storage.Stock, error)
	GetProductAndStockByID(int) (*storage.ProductStock, error)
}

type CoreSellProduct interface {
	CreateSellProduct(storage.SellProductmap) ([]storage.SoldProduct, error)
	GetSellProductByID(int) ([]storage.SellProduct, error)
}

type SellSvc struct {
	sellpb.UnimplementedSellServiceServer
	core   CoreSell
	coreC  CoreStock
	coreSP CoreSellProduct
}

func NewSellSvc(cs CoreSell, cst CoreStock, csp CoreSellProduct) *SellSvc {
	return &SellSvc{
		core:   cs,
		coreC:  cst,
		coreSP: csp,
	}
}

func (ps SellSvc) CreateSell(ctx context.Context, r *sellpb.CreateSellRequest) (*sellpb.CreateSellResponse, error) {

	sell := storage.Sell{
		CustomerId: int(r.GetCustomerId()),
		SelPro:     r.GetProductQuantity(),
	}
	if err := sell.Validate(); err != nil {
		fmt.Println(err)
		return nil, err 
	}
	for k, v := range sell.SelPro {
		s, err := ps.coreC.GetProductAndStockByID(int(k))
		if err != nil {
			fmt.Println(err)
		}
		if s.StockQuantity < int(v) {
			fmt.Printf("%v is not avilabe in %v quantity\n", s.ProductName, v)
			return nil, fmt.Errorf("quantity out of stock")
		}
	}

	p, err := ps.core.CreateSell(sell)
	if err != nil {
		return nil, err
	}

	sellPro := storage.SellProductmap{
		SellId: p.ID,
		SelPro: r.GetProductQuantity(),
	}

	sp, err := ps.coreSP.CreateSellProduct(sellPro)
	if err != nil {
		return nil, err
	}

	var (
		Sold       []*sellpb.SoldProduct
		totalPrice float32
	)
	for _, v := range sp {
		s, _ := ps.coreC.GetProductAndStockByID(v.ProductId)
		stProduct := storage.Stock{
			ProductId: v.ProductId,
			Quantity:  s.StockQuantity - v.Quantity,
		}
		_, err := ps.coreC.CreateOrUpdateStock(stProduct)
		if err != nil {
			fmt.Println(err)
		}

		totalPrice += v.TotalUnitPrice
		sold := &sellpb.SoldProduct{
			ProductId:      int32(v.ProductId),
			ProductName:    v.ProductName,
			UnitPrice:      v.UnitPrice,
			Quantity:       int32(v.Quantity),
			TotalUnitPrice: v.TotalUnitPrice,
		}
		Sold = append(Sold, sold)

	}
	p.TotalPrice = totalPrice
	newsp, err := ps.core.UpdateSell(p)
	if err != nil {
		return nil, err
	}
	
	return &sellpb.CreateSellResponse{
		Sell: &sellpb.Sell{
			ID:           int32(p.ID),
			CustomerId:   int32(p.CustomerId),
			SoldProducts: Sold,
			TotalPrice:   newsp.TotalPrice,
		},
	}, nil
}

func (ps SellSvc) ListSell(ctx context.Context, r *sellpb.ListSellRequest) (*sellpb.ListSellResponse, error){
	Sellfilts := storage.SellFilter{
		SearchTerm: r.GetSearchTerm(),
		Offset:     int(r.GetOffset()),
		Limit:      int(r.GetLimit()),
	}
	sl, err := ps.core.ListOfSell(Sellfilts)
	if err != nil {
		return nil, err
	}

	list := make([]*sellpb.Sell, len(sl))
	for i, sellPro := range sl {
		sp, err:= ps.coreSP.GetSellProductByID(sellPro.ID)
		if err != nil {
			return nil, err
		}
		var Sold []*sellpb.SoldProduct
		for _, v := range  sp{
			sold := &sellpb.SoldProduct{
				ProductId:      int32(v.ProductId),
				ProductName:    v.ProductName,
				UnitPrice:      v.UnitPrice,
				Quantity:       int32(v.Quantity),
				TotalUnitPrice: v.TotalUnitPrice,
			}
			Sold = append(Sold, sold)
		}
		list[i] = &sellpb.Sell{
			ID:           int32(sellPro.ID),
			CustomerId:   int32(sellPro.CustomerId),
			CustomerName: sellPro.CustomerName,
			SoldProducts: Sold,
			TotalPrice:   sellPro.TotalPrice,
		}
	}
	return &sellpb.ListSellResponse{
		Sells: list,
	}, nil
}

func (ps SellSvc) GetSell(ctx context.Context, r *sellpb.GetSellRequest) (*sellpb.GetSellResponse, error){
	sId := int(r.GetID())
	p, err := ps.core.GetSellByID(sId)
	if err != nil {
		return nil, err
	}

    sp, err:= ps.coreSP.GetSellProductByID(p.ID)
	if err != nil {
		return nil, err
	}

	var Sold []*sellpb.SoldProduct
	for _, v := range  sp{
		sold := &sellpb.SoldProduct{
			ProductId:      int32(v.ProductId),
			ProductName:    v.ProductName,
			UnitPrice:      v.UnitPrice,
			Quantity:       int32(v.Quantity),
			TotalUnitPrice: v.TotalUnitPrice,
		}
		Sold = append(Sold, sold)
	}
	return &sellpb.GetSellResponse{
		Sell: &sellpb.Sell{
			ID:           int32(p.ID),
			CustomerId:   int32(p.CustomerId),
			CustomerName: p.CustomerName,
			SoldProducts: Sold,
			TotalPrice:   p.TotalPrice,
		},
	},nil
}


