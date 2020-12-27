package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/rover10/model"
	"github.com/rover10/parseutil"
	"github.com/rover10/querybuilder"
)

func RegisterPatient(tx *gorm.DB, body map[string]interface{}) error {
	required := []string{"accountId", "firstName", "email", "phone", "genderId", "age", "countryId"}
	remove := []string{"uid", "createdOn", "updatedOn"}
	body = parseutil.RemoveFields(body, remove)
	missing := parseutil.EnsureRequired(body, required)
	if len(missing) != 0 {
		log.Println("missing", missing)
		return errors.New("Required fields missing: ")
	}

	stringFields := []string{"accountId", "firstName", "lastName", "phone", "email"}
	intFields := []string{"genderId", "age", "countryId"}
	//jsonFields := []string{"anyExistingMedicalCondition"}
	patient := model.Patient{}
	body, invalidType := parseutil.MapX(body, patient, stringFields, nil, intFields, nil, nil)
	fmt.Println(body)
	if len(invalidType) != 0 {
		log.Println("invalidType", invalidType)
		return errors.New("Invalid field datatype: ")
	}

	// Send to query builder BuildQuery(table string, model map[string]interface{}, returnfields []string)
	query, values := querybuilder.BuildInsertQuery(body, "patient")
	// Camel case can be utilize of RETURNING colum names are supposed to be user instead of table
	query = query + " RETURNING uid, account_id, first_name, last_name, age, email, phone, gender_id, country_id, created_on"

	fmt.Println(query)
	fmt.Println(values)
	// Execute query

	row := tx.Raw(query, values...)
	//lastName := sql.NullString{}
	if tx.Error != nil {
		log.Printf("\nDatabase Error: %+v", tx.Error)
		return errors.New("Database error: ")
	}

	row.Scan(&patient)
	if tx.Error != nil {
		log.Printf("\nDatabase Error: %+v", tx.Error)
		return errors.New("Database error: ")
	}
	return nil
}
