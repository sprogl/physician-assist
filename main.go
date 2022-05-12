package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sprogl/website/diagnosis"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

//Here, we define the templates as global viriables to be reachable within all functions
var dbconn *pgx.Conn

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
		fmt.Println("Err: line 53 of main.go")
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}
	//Print out the passed symptoms
	symps := pat.Symptoms
	for i := 0; i < len(symps); i++ {
		fmt.Println(symps[i])
	}
	//Analyse the symtoms and get the list of matched diseases
	diseases, err := pat.Diagnose(dbconn)
	if err != nil {
		fmt.Println("Err: line 65 of main.go")
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}
	//Prepare the data to feed into the template
	//In this case it contains a list of matched disease and
	//Marshal the input data
	dataJson, err := json.Marshal(resultPage{
		Items: diseases,
	})
	if err != nil {
		fmt.Println("Err: line 76 of main.go")
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}
	//Declare that the response data will be in json format
	wr.Header().Add("Access-Control-Allow-Origin", "*")
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
		log.Fatal(err)
	}
	defer dbconn.Close(context.Background())

	//Initialize the mux router
	router := mux.NewRouter().StrictSlash(true)
	//Set the respective handlers to uri addresses
	router.HandleFunc("/diagnosis/v1/index.html", dignosisFormHandler).Methods("Post")
	//Listen to the defined port and serve
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("DIAGAPIPORT")), router))
}
