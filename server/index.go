package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Index doctor
// On create & update
// Index hospitals/clinic
// On create & update
type IndexService interface {
	IndexDoctor(map[string]interface{}) error
}

type IndexServiceImpl struct {
	Client *http.Client
}

func (indexService *IndexServiceImpl) IndexDoctor(dr map[string]interface{}) error {
	// Get id

	// Get related fields
	// Prepare json body
	// Index
	indexDrUrl := "https://aw8akjyvyt:tyqzkp81x7@tarkol-8348260269.eu-central-1.bonsaisearch.net:443/doctors/_doc/" + dr["accountId"].(string)
	//indexService.Client.Post()

	reqBody, err := json.Marshal(dr)
	if err != nil {
		log.Printf("Error marshalling payload to json %v", err)
		return errors.New("Error marshaling payload")
	}

	endpoint := indexDrUrl
	request, _ := http.NewRequest("POST", endpoint, bytes.NewReader(reqBody))
	request.Header.Set("Content-Type", "application/json")
	response, err := indexService.Client.Do(request)
	if err != nil {
		errStr := fmt.Sprintf("Error connecting to server %v", err)
		log.Print(errStr)
		return errors.New(errStr)
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	fmt.Println("\n\nINDEXED DR. \n\n")
	fmt.Printf("\n\nErr= %v", err)
	return nil

}

// Search indexed doctor by disease & city
// Search clinic by name & city
