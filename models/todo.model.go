package models

import (
	db "todo/core/driver/mysql"
	"todo/schemes"
)

type TodoModel struct{}

func (m *TodoModel) GetTodos(userId int64) (res []map[string]interface{}, e error) {
	query := `
		SELECT
			id,
			user_id,
			title,
			description,
			start_date,
			end_date,
			status,
			created_at,
			updated_at
		FROM todos WHERE user_id = ?
	`

	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		row := schemes.TodoSchema{}

		err := rows.Scan(
			&row.ID,
			&row.UserId,
			&row.Title,
			&row.Description,
			&row.StartDate,
			&row.EndDate,
			&row.Status,
			&row.CreatedAt,
			&row.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		tmp := make(map[string]interface{})
		tmp["id"] = row.ID
		tmp["user_id"] = row.UserId
		tmp["title"] = row.Title
		tmp["description"] = row.Description.String
		tmp["start_date"] = row.StartDate
		tmp["end_date"] = row.EndDate
		tmp["status"] = row.Status
		tmp["created_at"] = row.CreatedAt
		tmp["updated_at"] = row.UpdatedAt.String

		res = append(res, tmp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return res, nil
}

func (m *TodoModel) DetailTodo(id int64, userId int64) (res map[string]interface{}, e error) {
	query := `
		SELECT
			id,
			user_id,
			title,
			description,
			start_date,
			end_date,
			status,
			created_at,
			updated_at
		FROM todos
		WHERE id = ? AND user_id = ?
	`

	row := schemes.TodoSchema{}
	if err := db.DB.QueryRow(query, id, userId).Scan(
		&row.ID,
		&row.UserId,
		&row.Title,
		&row.Description,
		&row.StartDate,
		&row.EndDate,
		&row.Status,
		&row.CreatedAt,
		&row.UpdatedAt,
	); err != nil {
		return nil, err
	}

	tmp := make(map[string]interface{})
	tmp["id"] = row.ID
	tmp["user_id"] = row.UserId
	tmp["title"] = row.Title
	tmp["description"] = row.Description.String
	tmp["start_date"] = row.StartDate
	tmp["end_date"] = row.EndDate
	tmp["status"] = row.Status
	tmp["created_at"] = row.CreatedAt
	tmp["updated_at"] = row.UpdatedAt.String

	return tmp, nil
}

func (m *TodoModel) CreateTodo(
	userId int64,
	title string,
	description string,
	startDate string,
	endDate string,
	status string,
) (lastId int64, e error) {
	query := `
		INSERT INTO todos(
			user_id,
			title,
			description,
			start_date,
			end_date,
			status
		) VALUES(?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(
		userId,
		title,
		description,
		startDate,
		endDate,
		status,
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

func (m *TodoModel) UpdateTodo(
	id int64,
	userId int64,
	title string,
	description string,
	startDate string,
	endDate string,
	status string,
) (affected int64, e error) {
	query := `
		UPDATE todos
		SET
			title = ?,
			description = ?,
			start_date = ?,
			end_date = ?,
			status = ?
		WHERE id = ? AND user_id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(
		title,
		description,
		startDate,
		endDate,
		status,
		id,
		userId,
	)

	if err != nil {
		return 0, err
	}

	affected, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func (m *TodoModel) UpdateStatus(
	id int64,
	userId int64,
	status string,
) (affected int64, e error) {
	query := `
		UPDATE todos
		SET status = ?
		WHERE id = ? AND user_id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, nil
	}

	res, err := stmt.Exec(status, id, userId)
	if err != nil {
		return 0, nil
	}

	affected, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func (m *TodoModel) DeleteTodo(id int64, userId int64) (affected int64, e error) {
	query := "DELETE FROM todos WHERE id = ? AND user_id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(id, userId)
	if err != nil {
		return 0, err
	}

	affected, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
}
