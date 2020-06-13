package main

import (
	"database/sql"

	"github.com/labstack/gommon/log"
)

func main() {

}

func DBConnect() (*sql.DB, error) {
	dbinfo := "user=postgres port=5432 password=root dbname=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Errorf("failed loading parameteres. Error :%+v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Errorf("failed to ping postgres Error :%+v", err)
		return nil, err
	}

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
