package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum/utils"
)

// var MyDB *sql.DB

// type test struct {
// 	Name string
// 	City string
// }

func Comments(w http.ResponseWriter, r *http.Request) {
	query := `SELECT * FROM users;`
	data, _ := utils.DB.Query(query)
	var users []utils.User
	for data.Next() {
		var user utils.User
		if err := data.Scan(&user.Id, &user.Username, &user.Email, &user.Password); err != nil {
			fmt.Println(err)
			return
		}
		users = append(users, user)
	}

	// t := &test{Name: "yassin", City: "oujda"}
	json.NewEncoder(w).Encode(users)
}
