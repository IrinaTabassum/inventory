package postgres

import (
	"fmt"
	"log"

	"codemen.org/inventory/usermgm/storage"
)

const createSellProductQuery = `INSERT INTO  sell_products (
	sell_id,
	product_id,
	quantity,
	unit_price,
	total_unit_price
) VALUES(
	:sell_id,
	:product_id,
	:quantity,
	:unit_price,
	:total_unit_price
) RETURNING *`

func (s PostgresStorage) CreateSellProduct(sp storage.SellProduct) (*storage.SellProduct, error) {

	stmt, err := s.DB.PrepareNamed(createSellProductQuery)
	if err != nil {
		return nil, err
	}
	if err := stmt.Get(&sp, sp); err != nil {
		return nil, err
	}

	if sp.ID == 0 {
		return nil, fmt.Errorf("unable to insert Sell purchase into db")
	}
	return &sp, nil
}

const getSellProductBYID =
`select p.id as product_id, p.name, sell_products.sell_id, sell_products.quantity, sell_products.unit_price, sell_products.total_unit_price 
FROM sell_products LEFT JOIN products p on p.id = sell_products.product_id
where sell_id in ($1)`

func (s PostgresStorage) GetSellProductByID(id int) ([]storage.SellProduct, error) {
	var sellProduct []storage.SellProduct
	if err := s.DB.Select(&sellProduct, getSellProductBYID, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return sellProduct, nil
}