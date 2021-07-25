package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/rover10/api"
	"github.com/rover10/config"
	"github.com/rover10/database"
	"github.com/rover10/server"
)

func main() {
	dborm, err := DBConnect()
	defer dborm.Close()
	//defer dbGo.Close()
	if err != nil {
		log.Errorf("failed to connect database. Error :%+v", err)
	}
	fmt.Println(dborm)
	config := config.Config{}
	config.APIPath = "/"
	config.DBHost = "docapp11"
	config.DBName = "postgres"
	config.DBPassword = "root"
	config.DBUser = "postgres"
	config.Host = "localhost"
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		config.Port = 5000
	} else {
		config.Port = port
	}
	log.Info("Port connected")
	log.Info(config.Port)

	config.WebDir = "web"
	server := server.NewServer(config)
	docdb := database.DocDB{}
	docdb.DBORM = dborm
	//docdb.DB = dbGo
	server.DB = &docdb
	api.Api(server)
	err = server.Start()
	if err != nil {
		log.Errorf("server failed to start. Error: %+v", err)
	}

}

func DBConnect() (*gorm.DB, error) {
	// postgres://blxiorkgqxoqts:c44c8a5f2b73f9e303446af8e2e0d53ca9c11a71686f27f413193828da72bd5b@ec2-107-21-10-179.compute-1.amazonaws.com:5432/d2hjt44nklvpml
	dbinfo := "postgres://blxiorkgqxoqts:c44c8a5f2b73f9e303446af8e2e0d53ca9c11a71686f27f413193828da72bd5b@ec2-107-21-10-179.compute-1.amazonaws.com:5432/d2hjt44nklvpml"
	// dbinfo := "user=postgres port=5432 password=root dbname=docapp host=localhost sslmode=disable"
	db, err := gorm.Open("postgres", dbinfo)
	//defer db.Close()
	if err != nil {
		log.Errorf("failed loading parameteres. Error :%+v", err)
		return nil, err
	}

	// db2, err := sql.Open("postgres", dbinfo)
	// //defer db2.Close()
	// err = db2.Ping()
	// if err != nil {
	// 	log.Errorf("failed to ping postgres Error :%+v", err)
	// 	return nil, nil, err
	// }

	log.Info("Successfully connected")

	return db, nil
}

func execute(tx *sql.Tx, sqlStr string, vals []interface{}) error {
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		log.Error("Err: stmt %+v", err)
		log.Error(err)
		return err
	}
	defer stmt.Close()
	//format all vals at once
	_, err = stmt.Exec(vals...)
	if err != nil {
		log.Error("Err: res %+v", err)
		return err
	}

	return nil
}

func (p *PathManager) Rollback() {
	p.tx.Rollback()
}

type PathManager struct {
	tx *sql.Tx
}
