package querybuilder

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rover10/mydocapp.git/model"
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
