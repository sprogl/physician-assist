package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/sprogl/website/diagnosis"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

//Here, we define the templates as global viriables to be reachable within all functions
var rootTmpl *template.Template
var notfoundTmpl *template.Template
var dbconn *pgx.Conn

//Some structs to deal with data used in program
type resultPage struct {
	Items []diagnosis.Disease `json:"diseases"`
}

//This handler handles the main page
func dignosisMainHandler(wr http.ResponseWriter, req *http.Request) {
	//Prepare the data to feed into the template
	//In this case it just contains the title of the page
	data := struct {
		Title string
	}{
		Title: "A suitable title",
	}
	//Set header, i.e., the date
	wr.Header().Set("Date", "Mon, 01 Jan 2020 00:00:00 GMT")
	//Feed the data into the root page template and serve it
	err := rootTmpl.Execute(wr, data)
	if err != nil {
		log.Fatal(err)
	}
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

//This function handles requests to undefined pages
func notfoundHandler(wr http.ResponseWriter, req *http.Request) {
	//Extract the requested uri
	url := req.URL.String()
	//Construct the data to be fed into the template
	//It, particularly, contains the page title and requested uri
	data := struct {
		Title string
		URL   string
	}{
		Title: "A suitable title",
		URL:   url,
	}
	//Set the header's cookies
	wr.Header().Set("Date", "Mon, 01 Jan 2020 00:00:00 GMT")
	wr.WriteHeader(http.StatusBadRequest)
	//Feed the data into the notfound page template and serve it
	err := notfoundTmpl.Execute(wr, data)
	if err != nil {
		log.Fatal(err)
	}
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

	/*
		//Get the executable's address to find the resources relatively
		templatesAdress, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		templatesAdress = filepath.Dir(templatesAdress) + "/templates/"
		//Read the templates from the respective html files
		rootTmpl = template.Must(template.ParseFiles(templatesAdress + "index.html"))
		notfoundTmpl = template.Must(template.ParseFiles(templatesAdress + "notfound.html"))

		//Set the respective handlers to uri addresses
		router.HandleFunc("/", http.RedirectHandler("/diagnosis/v1/index.html", 301).ServeHTTP)
		router.HandleFunc("/diagnosis/v1/index.html", dignosisMainHandler).Methods("Get")
		//Set notfound handler function to the wildcard
		router.HandleFunc("/{*}", notfoundHandler)
	*/

	//Initialize the mux router
	router := mux.NewRouter().StrictSlash(true)
	//Set the respective handlers to uri addresses
	router.HandleFunc("/diagnosis/v1/index.html", dignosisFormHandler).Methods("Post")
	//Listen to the defined port and serve
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("DIAGAPIPORT")), router))
}
