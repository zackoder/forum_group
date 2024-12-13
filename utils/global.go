package utils

import (
	"database/sql"
	"errors"
)

var DB *sql.DB

func GetUserByToken(token string) (int, error) {
	var userId int
	getUserIdQuery := `SELECT user_id FROM sessions WHERE token=?;`
	queryErr := DB.QueryRow(getUserIdQuery, token).Scan(&userId)
	if queryErr != nil {
		return userId, errors.New("user not valid")
	}
	return userId, nil
}
