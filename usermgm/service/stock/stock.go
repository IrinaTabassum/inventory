package stock

import (
	"context"

	stockpb "codemen.org/inventory/gunk/v1/stock"
	"codemen.org/inventory/usermgm/storage"
)


type CoreStock interface {
	ListOfStock(storage.StockFilter) ([]storage.Stock, error)
	GetProductAndStockByID(int) (*storage.ProductStock, error)
}

type StockSvc struct {
	stockpb.UnimplementedStockServiceServer
	core   CoreStock
}

func NewStockSvc(cs CoreStock) *StockSvc {
	return &StockSvc{
		core:   cs,
	}
}
func (ss StockSvc) ListStock(ctx context.Context, r *stockpb.ListStockRequest) (*stockpb.ListStockResponse, error){
	Stockfilts := storage.StockFilter{
		SearchTerm: r.GetSearchTerm(),
		Offset:     int(r.GetOffset()),
		Limit:      int(r.GetLimit()),
	}

	sl, err := ss.core.ListOfStock(Stockfilts)
	if err != nil {
		return nil, err
	}
	list := make([]*stockpb.Stock, len(sl))
	for i, st := range sl {
		list[i] = &stockpb.Stock{
			ID:          int32(st.ID),
			ProductId:   int32(st.ProductId),
			ProductNmae: st.ProductNane,
			Quantity:    int32(st.Quantity),
		}
	}
	return &stockpb.ListStockResponse{
		Stocks: list,
	}, nil
}

func (ss StockSvc) GetStock(ctx context.Context, r *stockpb.GetStockRequest) (*stockpb.GetStockResponse, error){
	pID := int(r.GetID())

	stock, err := ss.core.GetProductAndStockByID(pID)
	if err != nil {
		return nil, err
	}
	
	return &stockpb.GetStockResponse{
		Stock: &stockpb.Stock{
			ID:          int32(stock.ID),
			ProductId:   int32(stock.ProductId),
			ProductNmae: stock.ProductName,
			Quantity:    int32(stock.StockQuantity),
		},
	}, nil
}



