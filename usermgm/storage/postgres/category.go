package postgres

import (
	"fmt"
	"log"

	"codemen.org/inventory/usermgm/storage"
)

const createCategoryQuery = 
`INSERT INTO categories(
	name
) VALUES(
	:name
) RETURNING *`

func (s PostgresStorage) CreateCategory(sc storage.Category) (*storage.Category, error) {
	stmt, err := s.DB.PrepareNamed(createCategoryQuery)
	if err != nil {
		return nil, err
	}
	if err := stmt.Get(&sc, sc); err != nil {
		return nil, err
	}

	if sc.ID == 0 {
		return nil, fmt.Errorf("unable to insert user into db")
	}
	return &sc, nil
}

const listCategoryQuery = `WITH tot AS (select count(*) as total FROM categories
WHERE
	deleted_at IS NULL
	AND (name ILIKE '%%' || $1 || '%%'))
SELECT *, tot.total as total FROM categories
LEFT JOIN tot ON TRUE
WHERE
	deleted_at IS NULL
	AND (name ILIKE '%%' || $1 || '%%')
	ORDER BY id DESC
	OFFSET $2
	LIMIT $3`

func (s PostgresStorage) ListOfCategory(cf storage.CategoryFilter) ([]storage.Category, error) {
	var listCategory []storage.Category
	if cf.Limit == 0 {
		cf.Limit = 15
	}
	if err := s.DB.Select(&listCategory, listCategoryQuery, cf.SearchTerm, cf.Offset, cf.Limit); err != nil {
		log.Println(err)
		return nil, err
	}
	return listCategory, nil
}

const getCategoryByIDQuery = `SELECT * FROM categories WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetCategoryByID(id int) (*storage.Category, error) {
	var category storage.Category
	if err := s.DB.Get(&category, getCategoryByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &category, nil
}

const updateCategoryQuery = `UPDATE categories SET name = COALESCE(NULLIF(:name, ''), name) WHERE id = :id AND deleted_at IS NULL RETURNING *;`

func (s PostgresStorage) UpdateCategory(sp storage.Category) (*storage.Category, error) {
	stmt, err := s.DB.PrepareNamed(updateCategoryQuery)
	if err != nil {
		log.Fatalln(err)
	}
	if err := stmt.Get(&sp, sp); err != nil {
		log.Println(err)
		return nil, err
	}
	return &sp, nil
}
const deleteCategoryQuery = `UPDATE categories SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) DeleteCategoryByID(id int) error {
	res, err := s.DB.Exec(deleteCategoryQuery, id)
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
		return fmt.Errorf("unable to delete Category")
	}

	return nil
}

const getCategorieByNameQuery = `SELECT * FROM categories WHERE name=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetCategoryByName(name string) (*storage.Category, error) {
	var category storage.Category
	if err := s.DB.Get(&category, getCategorieByNameQuery, name); err != nil {
		log.Println(err)
		return nil, err
	}

	return &category, nil
}



