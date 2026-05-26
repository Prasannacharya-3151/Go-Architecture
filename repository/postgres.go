package repository

import (
	"database/sql" //manages the db connection like Exec,Query,QuryRow
	"fmt"

	"github.com/Prasannacharya-3151/Go-Bank-Service/db/models"
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
	query := `INSERT INTO customers (id,name,email,phone,address) VALUES  ($1,$2,$3,$4,$5)`
	_,err := r.db.Exec(query,customer.ID,customer.Name,customer.Email,customer.Phone,customer.Address)
	return err
}

func (r *PostgresRepository) GetCustomer(id string) (*models.Customer,error){
	query:= `SELECT id,name,email,phone,address FROM customers WHERE id = $1`
	var customer models.Customer
	err := r.db.QueryRow(query,id).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.Address,
	)

	if err!=nil{
		return nil,err
	}

	func (r *PostgresRepository) UpdateCustomer(customer *models.Customer) error {
		query := `UPDATE customers SET name=$1, email=$2, phone=$3, address=$4, WHERE id = $5`
		_,err := r.db.Exec(query,customer.Name,customer.Email,customer.Phone,customer.Address,customer.ID)
		return err
	}

	func (r *PostgresRepository) DeleteCustomer(id string) error{
		query:= `DELETE FROM customers WHERE id = $1`
		result,err:= r.db.Exec(query, id)
		if err!=nill {
			return fmt.Errorf(
				rowsAffected,err := result.RowsAffected()

				if rowsAffected == 0 {
					return fmt.Errorf("task with ID %s not found", id)
				}
				return err
		}

		func (r *PostgresRepository) ListCustomers()([]*models.Customer,error){
			query := `SELECT * FROM customers`
			rows,err := r.db.Query(query)
			if err!=nil{
				return nil, err
			}
			var customers []*models.Customer

			for rows.Next(){
				var customer models.Customererr
				if err := rows.Scan(&customer.ID,&customer.Name,&customer.Email,&customer.Phone.&customer.Address); err!=nil{
					return nil, err
				}
				customer = append(customers, &customer)

			}

			return customers, nil
		}