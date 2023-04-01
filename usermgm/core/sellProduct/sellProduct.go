package sellproduct

import (

	"codemen.org/inventory/usermgm/storage"
)


type SellProductStore interface {
	CreateSellProduct(storage.SellProduct) (*storage.SellProduct, error)
	GetProductAndStockByID(int) (*storage.ProductStock, error)
	GetSellProductByID(int) ([]storage.SellProduct, error) 
	
}
type CoreSellProduct struct { 
	store SellProductStore
}

func NewCoreSellProduct(sps SellProductStore) *CoreSellProduct {
	return &CoreSellProduct{
		store: sps,
	}
}

func (csp CoreSellProduct) CreateSellProduct(sp storage.SellProductmap) ([]storage.SoldProduct, error){
    var SoldProducts []storage.SoldProduct
	for k, v := range sp.SelPro {
       pro,_ := csp.store.GetProductAndStockByID(int(k))
	   sellPro := storage.SellProduct{
	   	SellId:         sp.SellId,
	   	ProductId:      pro.ProductId,
	   	Quantity:       int(v),
	   	UnitPrice:      pro.UnitPrice,
	   	TotalUnitPrice: pro.UnitPrice * float32(v),
	   }
	   sellProduct, err := csp.store.CreateSellProduct(sellPro)
	   if err != nil {
		   return nil, err
	   }

	   sP := storage.SoldProduct{
	   	ProductId:      sellProduct.ProductId,
	   	ProductName:    pro.ProductName,
	   	UnitPrice:      sellProduct.UnitPrice,
	   	Quantity:       sellProduct.Quantity,
	   	TotalUnitPrice: sellProduct.TotalUnitPrice,
	   }

       SoldProducts = append(SoldProducts, sP)
	}


	return SoldProducts, nil
}

func (csp CoreSellProduct) GetSellProductByID(id int) ([]storage.SellProduct, error)  { 
	sellproducts, err := csp.store.GetSellProductByID(id)
	if err != nil {
		return nil, err
	}
   
	return sellproducts, nil
}