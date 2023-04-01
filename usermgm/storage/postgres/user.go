package postgres

import (
	"fmt"
	"log"

	"codemen.org/inventory/usermgm/storage"
)

const registerQuery = `INSERT INTO users (
	first_name,
	last_name,
	username,
	email,
	password
) VALUES(
	:first_name,
	:last_name,
	:username,
	:email,
	:password
) RETURNING *`

func (s PostgresStorage) Register(u storage.User) (*storage.User, error) {
	stmt, err := s.DB.PrepareNamed(registerQuery)
	if err != nil {
		return nil, err
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("unable to insert user into db")
	}
	return &u, nil
}

const getUserByUsernameQuery = `SELECT * FROM users WHERE username=$1 AND deleted_at IS NULL`

func (ps PostgresStorage) GetUserByUsername(usernanme string) (*storage.User, error) {
	var user storage.User
	if err := ps.DB.Get(&user, getUserByUsernameQuery, usernanme); err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("unable to find user by username")
	}

	return &user, nil
}

const listQuery = `
	WITH tot AS (select count(*) as total FROM users
	WHERE
		deleted_at IS NULL
		AND (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%' OR username ILIKE '%%' || $1 || '%%' OR email ILIKE '%%' || $1 || '%%'))
	SELECT *, tot.total as total FROM users
	LEFT JOIN tot ON TRUE

	WHERE
		deleted_at IS NULL
		AND (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%' OR username ILIKE '%%' || $1 || '%%' OR email ILIKE '%%' || $1 || '%%')
		ORDER BY id DESC
		OFFSET $2
		LIMIT $3`

func (s PostgresStorage) ListUser(uf storage.UserFilter) ([]storage.User, error) {
	var listUser []storage.User
	if uf.Limit == 0 {
		uf.Limit = 15
	}
	if err := s.DB.Select(&listUser, listQuery, uf.SearchTerm, uf.Offset, uf.Limit); err != nil {
		log.Println(err)
		return nil, err
	}

	return listUser, nil
}

const updateUserQuery = `
	UPDATE users SET
		first_name = COALESCE(NULLIF(:first_name, ''), first_name),
		last_name = COALESCE(NULLIF(:last_name, ''), last_name),
	    is_active = COALESCE(:is_active, is_active)
	WHERE id = :id AND deleted_at IS NULL RETURNING *;
	`

func (s PostgresStorage) UpdateUser(u storage.User) (*storage.User, error) {
	stmt, err := s.DB.PrepareNamed(updateUserQuery)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := stmt.Get(&u, u); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const deleteUserByIdQuery = `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) DeleteUserByID(id int) (storage.User, error) {
	res, err := s.DB.Exec(deleteUserByIdQuery, id)
	if err != nil {
		fmt.Println(err)
		return storage.User{}, err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return storage.User{}, err
	}

	if rowCount <= 0 {
		return storage.User{}, fmt.Errorf("unable to delete user")
	}

	return storage.User{}, nil
}

const getUserByIDQuery = `SELECT * FROM users WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetUserByID(id int) (*storage.User, error) {
	var u storage.User
	if err := s.DB.Get(&u, getUserByIDQuery, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}

const getUserByEmailQuery = `SELECT * FROM users WHERE email=$1 AND deleted_at IS NULL`

func (ps PostgresStorage) GetUserByEmail(email string) (*storage.User, error) {
	var user storage.User
	if err := ps.DB.Get(&user, getUserByEmailQuery, email); err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("unable to find user by email")
	}

	return &user, nil
}