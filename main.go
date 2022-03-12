package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Here, we define the templates as lobal viriables to be reachable within all functions
var resultTmpl *template.Template
var rootTmpl *template.Template
var notfoundTmpl *template.Template
var startAdress string = "/home/soorena/"
var port int = 8080

//Some structs to deal with data used in program
type resultPage struct {
	Title string
	Items []diagnosis.disease
}

//This handler handles the main page
func rootHandler(wr http.ResponseWriter, req *http.Request) {
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
func formHandler(wr http.ResponseWriter, req *http.Request) {
	//This passes the post request to the formProcess function and gets the patient struct
	pat, err := formProcess(req)
	if err != nil {
		log.Fatal(err)
	}
	//Print out the passed symptoms
	for i := 0; i < len(pat.symptoms); i++ {
		fmt.Println(pat.symptoms[i])
	}
	//Prepare the data to feed into the template
	//In this case it contains the title of the page and matched disease
	data := resultPage{
		Title: "A suitable title",
		Items: []disease{
			{
				Name: "Cancer",
				Symps: []string{
					pat.symptoms[0],
					"symp1",
					"symp2",
					"symp3",
				},
			},

			{
				Name: "Aids",
				Symps: []string{
					"symp3",
					"symp4",
					"symp5",
				},
			},
		},
	}
	//Set the header's cookies
	wr.Header().Set("Date", "Mon, 01 Jan 2020 00:00:00 GMT")
	//Feed the data into the result page template and serve it
	err = resultTmpl.Execute(wr, data)
	if err != nil {
		log.Fatal(err)
	}
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
	//Feed the data into the notfound page template and serve it
	err := notfoundTmpl.Execute(wr, data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//Read the templates from the respective html files
	resultTmpl = template.Must(template.ParseFiles(startAdress + "templates/result.html"))
	rootTmpl = template.Must(template.ParseFiles(startAdress + "templates/index.html"))
	notfoundTmpl = template.Must(template.ParseFiles(startAdress + "templates/notfound.html"))
	//Initialize the mux router
	router := mux.NewRouter().StrictSlash(true)
	//Set the respective handlers to uri addresses
	router.HandleFunc("/", http.RedirectHandler("/index.html", 301).ServeHTTP)
	router.HandleFunc("/index.html", rootHandler)
	router.HandleFunc("/result.html", formHandler).Methods("Post")
	//Set notfound handler function to the wildcard
	router.HandleFunc("/{*}", notfoundHandler)
	//Listen to the defined port and serve
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
