package schemes

import (
	"database/sql"
)

type TodoSchema struct {
	ID          int64          `json:"id"`
	UserId      int64          `json:"user_id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	StartDate   string         `json:"start_date"`
	EndDate     string         `json:"end_date"`
	Status      string         `json:"status"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   sql.NullString `json:"updated_at"`
}
