package postgres

import (
	"fmt"
	"log"

	"codemen.org/inventory/usermgm/storage"
)

const createCustomerQuery = `INSERT INTO  customers(
	name
) VALUES(
	:name
) RETURNING *`

func (s PostgresStorage) CreateCustomer(sp storage.Customer) (*storage.Customer, error) {
	stmt, err := s.DB.PrepareNamed(createCustomerQuery)
	if err != nil {
		fmt.Println("errr1")
		return nil, err
	}

	if err := stmt.Get(&sp, sp); err != nil {
		fmt.Println("errr2")
		return nil, err
	}

	if sp.ID == 0 {
		return nil, fmt.Errorf("unable to insert user into db")
	}
	return &sp, nil
}

const listCustomerQuery = `WITH tot AS (select count(*) as total FROM customers
WHERE
	deleted_at IS NULL
	AND (name ILIKE '%%' || $1 || '%%'))
SELECT *, tot.total as total FROM customers
LEFT JOIN tot ON TRUE
WHERE
	deleted_at IS NULL
	AND (name ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
	OFFSET $2
	LIMIT $3`

func (s PostgresStorage) ListOfCustomer(cf storage.CustomerFilter) ([]storage.Customer, error) {
	var listCustomer []storage.Customer
	if cf.Limit == 0 {
		cf.Limit = 15
	}
	if err := s.DB.Select(&listCustomer, listCustomerQuery, cf.SearchTerm, cf.Offset, cf.Limit); err != nil {
		log.Println(err)
		return nil, err
	}
	return listCustomer, nil
}

const getCustomerByIDQuery = `SELECT * FROM customers WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetCustomerByID(id int) (*storage.Customer, error) {
	var customer storage.Customer
	if err := s.DB.Get(&customer, getCustomerByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &customer, nil
}

const updateCustomerQuery = `UPDATE customers SET name = COALESCE(NULLIF(:name, ''), name) WHERE id = :id AND deleted_at IS NULL RETURNING *;`

func (s PostgresStorage) UpdateCustomer(sp storage.Customer) (*storage.Customer, error) {
	stmt, err := s.DB.PrepareNamed(updateCustomerQuery)
	if err != nil {
		log.Fatalln(err)
	}
	if err := stmt.Get(&sp, sp); err != nil {
		log.Println(err)
		return nil, err
	}
	return &sp, nil
}
const deleteCustomerQuery = `UPDATE customers SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) DeleteCustomerByID(id int) error {
	res, err := s.DB.Exec(deleteCustomerQuery, id)
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
		return fmt.Errorf("unable to delete customer")
	}

	return nil
}

const getCustomerByNmaeQuery = `SELECT * FROM customers WHERE name=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetCustomerByName(name string) (*storage.Customer, error) {
	var customer storage.Customer
	if err := s.DB.Get(&customer, getCustomerByNmaeQuery, name); err != nil {
		log.Println(err)
		return nil, err
	}

	return &customer, nil
}