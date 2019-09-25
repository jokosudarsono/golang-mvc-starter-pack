package models

import (
	"todo/core/driver/mysql"
	"todo/schemes"
)

type AuthModel struct{}

func (m *AuthModel) Signup(
	firstname string,
	lastname string,
	email string,
	password string,
) (lastId int64, err error) {
	query := `
		INSERT INTO users(
			firstname,
			lastname,
			email,
			password
		) VALUES(?, ?, ?, ?)
	`

	stmt, err := mysql.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(
		firstname,
		lastname,
		email,
		password,
	)

	if err != nil {
		return 0, err
	}

	lastId, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (m *AuthModel) GetUserByEmail(email string) (res map[string]interface{}, e error) {
	query := `
		SELECT
			id,
			firstname,
			lastname,
			email,
			password,
			created_at,
			updated_at
		FROM users WHERE email = ?
	`

	row := schemes.UserSchema{}

	if err := mysql.DB.QueryRow(query, email).Scan(
		&row.ID,
		&row.FirstName,
		&row.LastName,
		&row.Email,
		&row.Password,
		&row.CreatedAt,
		&row.UpdatedAt,
	); err != nil {
		return nil, err
	}

	tmp := make(map[string]interface{})
	tmp["id"] = row.ID
	tmp["firstname"] = row.FirstName
	tmp["lastname"] = row.LastName.String
	tmp["email"] = row.Email
	tmp["password"] = row.Password
	tmp["created_at"] = row.CreatedAt
	tmp["updated_at"] = row.UpdatedAt.String

	return tmp, nil
}
