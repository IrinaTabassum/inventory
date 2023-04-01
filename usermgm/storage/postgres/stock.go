package postgres

import (
	"fmt"
	"log"

	"codemen.org/inventory/usermgm/storage"
)

const createStockQuery = `INSERT INTO  stocks (
	product_id,
	quantity
) VALUES(
	:product_id,
	:quantity
)ON CONFLICT(product_id) DO UPDATE SET 
product_id = EXCLUDED.product_id, 
quantity = EXCLUDED.quantity RETURNING *`

func (s PostgresStorage) CreateOrUpdateStock(sp storage.Stock) (*storage.Stock, error) {
	stmt, err := s.DB.PrepareNamed(createStockQuery)
	if err != nil {
		return nil, err
	}

	if err := stmt.Get(&sp, sp); err != nil {
		return nil, err
	}

	if sp.ID == 0 {
		return nil, fmt.Errorf("unable to insert stock into db")
	}
	return &sp, nil
}

const listStockQuery = `
WITH tot AS (select count(*) as total FROM stocks 
LEFT JOIN products ON stocks.product_id = products.id
WHERE (name ILIKE '%%' || $1 || '%%'))
SELECT s.id, s.product_id, products.name, s.quantity, s.created_at, s.updated_at, tot.total as total 
FROM stocks as s LEFT JOIN products ON s.product_id = products.id
LEFT JOIN tot ON TRUE
WHERE (name ILIKE '%%' || $1 || '%%')
	ORDER BY products.id DESC
	OFFSET $2
	LIMIT $3`

func (s PostgresStorage) ListOfStock(sf storage.StockFilter) ([]storage.Stock, error) {
	var listStock []storage.Stock
	if sf.Limit == 0 {
		sf.Limit = 10
	}
	if err := s.DB.Select(&listStock, listStockQuery, sf.SearchTerm, sf.Offset, sf.Limit); err != nil {
		log.Println(err)
		return nil, err
	}
	return listStock, nil
}


const getProductQuanAndStockBYID = `select COALESCE(stocks.id, 0) as id,  products.id as product_id, products.name, products.price, COALESCE(stocks.quantity, 0) as quantity FROM products
LEFT JOIN stocks
ON products.id = stocks.product_id where products.id=$1;`

func (s PostgresStorage) GetProductAndStockByID(id int) (*storage.ProductStock, error) {
	var productStock storage.ProductStock
	if err := s.DB.Get(&productStock, getProductQuanAndStockBYID, id); err != nil {
		log.Println(err)
		return nil, err
	}
	
	return &productStock, nil
}
