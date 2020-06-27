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
		"bool":       "bool",
		"float64":    "float64",
		"int32":      "int",
		"int64":      "int",
		"int":        "int",
		"string":     "string",
		"uuid.UUID":  "string",
		"*bool":      "bool",
		"*float64":   "float64",
		"*int32":     "int",
		"*int64":     "int",
		"*int":       "int",
		"*string":    "string",
		"*uuid.UUID": "string",
		//"*": "sql.NullTime"
	}

	//create-required:true
	//update-required:true

	//return all the fields in all the case
	//tag := f.Tag.Get("insert-required")
	createRequired := make([]string, 0)
	createRemove := make([]string, 0)
	//tag := f.Tag.Get("post-require")
	//tag := f.Tag.Get("update-required")
	updateRequired := make([]string, 0)
	updateRemove := make([]string, 0)

	//return fields
	returnFields := make([]string, 0)
	// Scan values
	// Mark fields which does not need serialization
	t := reflect.TypeOf(dr)
	modelVariable := "model"
	stringField := make([]string, 0)
	boolField := make([]string, 0)
	intField := make([]string, 0)
	floatField := make([]string, 0)
	jsonField := make([]string, 0)
	//scanline reads all fi
	scanStatement := ""
	//nullableVariableDeclartion
	nullableVariableDeclartion := make([]string, 0)
	scanStatementArgs := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name)
		tf := fmt.Sprintf("%s", f.Type)
		key := fmt.Sprintf("%s, %s, %s", f.Name, f.Type, nullTypes[tf])

		// Create
		createReq := f.Tag.Get("create-required")
		createRem := f.Tag.Get("create-remove")
		if createRem == "True" && createReq == "True" {
			fmt.Println("Ambiguous requirement, either of 'create-required' or 'create-remove' should be applied at time.")
			return
		} else if createReq == "True" {
			createRequired = append(createRequired, tf)
		} else if createRem == "True" {
			createRemove = append(createRemove, tf)
		}

		// Update
		updateReq := f.Tag.Get("update-required")
		updateRem := f.Tag.Get("update-remove")
		if updateReq == "True" && updateRem == "True" {
			fmt.Println("Ambiguous requirement, either of 'update-required' or 'update-remove' should be applied at time.")
			return
		} else if updateReq == "True" {
			updateRequired = append(updateRequired, tf)
		} else if updateRem == "True" {
			updateRemove = append(updateRemove, tf)
		}
		//For pointer type NullType
		//textT := fmt.Sprintf("%s", tf)
		//fmt.Println("\ntext = %s", t)
		//fmt.Println(fmt.Sprintf("\n---->>>>>%s", textT))
		jsonField := f.Tag.Get("json")
		dataType := primitiveTypes[tf]
		if dataType != "int" {
			intField = append(intField, jsonField)
		} else if dataType != "float64" {
			floatField = append(floatField, jsonField)
		} else if dataType != "string" {
			stringField = append(stringField, jsonField)
		} else if dataType != "bool" {
			boolField = append(boolField, jsonField)
		}
		queryReturnField := querybuilder.SnakeToCamelCase(f.Tag.Get("json"))
		if strings.Contains(key, "*") {
			sqlNullType := nullTypes[tf]
			if sqlNullType != nil {
				nullableVariable := fmt.Sprintf("%s := %s{}", jsonField, sqlNullType)
				println("--->>>" + nullableVariable)
				returnFields = append(returnFields, queryReturnField)
				nullableVariableDeclartion = append(nullableVariableDeclartion, nullableVariable)
				scanStatement = fmt.Sprintf("&%s ", f.Tag.Get("json"))
				scanStatementArgs = append(scanStatementArgs, scanStatement)
			}
		} else {
			//templateCode := fmt.Sprintf("%s := %s{}", f.Tag.Get("json"), tf)
			//println(templateCode)
			primType := primitiveTypes[tf]
			if primType != nil {
				returnFields = append(returnFields, queryReturnField)
				scanStatement = fmt.Sprintf("&%s.%s ", modelVariable, f.Name)
				scanStatementArgs = append(scanStatementArgs, scanStatement)
			}
		}

		// excluded fields for PUT,POST and GET
		// required: required tag
		// datatype categorization int, float64, string, boolean
		// other primitive type should be read with sql.Null* datatype
		// Any non primitive type should be handled manually in version 1
		// jsonb should be handled manually
		// scan
		// Value assignment to struct

	}

	// Parse request body
	fmt.Println(PARSE_REQUEST_BODY)
	// Required and remove field for create
	createRequiredStmt := fmt.Sprintf("createRequired := []string{\"%s\"}", strings.Join(createRequired, "\",\""))
	createRemoveStmt := fmt.Sprintf("createRemove := []string{\"%s\"}", strings.Join(createRemove, "\",\""))

	// Required and remove field for update
	updateRequiredStmt := fmt.Sprintf("updateRequired := []string{\"%s\"}", strings.Join(updateRequired, "\",\""))
	updateRemoveStmt := fmt.Sprintf("updateRemove := []string{\"%s\"}", strings.Join(updateRemove, "\",\""))

	//CREATE remove & required
	createRemoveAndRequiredStmts := fmt.Sprintf(REQUIRED_AND_REMOVE, "createRemove", "createRequired")
	fmt.Println(createRequiredStmt)
	fmt.Println(createRemoveStmt)
	fmt.Println(createRemoveAndRequiredStmts)

	//UPDATE remove & required
	updateRemoveAndRequiredStmts := fmt.Sprintf(REQUIRED_AND_REMOVE, "updateRemove", "updateRequired")
	fmt.Println(updateRequiredStmt)
	fmt.Println(updateRemoveStmt)
	fmt.Println(updateRemoveAndRequiredStmts)

	// Datatype statements
	stringFieldStmt := fmt.Sprintf("stringField := []string{\"%s\"}", strings.Join(stringField, "\",\""))
	fmt.Println(stringFieldStmt)

	intFieldStmt := fmt.Sprintf("intField := []string{\"%s\"}", strings.Join(intField, "\",\""))
	fmt.Println(intFieldStmt)

	floatFieldStmt := fmt.Sprintf("floatField := []string{\"%s\"}", strings.Join(floatField, "\",\""))
	fmt.Println(floatFieldStmt)

	boolFieldStmt := fmt.Sprintf("boolField := []string{\"%s\"}", strings.Join(boolField, "\",\""))
	fmt.Println(boolFieldStmt)

	jsonFieldStmt := fmt.Sprintf("jsonField := []string{\"%s\"}", strings.Join(jsonField, "\",\""))
	fmt.Println(jsonFieldStmt)
	//
	model := fmt.Sprintf("%s{}", t)
	modelSplit := strings.Split(model, ".")
	if len(modelSplit) > 1 {
		if modelSplit[0] == "main" {
			model = t.Name()
		}
	}

	invalidDataTypeCheck := fmt.Sprintf(INVALID_DATA_TYPE, model)
	fmt.Println(invalidDataTypeCheck)
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
	PARSE_REQUEST_BODY = `
body, err := parseutil.ParseJSON(context)
if err != nil {
	log.Printf("\nError: %+v", err)
}
	`

	REQUIRED_AND_REMOVE = `
body = parseutil.RemoveFields(body, %s)
missing := parseutil.EnsureRequired(body, %s)
if len(missing) != 0 {
	log.Println("missing", missing)
	return context.JSON(http.StatusBadRequest, missing)
}
	`

	INVALID_DATA_TYPE = `
body, invalidType := parseutil.MapX(body, %s, stringFields, floatField, intField, boolField, jsonField)
if len(invalidType) != 0 {
	log.Println("invalidType", invalidType)
	return context.JSON(http.StatusBadRequest, invalidType)
}
`
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
