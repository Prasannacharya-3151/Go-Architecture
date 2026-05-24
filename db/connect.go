package db //package is the keyword used to define a package in go an here we are defining a package named db

import ( //in here importing a required packages for the database connection
	"database/sql" //sql package is used for working with databases in go and it provides a generic interface for interacting with differnt databases
	"log"
"os"
_ "github.com/jackc/pgx/v5/stdlib" //thisis the external packages used for the postgres database driver and the _ is used to import the packages for its side effects without actually using any of its exported identifiers in the code)

)

func Connect() (*sql.DB, error){ //function , function name , no parameters passed , adn then retun types
	url := os.Getenv("POSTGRESS_URL") //getting a env variable with the key "POSTGRESS_URL" and storing it in the url variable

	if url == "" { //check the empty string if the url is empty it means the env variable is not loaded or not set if env not set crash the programm immediatly
		log.Fatalln("postgrss url is not loaded")
	}
	db, err :=sql.Open("pgx",url) //db an object and that error , sql.open opens a db "pgx" is the driver name and connection string url

	if err != nil {
		log.Fatalln("failed to connect to db:", err)
		return nil, err
	}

	err = db.Ping()

	if err != nil{
		return nil, err
	}
	return db, nil //db is the connection object and nill is the no error means database connection is successfull 
}