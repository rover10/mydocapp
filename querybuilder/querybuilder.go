package querybuilder

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rover10/mydocapp/model"
)

//BuildInsertQuery build a insert query
func BuildInsertQuery(body map[string]interface{}, table string) (string, []interface{}) {
	attributes := make([]string, 0)
	values := make([]interface{}, 0)
	placeholders := make([]string, 0)
	i := 1

	for k, v := range body {
		//convert camelcase to
		attributes = append(attributes, SnakeToCamelCase(k))
		placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		values = append(values, v)
		i++
	}

	insert := "INSERT INTO "
	insert = insert + table
	insert = insert + " ("
	insert = insert + strings.Join(attributes, ",")
	insert = insert + ")"
	insert = insert + " VALUES"
	insert = insert + "("
	insert = insert + strings.Join(placeholders, ",")
	insert = insert + ")"
	t := reflect.TypeOf(model.Doctor{})
	t.String()
	t.PkgPath()
	return insert, values
}

func BuildUpdateQuery(body map[string]interface{}, table string) (string, []interface{}) {
	updateFields := make([]string, 0)
	values := make([]interface{}, 0)
	//placeholders := make([]string, 0)
	i := 1

	for k, v := range body {
		//convert camelcase to
		updateQr := fmt.Sprintf("%s = $%d", SnakeToCamelCase(k), i)
		updateFields = append(updateFields, updateQr)
		//placeholders = append(placeholders, fmt.Sprintf("$%d", i))
		values = append(values, v)
		i++
	}

	sql := "UPDATE "
	sql = sql + table
	sql = sql + " SET "
	sql = sql + strings.Join(updateFields, ",")

	//t := reflect.TypeOf(model.Doctor{})
	fmt.Println(sql)
	//t.PkgPath()
	return sql, values
}

//SnakeToCamelCase convert snakeCase to camel_case
func SnakeToCamelCase(key string) string {
	camelcase := ""
	for _, y := range key {
		if y > 64 && y < 65+26 {
			camelcase = camelcase + fmt.Sprintf("_%c", y+(97-65))
		} else {
			camelcase = camelcase + fmt.Sprintf("%c", y)
		}
	}
	return camelcase
}
