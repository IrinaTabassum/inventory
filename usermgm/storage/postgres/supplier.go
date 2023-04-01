package postgres

import (
	"fmt"
	"log"

	"codemen.org/inventory/usermgm/storage"
)

const createSupplierQuery = `INSERT INTO  suppliers(
	name
) VALUES(
	:name
) RETURNING *`

func (s PostgresStorage) CreateSupplier(sp storage.Supplier) (*storage.Supplier, error) {
	stmt, err := s.DB.PrepareNamed(createSupplierQuery)
	if err != nil {
		return nil, err
	}

	if err := stmt.Get(&sp, sp); err != nil {
		return nil, err
	}

	if sp.ID == 0 {
		return nil, fmt.Errorf("unable to insert user into db")
	}
	return &sp, nil
}

const listSupplierQuery = `WITH tot AS (select count(*) as total FROM suppliers
WHERE
	deleted_at IS NULL
	AND (name ILIKE '%%' || $1 || '%%'))
SELECT *, tot.total as total FROM suppliers
LEFT JOIN tot ON TRUE
WHERE
	deleted_at IS NULL
	AND (name ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
	OFFSET $2
	LIMIT $3`

func (s PostgresStorage) ListOfSupplier(sf storage.SupplierFilter) ([]storage.Supplier, error) {
	var listSupplier []storage.Supplier
	if sf.Limit == 0 {
		sf.Limit = 15
	}
	if err := s.DB.Select(&listSupplier, listSupplierQuery, sf.SearchTerm, sf.Offset, sf.Limit); err != nil {
		log.Println(err)
		return nil, err
	}
	return listSupplier, nil
}

const getSupplierByIDQuery = `SELECT * FROM suppliers WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetSupplierByID(id int) (*storage.Supplier, error) {
	var supplier storage.Supplier
	if err := s.DB.Get(&supplier, getSupplierByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &supplier, nil
}

const getSupplierByNameQuery = `SELECT * FROM suppliers WHERE name=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetSupplierByName(name string) (*storage.Supplier, error) {
	var supplier storage.Supplier
	if err := s.DB.Get(&supplier, getSupplierByNameQuery, name); err != nil {
		log.Println(err)
		return nil, err
	}

	return &supplier, nil
}

const updateSupplierQuery = `
	UPDATE suppliers SET
		name = COALESCE(NULLIF(:name, ''), name)
	WHERE id = :id AND deleted_at IS NULL RETURNING *;
	`

func (s PostgresStorage) UpdateSupplier(sp storage.Supplier) (*storage.Supplier, error) {
	stmt, err := s.DB.PrepareNamed(updateSupplierQuery)
	if err != nil {
		log.Fatalln(err)
	}
	if err := stmt.Get(&sp, sp); err != nil {
		log.Println(err)
		return nil, err
	}
	return &sp, nil
}
const deleteSupplierQuery = `UPDATE suppliers SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) DeleteSupplierByID(id int) error {
	res, err := s.DB.Exec(deleteSupplierQuery, id)
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
		return fmt.Errorf("unable to delete supplier")
	}

	return nil
}

