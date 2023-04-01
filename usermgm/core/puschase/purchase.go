package puschase

import (
	"log"

	"codemen.org/inventory/usermgm/storage"
)

type PurchaseStore interface {
	CreatePurchase(storage.Purchase) (*storage.Purchase, error)
	ListOfPurchase(storage.PurchaseFilter) ([]storage.Purchase, error)
	GetPurchaseByID(int) (*storage.Purchase, error)
	UpdatePurchase(storage.Purchase) (*storage.Purchase, error)
	CreateOrUpdateStock(storage.Stock) (*storage.Stock, error)
	GetProductAndStockByID(int) (*storage.ProductStock, error)
}

type CorePurchase struct {
	store PurchaseStore
}

func NewCorePurchase(ps PurchaseStore) *CorePurchase {
	return &CorePurchase{
		store: ps,
	}
}

func (cp CorePurchase) CreatePurchase(p storage.Purchase) (*storage.Purchase, error) {
	pro, err := cp.store.GetProductAndStockByID(p.ProductId)
	if err != nil {
		log.Println(err)
	}

	pur := storage.Purchase{
		SupplierId: p.SupplierId,
		ProductId:  p.ProductId,
		Quantity:   p.Quantity,
		UnitPrice:  pro.UnitPrice,
		TotalPrice: pro.UnitPrice * float32(p.Quantity),
	}
	purchase, err := cp.store.CreatePurchase(pur)
	if err != nil {
		return nil, err
	}
	var quantity int
	if pro == nil {
		quantity = p.Quantity
	} else {
		quantity = pro.StockQuantity + p.Quantity
	}
	stok := storage.Stock{
		ProductId: p.ProductId,
		Quantity:  quantity,
	}
	_, err = cp.store.CreateOrUpdateStock(stok)
	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (cp CorePurchase) ListOfPurchase(sf storage.PurchaseFilter) ([]storage.Purchase, error) {
	purchaselist, err := cp.store.ListOfPurchase(sf)
	if err != nil {
		return nil, err
	}
	return purchaselist, nil
}

func (cp CorePurchase) GetPurchaseByID(id int) (*storage.Purchase, error) {

	rs, err := cp.store.GetPurchaseByID(id)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (cp CorePurchase) UpdatePurchase(sp storage.Purchase) (*storage.Purchase, error) {
	oldPurch, err := cp.store.GetPurchaseByID(sp.ID)
	if err != nil {
		log.Println(err)
	}
	oldProduct, err := cp.store.GetProductAndStockByID(oldPurch.ProductId)
	if err != nil {
		log.Println(err)
	}

	stok := storage.Stock{
		ProductId: oldPurch.ProductId,
		Quantity:  oldProduct.StockQuantity - oldPurch.Quantity,
	}
	_, err = cp.store.CreateOrUpdateStock(stok)
	if err != nil {
		return nil, err
	}

	newProduct, err := cp.store.GetProductAndStockByID(sp.ProductId)
	if err != nil {
		log.Println(err)
	}
	pur := storage.Purchase{
		ID:         sp.ID,
		SupplierId: sp.SupplierId,
		ProductId:  sp.ProductId,
		Quantity:   sp.Quantity,
		UnitPrice:  newProduct.UnitPrice,
		TotalPrice: newProduct.UnitPrice * float32(sp.Quantity),
	}

	updatePur, err := cp.store.UpdatePurchase(pur)
	if err != nil {
		return nil, err
	}

	stok = storage.Stock{
		ProductId: sp.ProductId,
		Quantity:  newProduct.StockQuantity + sp.Quantity,
	}
	_, err = cp.store.CreateOrUpdateStock(stok)
	if err != nil {
		return nil, err
	}

	return updatePur, nil
}
