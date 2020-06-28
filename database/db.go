package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/rover10/mydocapp.git/model"
)

type DocDB struct {
	DB *sql.DB
}

func (db *DocDB) RetriveUserCred(email string) (model.User, error) {
	query := " SELECT uid, password, first_name, email, phone, user_type, country_id, is_active from users where email = $1"
	fmt.Println(query)
	row := db.DB.QueryRow(query, email)
	user := model.User{}
	err := row.Scan(&user.UID, &user.Password, &user.FirstName, &user.Email, &user.Phone, &user.UserType, &user.Country, &user.IsActive)
	if err != nil {
		log.Printf("\nDatabase Error: %+v", err)
		return user, err
	}
	return user, nil
}
