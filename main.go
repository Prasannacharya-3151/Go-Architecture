package main

import (  //import used for a external and internal packages
	"log"  //built-in go package for a logging like printing logs and hanling a errors
	"net/http" //http server packages used for the creating a server and handling a request from the client 
	"os" //os package used for a readung environment variables and interacting with system

	"github.com/joho/godotenv" //this one is the external library used for the loading .env file into the enviroment variable without this one .env values wont be accessibel


)

func main() { //in here no parameters assigned beacuse the main function is the entry point of the programm and it will be called by the go runtime when the program starts.
	if err := godotenv.load(); err != nil {  //; sepearater used to separate the variable declaration ad the if statement and the err variable is used to store the error if there is any error while loading the env file 
		log.Fatalln("env file is not loaded") // Fatal is the stops the program and print the error message to the log
	} //try to load the env file if there is an error it will log the error and exist the program

	serverAddress := os.Getenv("PORT") //:= it means deaclire + assign the value to the variable serverAddress and os.Getenv is used to get the value os the environment variable with the key "PORT"

	dbConn, err := connect()
	if err != nil {
		log.Fatalln("failed to connect to db:", err)
	}
	log.Println("DB connected")

	repo := repository.NewPostgressRepository(dbConn) //repo is an varible and NewPostgrsRepository is a constructor fucntion and then dbConn passed as dependency
   //we have created a repository instanceby calling the NewPostgressRepository function and passing the dbConn as an argument

   if err := repo.Init(); err != nil {
	log.Fatalln(err)
   }

   router := routes.MountRouters(repo) //router http hanlding and MountRoutes funvtion and repo passed inside
   //we have created a router instance by calling the MountRouters function and passing the repo as an argument

   log.Print(serverAddress) //it will print the server address to the log

   http.ListenAndServe(serverAddress, router)
   //http is the package and ListenAndserve is a function that starts an Http server and it takes two arguments the first one is the address to Listen and the second one is the handler to handle the incoming request in this case we are passing the router as a handler 
}