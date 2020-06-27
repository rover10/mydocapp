package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/rover10/mydocapp.git/querybuilder"
	uuid "github.com/satori/go.uuid"
)

func main() {
	dr := Doctor{}
	//fmt.Println(dr.main.D)
	nullTypes := map[string]interface{}{
		"*bool":      "sql.NullBool",
		"*float64":   "sql.NullFloat64",
		"*int32":     "sql.NullInt32",
		"*int64":     "sql.NullInt64",
		"*int":       "sql.NullInt64",
		"*string":    "sql.NullString",
		"*uuid.UUID": "sql.NullString",
		//"*": "sql.NullTime"
	}

	primitiveTypes := map[string]interface{}{
		"bool":      true,
		"float64":   true,
		"int32":     true,
		"int64":     true,
		"int":       true,
		"string":    true,
		"uuid.UUID": true,
		//"*": "sql.NullTime"
	}

	//insert-required:true
	//post-required:true
	//update-required:true
	//readonly
	//return all the fields in all the case
	//tag := f.Tag.Get("insert-required")
	insertRequired := make([]string, 1)
	insertRequired[0] = "0"
	//tag := f.Tag.Get("post-require")

	//tag := f.Tag.Get("update-required")
	updateRequired := make([]string, 1)
	updateRequired[0] = "0"
	//writeIgnore:"True"
	writeIgnore := make([]string, 1)
	writeIgnore[0] = "0"
	//return fields
	returnFields := make([]string, 0)

	// Scan values
	// snake := convertToSnake(tag)
	// generate get returning all fields
	// generate post
	// generate put
	// Mark fields which does not need serialization
	t := reflect.TypeOf(dr)
	modelVariable := "model"

	//scanline reads all fi
	scanStatement := "err = row.Scan("
	model := fmt.Sprintf("%s := %s{}", modelVariable, t)
	fmt.Println(model)
	//nullableVariableDeclartion
	nullableVariableDeclartion := make([]string, 0)
	scanStatementArgs := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name)

		tf := fmt.Sprintf("%s", f.Type)
		key := fmt.Sprintf("%s, %s, %s", f.Name, f.Type, nullTypes[tf])
		insertReq := f.Tag.Get("insert-required")
		if insertReq == "True" {
			insertRequired = append(insertRequired, tf)
		}

		updateReq := f.Tag.Get("update-required")
		if updateReq == "True" {
			updateRequired = append(updateRequired, tf)
		}

		//
		//For pointer type NullType
		//textT := fmt.Sprintf("%s", tf)
		//fmt.Println("\ntext = %s", t)
		//fmt.Println(fmt.Sprintf("\n---->>>>>%s", textT))
		jsonField := querybuilder.SnakeToCamelCase(f.Tag.Get("json"))
		if strings.Contains(key, "*") {
			sqlNullType := nullTypes[tf]
			if sqlNullType != nil {
				nullableVariable := fmt.Sprintf("%s := %s{}", jsonField, sqlNullType)
				println("--->>>" + nullableVariable)
				returnFields = append(returnFields, jsonField)
				nullableVariableDeclartion = append(nullableVariableDeclartion, nullableVariable)
				scanStatement = fmt.Sprintf("&%s ", f.Tag.Get("json"))
				scanStatementArgs = append(scanStatementArgs, scanStatement)
			}

		} else {
			//templateCode := fmt.Sprintf("%s := %s{}", f.Tag.Get("json"), tf)
			//println(templateCode)
			primType := primitiveTypes[tf]
			if primType != nil {
				returnFields = append(returnFields, jsonField)
				scanStatement = fmt.Sprintf("&%s.%s ", modelVariable, f.Name)
				scanStatementArgs = append(scanStatementArgs, scanStatement)
			}
		}

		//
		// excluded fields for PUT,POST and GET
		// required: required tag
		// datatype categorization int, float64, string, boolean
		// other primitive type should be read with sql.Null* datatype
		// Any non primitive type should be handled manually in version 1
		// jsonb should be handled manually
		// scan
		// Value assignment to struct

	}

	// Query builder
	modelName := strings.Split(t.Name(), ".")
	table := t.Name()
	if len(modelName) == 2 {
		table = modelName[1]
	}
	table = querybuilder.SnakeToCamelCase(table)
	table = strings.TrimLeft(table, "_")
	queryBuilderStmt := fmt.Sprintf("query, values := querybuilder.BuildInsertQuery(body, \"%s\")", table)
	fmt.Println(queryBuilderStmt)
	//Return value
	returnFieldsStatement := "RETURNING " + strings.Join(returnFields, ",")
	fmt.Println(returnFieldsStatement)
	// Execute SQL
	fmt.Println(EXECUTE_SQL)
	//Nullable variable declaration
	for _, v := range nullableVariableDeclartion {
		fmt.Println(v)
	}
	//Scan query result
	scanResult := "err = row.Scan(" + strings.Join(scanStatementArgs, ",") + ")"
	fmt.Println(scanResult)
	// After scan, commit
	fmt.Println(AFTER_RESULT_SCAN)
	fmt.Println("--->")
}

type Doctor struct {
	AccountID  uuid.UUID  `json:"accountId"`
	AccountID2 *uuid.UUID `json:"accountId2"`
	Fee        *float64   `json:"float64" required:"True" writeIgnore:"True"`
	Tee        *int64     `json:"int64"`
	See        *string    `json:"string"`
	Int2       *int       `json:"intPtr"`
	Float32f   *int       `json:"f32"`
	Integer    int        `json:"integer"`
	Dd         D          `json:"dd"`
	Dd2        *D         `json:"ddPtr"`
}

type D struct {
	A string `json:"a"`
}

const (
	EXECUTE_SQL = `
tx, err := s.DB.Begin()
if err != nil {
	return context.JSON(http.StatusInternalServerError, err)
}
row := tx.QueryRow(query, values...)
	`
	AFTER_RESULT_SCAN = `	
if err != nil {
	log.Printf("\nDatabase Error: %+v", err)
	return context.JSON(http.StatusInternalServerError, err)
}
err = tx.Commit()
if err != nil {
	log.Printf("\nDatabase Commit Error: %+v", err)
	return context.JSON(http.StatusInternalServerError, err)
}
return context.JSON(http.StatusOK, model)
	`
)

/*
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32
     // represents a Unicode code point
float32 float64
complex64 complex128
*/
