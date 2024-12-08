package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (db *Date) Insert(user User) (int, int64, error) {
	query := `INSERT INTO user (user_name , email , passwd) 
		VALUES (?, ? , ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return http.StatusInternalServerError, -1, fmt.Errorf("sorry but there are error in server try anther time")
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.User_name, user.Email, user.Passwd)
	if err != nil {
		if strings.Contains(err.Error(), "user_name") {
			return http.StatusFound, -1, fmt.Errorf("user name already used try anther user name")
		} else if strings.Contains(err.Error(), "email") {
			return http.StatusFound, -1, fmt.Errorf("email already used try anther email")
		}
		return http.StatusInternalServerError, -1, fmt.Errorf("sorry but there are error in server try anther time")
	}

	lastID, err := res.LastInsertId()
	if err == nil {
		return http.StatusOK, lastID, nil
	}
	return http.StatusInternalServerError, -1, fmt.Errorf("sorry but there are error in server try anther time")
}

func (db *Date) Select(userIfo, passwd string) (int, error) {
	query := `SELECT id , passwd FROM user
		WHERE user_name = ? OR email = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return -2, fmt.Errorf("problem in server try anther time")
	}
	defer stmt.Close()
	var hashpasswd string
	var id int
	err = stmt.QueryRow(userIfo, userIfo).Scan(&id, &hashpasswd)
	if err == sql.ErrNoRows {
		return -1, fmt.Errorf("user or password not correct")
	} else if err != nil {
		return -1, fmt.Errorf("check your input")
	}
	if bcrypt.CompareHashAndPassword([]byte(hashpasswd), []byte(passwd)) != nil {
		return -1, fmt.Errorf("user or password not correct")
	}
	return id, nil
}

func (db *Date) CheckEXist(checker string) bool {
	exist := false
	err := db.DB.QueryRow(`
		SELECT EXISTS(
			SELECT 1 
			FROM session
			WHERE uid = ?
		)`, checker, checker).Scan(&exist)
	return err == nil && exist
}

func (db *Date) CraeteSession(userid int, session string) error {
	query := `INSERT INTO session(user_id , uid)
		VALUES(?,?)
			ON CONFLICT DO UPDATE SET
				uid = EXCLUDED.uid ,
				create_date = CURRENT_TIMESTAMP
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userid, session)

	return err
}

func (db *Date) DeleteSession() {
	diff_time := time.Now().Add(-time.Hour * 24)
	query := `SELECT id, create_date FROM session`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	tx, err := db.DB.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tx.Rollback()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var lasttime time.Time
	id := 0
	for rows.Next() {
		err = rows.Scan(&id, &lasttime)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if lasttime.Second() <= diff_time.Second() {
			_, err = tx.Exec(`DELETE FROM session WHERE id = ?`, id)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	if err := tx.Commit(); err != nil {
		fmt.Println("Error committing transaction:", err)
	}
}
