package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sprogl/website/diag_api/diagnosis"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

//Here, we define the templates as global viriables to be reachable within all functions
var dbconn *pgx.Conn

//Following environment variables are needed for this programm to run
//1- DIAGAPIPORT	the port on which the programm responds
//2- DBUSER			the databse username
//3- DBPASS			the database password
//4- DBIP			the databse IP
//5- DBPORT			the databse port
//6- DATABASE		the database name

//Some structs to deal with data used in program
type resultPage struct {
	Items []diagnosis.Disease `json:"diseases"`
}

//This function handles the request to the disease form
//and returns the sugests the fitting disease
func dignosisFormHandler(wr http.ResponseWriter, req *http.Request) {
	//This passes the post request to the formProcess function and gets the patient struct
	pat, err := diagnosis.FormProcess(req)
	if err != nil {
		fmt.Println("Err: line 39 of main.go")
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}
	//Analyse the symtoms and get the list of matched diseases
	diseases, err := pat.Diagnose(dbconn)
	if err != nil {
		fmt.Println("Err: line 46 of main.go")
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}
	//Prepare the data to feed into the template
	//In this case it contains a list of matched disease and
	//Marshal the input data
	dataJson, err := json.Marshal(resultPage{
		Items: diseases,
	})
	if err != nil {
		fmt.Println("Err: line 57 of main.go")
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}
	//Declare that the response data will be in json format
	// wr.Header().Add("Access-Control-Allow-Origin", "*")
	wr.Header().Set("Content-Type", "application/json")
	//Feed the data into the result page template and serve it
	fmt.Fprintf(wr, string(dataJson))
}

func main() {
	//Get the databse address from the environment variables
	DBAdrress := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBIP"), os.Getenv("DBPORT"), os.Getenv("DATABASE"))
	//Initiate the database connection
	dbconn, err := pgx.Connect(context.Background(), DBAdrress)
	if err != nil {
		fmt.Println("Err: line 74 of main.go")
		log.Fatal(err)
	}
	defer dbconn.Close(context.Background())
	//Initialize the mux router
	router := mux.NewRouter().StrictSlash(true)
	//Set the respective handlers to uri addresses
	router.HandleFunc("/diagnosis/v1/index.html", dignosisFormHandler).Methods("Post")
	//Listen to the defined port and serve
	err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("DIAGAPIPORT")), router)
	if err != nil {
		fmt.Println("err: line 85 of main.go")
		log.Fatal(err)
	}
}
