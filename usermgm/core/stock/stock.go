package stock


import (


	"codemen.org/inventory/usermgm/storage"
)



type StockStore interface {
	CreateOrUpdateStock(storage.Stock) (*storage.Stock, error)
	ListOfStock(storage.StockFilter) ([]storage.Stock, error)
	GetProductAndStockByID(int) (*storage.ProductStock, error)
}
type CoreStock struct { 
	store StockStore
}

func NewCoreStock(ss StockStore) *CoreStock {
	return &CoreStock{
		store: ss,
	}
}

func (cp CoreStock) CreateOrUpdateStock(p storage.Stock) (*storage.Stock, error){
	rs, err :=cp.store.CreateOrUpdateStock(p)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cp CoreStock) ListOfStock(sf storage.StockFilter) ([]storage.Stock, error){
	purchaselist, err:= cp.store.ListOfStock(sf)
	if err != nil{
		return nil, err
	}
	return purchaselist, nil
}
func (s CoreStock) GetProductAndStockByID(pid int) (*storage.ProductStock, error) {
    prores,err := s.store.GetProductAndStockByID(pid)
	if err != nil{
		return nil, err
	}
	return prores, nil
}


