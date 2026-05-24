package models

type Customer struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
}

type CustomerAccount struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
}

type CustomerResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Emial string `json:"email"`
}


//this this one interface is very imp give a flexibilty we can swith the DB into the postgrss to mongoDB and also easy to testing and controllers not depend on the implimentation of the database and we can easily mock the database for testing purpose and also it will help us to achieve the separation of concerns and also it will help us to achieve the single responsibility principle because the repository will be responsible for handling all the database operations and the controller will be responsible for handling all the http request and response
type BankRepo interface {
	CreateCustomer(customer *Customer) error //function name and customer *Customer this is the pointer to the customer struct and then error is the return typeod of the function it will return an error 
	GetCustomer(id string) (Customer, error)
	updateCustomer(customer *Customer) error 
	DeleteCustomer(id string) error
	ListOfCustomers() ([]*Customer, error) //[] slice (array-like) and *Customer is the  pointer means return the list of the customers and error is the retunr type of the function it will return an error if there is any error while fetching the list of the custmores

}

