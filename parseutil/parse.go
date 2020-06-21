package parseutil

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"
)

//ParseJSON parses the request json
func ParseJSON(context echo.Context) (map[string]interface{}, error) {
	contentType := context.Request().Header.Get("Content-Type")
	if contentType != "application/json" {

	}

	defer context.Request().Body.Close()
	stringJSONBody, err := ioutil.ReadAll(context.Request().Body)
	if err != nil {
		return nil, err
	}

	var body map[string]interface{}
	err = json.Unmarshal([]byte(stringJSONBody), &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//RemoveFields the fields which should not be inserted or updated by user
func RemoveFields(body map[string]interface{}, removeFields []string) map[string]interface{} {
	if body != nil {
		for _, v := range removeFields {
			delete(body, v)
		}
	}
	return body
}

//EnsureRequired ensure required fields are present
func EnsureRequired(body map[string]interface{}, requiredFields []string) []string {
	missing := make([]string, 0)
	if body != nil {
		for _, v := range requiredFields {
			if body[v] == nil {
				missing = append(missing, v)
			}
		}
	}
	return missing
}

// switch k {
// //All string
// case "firstName", "lastName", "phone", "email":
// 	value, valid := ensureString(v)
// 	validBody, invalidDataType = setValue(validBody, invalidDataType, k, value, valid)
// //All integer
// case "userType", "gender", "country":
// 	value, valid := ensureString(v)
// 	validBody, invalidDataType = setValue(validBody, invalidDataType, k, value, valid)
// }
//Others

//MapX takes the model which needs to be inserted/updated
func MapX(requestBody map[string]interface{}, datamodel interface{}, kString []string, kFloat64 []string, kInt []string, kBool []string) (map[string]interface{}, []string) {
	invalidDataType := make([]string, 0)
	validBody := make(map[string]interface{}, 0)

	// string, float64, int, bool
	mString := mapKeys(kString)
	mFloat64 := mapKeys(kFloat64)
	mInt := mapKeys(kInt)
	mBool := mapKeys(kBool)

	// Convert the structure datamodel {User, Doctor, Staff..etc} to map
	inrec, _ := json.Marshal(&datamodel)
	var modelMap map[string]interface{}
	json.Unmarshal(inrec, &modelMap)

	for k := range modelMap {
		// Get value from the requestBody
		v := requestBody[k]
		if v != nil {
			if mString[k] != nil {
				value, valid := ensureString(v)
				validBody, invalidDataType = setValue(validBody, invalidDataType, k, value, valid)
			} else if mFloat64[k] != nil {
				value, valid := ensureFloat64(v)
				validBody, invalidDataType = setValue(validBody, invalidDataType, k, value, valid)
			} else if mInt[k] != nil {
				value, valid := ensureInt(v)
				validBody, invalidDataType = setValue(validBody, invalidDataType, k, value, valid)
			} else if mBool[k] != nil {
				value, valid := ensureBool(v)
				validBody, invalidDataType = setValue(validBody, invalidDataType, k, value, valid)
			}
		}
	}
	// validBody will only contains fields which are in database table so query builder can easily build query
	// If len(invaliddataType) > 0: Invalid datatype
	return validBody, invalidDataType
}

//get map
func mapKeys(keys []string) map[string]interface{} {
	kMap := map[string]interface{}{}
	if keys != nil {
		for _, v := range keys {
			kMap[v] = ""
		}
	}
	return kMap
}

func setValue(validBody map[string]interface{}, invalidDataType []string, key string, value interface{}, valid bool) (map[string]interface{}, []string) {
	if !valid {
		invalidDataType = append(invalidDataType, key)
	} else {
		validBody[key] = value
	}
	return validBody, invalidDataType
}

func ensureString(v interface{}) (string, bool) {
	value, ok := v.(string)
	if !ok {
		return "", false
	}
	return value, true
}

func ensureInt(v interface{}) (int, bool) {
	value, ok := v.(float64)
	if !ok {
		return 0, false
	}
	return int(value), true
}

func ensureFloat64(v interface{}) (float64, bool) {
	value, ok := v.(float64)
	if !ok {
		return 0, false
	}
	return value, true
}

func ensureBool(v interface{}) (bool, bool) {
	value, ok := v.(bool)
	if !ok {
		return false, false
	}
	return value, true
}
