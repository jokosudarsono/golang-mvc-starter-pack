package schemes

import (
	"database/sql"
)

type UserSchema struct {
	ID        int64          `json:"id"`
	FirstName string         `json:"firstname"`
	LastName  sql.NullString `json:"lastname"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}
