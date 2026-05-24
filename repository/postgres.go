package repository

import (
	"database/sql" //manages the db connection like Exec,Query,QuryRow
	"fmt"
	"github.com/Prasannacharya-3151/Go-Bank-Service/models"
)

type PostgresRepository struct {
	db *sql.DB //db connection stored means this struct holdes the db connection 
}

func NewPostgresrepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db : db}
}

//r is the recievier of the method and it is a pointer to the PostgesRepository struct and attached to the struct and it will be used to access the db connection and then it will return an error if there is any error while creating the table in the database
func (r *PostgresRepository) Init() error{ //init is a method of the PostgresRepository struct and it will return an error if there is any error while creating the table in the database 
	query := `CREATE TABLE IF NOT EXISTS customers (
	id UUID PRIMERY KEY,
	name VARCHAR(255),
	email VARCHER(255),
	phone VARCHER(20),
	address TEXT
	);`
	_, err := r.db.Exec(query)
	return err 
}

func (r *PostgresRepository) CreateCustomer(customer *models.Customer) error{
	query := 
}