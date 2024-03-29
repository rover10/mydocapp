package main

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/rover10/mydocapp.git/api"
	"github.com/rover10/mydocapp.git/config"
	"github.com/rover10/mydocapp.git/database"
	"github.com/rover10/mydocapp.git/server"
)

func main() {
	dborm, dbGo, err := DBConnect()
	defer dborm.Close()
	defer dbGo.Close()
	if err != nil {
		log.Errorf("failed to connect database. Error :%+v", err)
	}
	fmt.Println(dborm)
	config := config.Config{}
	config.APIPath = "/"
	config.DBHost = "localhost"
	config.DBName = "docapp"
	config.DBPassword = "root"
	config.DBUser = "postgres"
	config.Host = "localhost"
	config.Port = 3000
	config.WebDir = "web"
	server := server.NewServer(config)
	docdb := database.DocDB{}
	docdb.DBORM = dborm
	docdb.DB = dbGo
	server.DB = &docdb
	api.Api(server)
	err = server.Start()
	if err != nil {
		log.Errorf("server failed to start. Error: %+v", err)
	}

}

func DBConnect() (*gorm.DB, *sql.DB, error) {
	dbinfo := "user=postgres port=5432 password=root dbname=docapp host=localhost sslmode=disable"
	db, err := gorm.Open("postgres", dbinfo)
	//defer db.Close()
	if err != nil {
		log.Errorf("failed loading parameteres. Error :%+v", err)
		return nil, nil, err
	}

	db2, err := sql.Open("postgres", dbinfo)
	//defer db2.Close()
	err = db2.Ping()
	if err != nil {
		log.Errorf("failed to ping postgres Error :%+v", err)
		return nil, nil, err
	}

	log.Info("Successfully connected")

	return db, db2, nil
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
