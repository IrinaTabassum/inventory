package postgres

import (
	"fmt"
	"log"

	"codemen.org/inventory/usermgm/storage"
)

const createPurchaseQuery = `INSERT INTO  purchases(
	supplier_id,
	product_id,
	quantity,
	unit_price,
	total_price
) VALUES(
	:supplier_id,
	:product_id,
	:quantity,
	:unit_price,
	:total_price
) RETURNING *`

func (s PostgresStorage) CreatePurchase(sp storage.Purchase) (*storage.Purchase, error) {
	stmt, err := s.DB.PrepareNamed(createPurchaseQuery)
	if err != nil {
		return nil, err
	}

	if err := stmt.Get(&sp, sp); err != nil {
		return nil, err
	}

	if sp.ID == 0 {
		return nil, fmt.Errorf("unable to insert purchase into db")
	}
	return &sp, nil
}

const listPurchaseQuery = `WITH tot AS (select count(*) as total FROM purchases as p LEFT JOIN  suppliers as s ON p.supplier_id = s.id
WHERE (name ILIKE '%%' || $1 || '%%'))
SELECT p.id, p.supplier_id, s.name, 
p.product_id, p.unit_price, p.quantity, 
p.total_price, p.created_at,p.updated_at, tot.total as total
FROM purchases as p LEFT JOIN  suppliers as s ON p.supplier_id = s.id
LEFT JOIN tot ON TRUE
WHERE (name ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
	OFFSET $2
	LIMIT $3`

func (s PostgresStorage) ListOfPurchase(pf storage.PurchaseFilter) ([]storage.Purchase, error) {
	var listPurchase []storage.Purchase
	if pf.Limit == 0 {
		pf.Limit = 10
	}
	if err := s.DB.Select(&listPurchase, listPurchaseQuery, pf.SearchTerm, pf.Offset, pf.Limit); err != nil {
		log.Println(err)
		return nil, err
	}
	return listPurchase, nil
}

const getPurchaseByIDQuery = `SELECT * FROM purchases WHERE id=$1`

func (s PostgresStorage) GetPurchaseByID(id int) (*storage.Purchase, error) {
	var purchase storage.Purchase
	if err := s.DB.Get(&purchase, getPurchaseByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &purchase, nil
}

const updatePurchaseQuery = `
    UPDATE purchases SET
		supplier_id = :supplier_id,
		product_id = :product_id,
		quantity = :quantity,
		unit_price = :unit_price,
		total_price = :total_price
	WHERE id = :id RETURNING *;`

func (s PostgresStorage) UpdatePurchase(sp storage.Purchase) (*storage.Purchase, error) {
	stmt, err := s.DB.PrepareNamed(updatePurchaseQuery)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := stmt.Get(&sp, sp); err != nil {
		log.Println(err)
		return nil, err
	}
	
	return &sp, nil
	
}
