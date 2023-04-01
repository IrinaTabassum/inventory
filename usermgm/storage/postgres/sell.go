package postgres

import (
	"fmt"
	"log"

	"codemen.org/inventory/usermgm/storage"
)

const createSellQuery = `INSERT INTO sells(
	customer_id,
	total_price
) VALUES(
	:customer_id,
	:total_price
) RETURNING *`

func (ps PostgresStorage) CreateSell(s storage.Sell) (storage.Sell, error) {
	stmt, err := ps.DB.PrepareNamed(createSellQuery)
	if err != nil {
		return storage.Sell{}, err
	}

	if err := stmt.Get(&s, s); err != nil {
		return storage.Sell{}, err
	}

	if s.ID == 0 {
		return storage.Sell{}, fmt.Errorf("unable to insert purchase into db")
	}
	return s, nil
}

const listSellQuery = `WITH tot AS (select count(*) as total FROM sells LEFT JOIN customers ON sells.customer_id = customers.id
WHERE (customers.name ILIKE '%%' || $1 || '%%'))
SELECT sells.id as id, sells.customer_id, sells.total_price, sells.created_at, customers.name, tot.total as total FROM sells LEFT JOIN customers ON sells.customer_id = customers.id
LEFT JOIN tot ON TRUE
WHERE (customers.name ILIKE '%%' || $1 || '%%')
	ORDER BY sells.id DESC
	OFFSET $2
	LIMIT $3`

func (s PostgresStorage) ListOfSell(sf storage.SellFilter) ([]storage.Sell, error) {
	var listSell []storage.Sell
	if sf.Limit == 0 {
		sf.Limit = 10
	}
	if err := s.DB.Select(&listSell, listSellQuery, sf.SearchTerm, sf.Offset, sf.Limit); err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println(listSell)
	return listSell, nil
}

const getSellByIDQuery = `
SELECT sells.id as id,
 sells.customer_id, 
 sells.total_price, 
 sells.created_at, 
 customers.name 
 FROM sells LEFT JOIN customers ON sells.customer_id = customers.id WHERE sells.id=$1`

func (s PostgresStorage) GetSellByID(id int) (*storage.Sell, error) {
	var purchase storage.Sell
	if err := s.DB.Get(&purchase, getSellByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &purchase, nil
}

const updateSellQuery = `
    UPDATE sells SET
		total_price = :total_price
	WHERE id = :id RETURNING *;`

func (ps PostgresStorage) UpdateSell(s storage.Sell) (*storage.Sell, error) {
	stmt, err := ps.DB.PrepareNamed(updateSellQuery)
	if err != nil {
		log.Fatalln(err)
	}
	if err := stmt.Get(&s, s); err != nil {
		log.Println(err)
		return nil, err
	}
	return &s, nil
}

