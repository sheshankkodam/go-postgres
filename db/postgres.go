package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbUser      = "postgres"
	dbName      = "mytestdb"
	insertQuery = "INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;"
)

//go:generate mockgen --destination=mocker/mock.go -package=mocker github.com/sheshankkodam/MyPostgres/db PostgresOperations
type PostgresOperations interface {
	Insert(username, departmentNam, createdDate string) error
}

type PostgresService struct {
	Ops PostgresOperations
}

func (p *PostgresService) Insert(user, departmentName, createdDate string) error {
	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbUser, dbName)
	dbClient, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return err
	}
	log.Println("Database initiated successfully")
	var lastInsertId int
	err = dbClient.QueryRow(insertQuery, user, departmentName, createdDate).Scan(&lastInsertId)
	if err != nil {
		return err
	}
	log.Println("last inserted id =", lastInsertId)
	return nil
}
