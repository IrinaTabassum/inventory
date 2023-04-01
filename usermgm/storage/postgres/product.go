package postgres

import (
	"fmt"
	"log"

	"codemen.org/inventory/usermgm/storage"
)

const createProductQuery = `INSERT INTO  products(
	category_id,
	name,
	description,
	price
) VALUES(
	:category_id,
	:name,
	:description,
	:price
) RETURNING *`

func (s PostgresStorage) CreateProduct(sp storage.Product) (*storage.Product, error) {
	stmt, err := s.DB.PrepareNamed(createProductQuery)
	if err != nil {
		return nil, err
	}

	if err := stmt.Get(&sp, sp); err != nil {
		return nil, err
	}

	if sp.ID == 0 {
		return nil, fmt.Errorf("unable to insert product into db")
	}
	return &sp, nil
}

const listProductQuery = `WITH tot AS (select count(*) as total FROM products WHERE deleted_at IS NULL AND (name ILIKE '%%' || $1 || '%%'))
SELECT P.id, P.name, P.category_id, categories.name as category_name,  P.description, P.price, COALESCE(stocks.quantity, 0) as quantity, P.deleted_at, tot.total as total
FROM products as P LEFT JOIN stocks ON P.id = stocks.product_id LEFT JOIN categories ON P.category_id = categories.id
LEFT JOIN tot ON TRUE
WHERE p.deleted_at IS NULL AND (p.name ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
	OFFSET $2
	LIMIT $3`

func (s PostgresStorage) ListOfProduct(pf storage.ProductFilter) ([]storage.Product, error) {
	var listProduct []storage.Product
	if(pf.Limit == 0){
		pf.Limit = 10
	}
	if err := s.DB.Select(&listProduct, listProductQuery, pf.SearchTerm, pf.Offset, pf.Limit); err != nil {
		return nil, err
	}
	
	return listProduct, nil
}

const getProductByIDQuery = `SELECT * FROM products WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetProductByID(id int) (*storage.Product, error) {
	var product storage.Product
	if err := s.DB.Get(&product, getProductByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &product, nil
}

const updateProductQuery = `
	UPDATE products SET
		name = :name,
		category_id = :category_id,
		description = :description,
		price = :price
	WHERE id = :id AND deleted_at IS NULL RETURNING *;
	`

func (s PostgresStorage) UpdateProduct(sp storage.Product) (*storage.Product, error) {
	stmt, err := s.DB.PrepareNamed(updateProductQuery)
	if err != nil {
		log.Fatalln(err)
	}
	if err := stmt.Get(&sp, sp); err != nil {
		log.Println(err)
		return nil, err
	}
	return &sp, nil
}

const deleteProductQuery = `UPDATE products SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) DeleteProductByID(id int) error {
	res, err := s.DB.Exec(deleteProductQuery, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return fmt.Errorf("unable to delete product")
	}

	return nil
}
