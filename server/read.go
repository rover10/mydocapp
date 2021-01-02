package server

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/rover10/model"
	"github.com/rover10/response"
	"github.com/rover10/token"
)

//Doctor read a doctor
func Doctor(context echo.Context) error {
	return nil
}

//Clinic read a clinic detail
func (s *Server) Clinic(context echo.Context) error {
	log.Println("Clinic")
	clinics := []model.Clinic{}
	if err := s.DB.DBORM.Table("clinic").Find(&clinics).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusOK, clinics)
}

//Disease reads a disease detail
func (s *Server) Disease(context echo.Context) error {
	log.Println("Disease")
	clinics := []model.Disease{}
	if err := s.DB.DBORM.Table("disease").Find(&clinics).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusOK, clinics)
}

//UserType reads user type
func (s *Server) UserType(context echo.Context) error {
	userType := []model.UserType{}
	if err := s.DB.DBORM.Table("user_type").Find(&userType).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusOK, userType)
}

//Country reads countries
func (s *Server) Country(context echo.Context) error {
	countries := []model.Country{}
	if err := s.DB.DBORM.Table("country").Find(&countries).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusOK, countries)
}

//Clinic read a clinic detail
func (s *Server) GenerateQr(context echo.Context) error {
	//login := token.GetLoggedIn(context)
	//accountID := login["uid"].(string)
	log.Println("GenerateQr")
	if appointmentID := context.QueryParam("appointmentId"); appointmentID != "" {
		fmt.Println(appointmentID)
		qrCode, _ := GenerateQrCode(appointmentID, 128, 128)
		return png.Encode(context.Response().Writer, qrCode)
	}
	return context.JSON(http.StatusNoContent, "")
	//return context.JSON(http.StatusOK, clinics)
}

func GenerateQrCode(data string, l int, w int) (barcode.Barcode, error) {
	qrCode, err := qr.Encode(data, qr.L, qr.Auto)
	if err != nil {
		return nil, err
	}
	return barcode.Scale(qrCode, l, w)
}

//Appointment read appointment
func (s *Server) Appointment(context echo.Context) error {
	log.Println("Appointment")
	login := token.GetLoggedIn(context)
	accountID := login["uid"].(string)
	fmt.Println(accountID)
	appointments := []model.Appointment{}
	if err := s.DB.DBORM.Table("appointment").Where("account_id = ?", accountID).Find(&appointments).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusOK, appointments)
	//s.DB.RetrieveAppointment(sid)

}

//Appointment read appointment
func (s *Server) AppointmentV2(context echo.Context) error {
	log.Println("AppointmentV2")
	login := token.GetLoggedIn(context)
	accountID := login["uid"].(string)
	fmt.Println(accountID)
	appointments := []response.Appointment{}

	cancelledS := context.QueryParam("isCancelled")
	historyS := context.QueryParam("history")
	isCancelled := false
	history := false

	if cancelledS != "" {
		if strings.ToLower(cancelledS) == "true" {
			isCancelled = true
		}
	}

	if historyS != "" {
		if strings.ToLower(historyS) == "true" {
			history = true
		}
	}
	//.Preload("Clinic").Preload("Patient").Preload("Doctor")
	//clinic := response.Clinic{}
	ret := s.DB.DBORM.Debug().
		//Preload("Clinic").
		Table("appointment").
		Preload("Clinic", func(db *gorm.DB) *gorm.DB {
			return db.
				Table("clinic")
		}).
		Preload("Patient", func(db *gorm.DB) *gorm.DB {
			return db.Table("patient")
		}).
		Preload("Doctor", func(db *gorm.DB) *gorm.DB {
			return db.Table("doctor")
		})

	//	Where("account_id = ?", accountID)
	if isCancelled {
		ret = ret.Where("account_id = ? AND cancelled = 'True'", accountID)
	} else if history {
		ret = ret.Where("account_id = ? AND slot_date_time < now()", accountID)
	} else {
		ret = ret.Where("account_id = ? AND slot_date_time > now() and cancelled = 'False'", accountID)
	}

	if err := ret.Find(&appointments).Error; err != nil {
		return err
	}

	//s.DB.DBORM.
	return context.JSON(http.StatusOK, appointments)
	//s.DB.RetrieveAppointment(sid)

}

//Patient read linked patient
func (s *Server) Patient(context echo.Context) error {

	log.Println("Patient")
	login := token.GetLoggedIn(context)
	accountID := login["uid"].(string)
	patients := []model.Patient{}
	if err := s.DB.DBORM.Table("patient").Where("account_id = ?", accountID).Find(&patients).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusOK, patients)
}

//Treatment read treatment detail which includes patient_problem_description, prescription, test
func Treatment(context echo.Context) error {
	return nil
}

//
func (s *Server) Search(context echo.Context) error {
	//search()
	// open a new index
	index := s.Index

	data := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}{
		Name:        "your_data you This us your data",
		Description: "I have something to search",
	}

	//index.
	// index some data
	err := index.Index("id1", data)

	data.Name = "Do you have something to search is something?"
	data.Description = "Do you have something to search is something?"
	err = index.Index("id2", data)
	data.Name = "Do you need something"
	data.Description = "What is something you want to search?"
	err = index.Index("id3", data)

	// search for some text
	query := bleve.NewFuzzyQuery("you") //bleve.NewMatchQuery("you")
	query.SetFuzziness(2)
	//query.SetField("name")

	//query := bleve.NewWildcardQuery("something")
	// sr := make([]string, 0)
	// sr = append(sr, "have")
	// sr = append(sr, "something")

	//query2 := bleve.NewPhraseQuery(sr, "Name")
	fmt.Println(query)
	//query.SetFuzziness(3)
	//bleve.NewPhraseQuery("you")
	//index.
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	//index.Se
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}
	fmt.Println("--------> --> --> ")
	fmt.Println(searchResults)
	context.JSON(http.StatusOK, searchResults)

	return nil
}

func search() {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	}{
		Name: "text",
	}

	// index some data
	index.Index("id", data)

	// search for some text
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}
