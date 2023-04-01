package sell

import (
	"codemen.org/inventory/usermgm/storage"
)



type SellStore interface {
	CreateSell(storage.Sell) (storage.Sell, error)
	GetSellByID(int) (*storage.Sell, error)
	UpdateSell(storage.Sell) (*storage.Sell, error)
	ListOfSell(storage.SellFilter) ([]storage.Sell, error)
}
type CoreSell struct { 
	store SellStore
}

func NewCoreSell(ss SellStore) *CoreSell {
	return &CoreSell{
		store: ss,
	}
}

func (ps CoreSell) CreateSell(s storage.Sell) (storage.Sell, error){
	sell, err := ps.store.CreateSell(s)
	if err != nil {
		return storage.Sell{}, err
	}

	return sell, nil
}

func (ps CoreSell) UpdateSell(s storage.Sell) (*storage.Sell, error){
	sell, err := ps.store.UpdateSell(s)
	if err != nil {
		return nil, err
	}

	return sell, nil
}

func (ps CoreSell) ListOfSell(pf storage.SellFilter) ([]storage.Sell, error){
	productlist, err:= ps.store.ListOfSell(pf)
	if err != nil{
		return nil, err
	}
	return productlist, nil
}

func (ps CoreSell) GetSellByID(id int) (*storage.Sell, error){
	sell, err := ps.store.GetSellByID(id)
	if err != nil {
		return nil, err
	}

	return sell, nil
}



