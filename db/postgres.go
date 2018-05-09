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
	dbConn *sql.DB
	Ops PostgresOperations
}

func NewPostgresService() (*PostgresService, error) {
	dbInfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbUser, dbName)
	dbClient, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}
	log.Println("Database initiated successfully")
	p := &PostgresService{dbConn:dbClient}
	return p, nil
}

func (p *PostgresService) Insert(user, departmentName, createdDate string) error {
	var lastInsertId int
	if err := p.dbConn.QueryRow(insertQuery, user, departmentName, createdDate).Scan(&lastInsertId); err != nil {
		return err
	}
	log.Println("Inserted successfully")
	return nil
}
