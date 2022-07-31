package main

import (
	"context"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
	mux "github.com/gorilla/mux"
	pgx "github.com/jackc/pgx/v4"
)

//Following environment variables are needed for this programm to run
//1- AUTHAPIPORT	the port on which the programm responds
//2- DBUSER			the databse username
//3- DBPASS			the database password
//4- DBIP			the databse IP
//5- DBPORT			the databse port
//6- DATABASE		the database name
//7- PRIVKEY		the private key for authentication

//Here, we define the templates as global viriables to be reachable within all functions
var dbconn *pgx.Conn
var pubKeyStr string
var privKey ed25519.PrivateKey

//Introduce the struct LogIn and SignUp
type LogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUp struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//This handler serves the authentication public key
//This key is used to verify the JW token inside other components of the website
// TODO: Add captcha
func pubKeyHandler(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(wr, pubKeyStr)
}

//This function handles the request to the disease form
//and returns the sugests the fitting disease
func logInHandler(wr http.ResponseWriter, req *http.Request) {
	const authQuery = `
SELECT count(id), id AS userID
FROM user_table
WHERE email=$1 AND password_hash=$2;
`
	var claims = jwt.StandardClaims{
		Issuer: "Docassist aut",
	}
	var l LogIn
	//Specify the output format
	wr.Header().Add("Access-Control-Allow-Origin", "*")
	wr.Header().Set("Content-Type", "text/plain")
	//Parse the posted form and extract it for further process
	jsonDecoder := json.NewDecoder(req.Body)
	err := jsonDecoder.Decode(&l)
	if err != nil {
		http.Error(wr, "wrong request format", http.StatusBadRequest)
		return
	}
	_, err = mail.ParseAddress(l.Email)
	if err != nil {
		http.Error(wr, "wrong email format", http.StatusBadRequest)
		return
	}
	// TODO: add salt to hash
	passHash := sha256.Sum256([]byte(l.Password))
	passHashBase64 := base64.StdEncoding.EncodeToString(passHash[:])
	rows, err := dbconn.Query(context.Background(), authQuery, l.Email, passHashBase64)
	if err != nil {
		http.Error(wr, "Something went wrong!", http.StatusInternalServerError)
		log.Fatal(err)
	}
	defer rows.Close()
	var n int
	var sub string
	err = rows.Scan(&n, &sub)
	if err != nil {
		http.Error(wr, "something went wrong.", http.StatusInternalServerError)
		log.Fatal(err)
	}
	if n == 0 {
		http.Error(wr, "authentication failed.", http.StatusUnauthorized)
		return
	}
	if n != 1 {
		http.Error(wr, "something went wrong.", http.StatusInternalServerError)
		log.Fatal(err)
	}
	claims.ExpiresAt = time.Now().Unix() + 432000
	claims.Subject = sub
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	signedToken, err := token.SignedString(privKey)
	if err != nil {
		http.Error(wr, "something went wrong.", http.StatusInternalServerError)
		log.Fatal(err)
	}
	//Feed the data into the result page template and serve it
	wr.WriteHeader(http.StatusOK)
	fmt.Fprint(wr, signedToken)
}

func registerationHandler(wr http.ResponseWriter, req *http.Request) {
	const newidQuery = `
	SELECT count(id)
	FROM user_table;
	`
	const signUpQuery = `
	INSERT INTO user_table
		VALUES (
			$1,
			$2,
			$3,
			$4
		);
	`
	var duplicateIDErr = errors.New("duplicate key value violates unique constraint \"user_table_pkey\"")
	var duplicateEmailErr = errors.New("duplicate key value violates unique constraint \"user_table_email_key\"")
	var s SignUp
	// Specify the output format
	wr.Header().Add("Access-Control-Allow-Origin", "*")
	wr.Header().Set("Content-Type", "text/plain")
	//Parse the posted form and extract it for further process
	jsonDecoder := json.NewDecoder(req.Body)
	err := jsonDecoder.Decode(&s)
	if err != nil {
		http.Error(wr, "wrong request format", http.StatusBadRequest)
		return
	}
	_, err = mail.ParseAddress(s.Email)
	if err != nil {
		http.Error(wr, "wrong email format", http.StatusBadRequest)
		return
	}
	// TODO: add salt to hash
	passHash := sha256.Sum256([]byte(s.Password))
	passHashBase64 := base64.StdEncoding.EncodeToString(passHash[:])
	rows, err := dbconn.Query(context.Background(), newidQuery)
	if err != nil {
		http.Error(wr, "something went wrong.", http.StatusInternalServerError)
		log.Fatal(err)
	}
	var id int
	err = rows.Scan(&id)
	if err != nil {
		http.Error(wr, "something went wrong.", http.StatusInternalServerError)
		log.Fatal(err)
	}
	for err := duplicateIDErr; err == duplicateIDErr; id++ {
		rows, err = dbconn.Query(context.Background(), signUpQuery, id, s.Name, s.Email, passHashBase64)
		switch err {
		case duplicateIDErr:
			continue
		case duplicateEmailErr:
			http.Error(wr, "email already exists.", http.StatusBadRequest)
			return
		default:
			http.Error(wr, "something went wrong.", http.StatusInternalServerError)
			log.Fatal(err)
		}
	}
	defer rows.Close()
	wr.WriteHeader(http.StatusOK)
	fmt.Fprint(wr, "user successfully signed up.")
}

func main() {
	var err error
	//Get the databse address from the environment variables
	DBAdrress := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBIP"), os.Getenv("DBPORT"), os.Getenv("DATABASE"))
	//Initiate the database connection
	dbconn, err = pgx.Connect(context.Background(), DBAdrress)
	if err != nil {
		log.Fatal(err)
	}
	defer dbconn.Close(context.Background())
	//Read the private key from the environment variables
	//and generate the public key from it
	privKeyBytes, err := base64.StdEncoding.DecodeString(os.Getenv("PRIVKEY"))
	if err != nil {
		log.Fatal(err)
	}
	if len(privKeyBytes) != ed25519.PrivateKeySize {
		log.Fatal(errors.New("inapropirate private key length"))
	}
	privKey = ed25519.PrivateKey(privKeyBytes)
	pubKey := privKey.Public()
	pubKeyStr = base64.StdEncoding.EncodeToString(pubKey.(ed25519.PublicKey))
	//Initialize the mux router
	router := mux.NewRouter().StrictSlash(true)
	//Set the respective handlers to uri addresses
	router.HandleFunc("/authentication/v1/login", logInHandler).Methods(http.MethodPost)
	router.HandleFunc("/authentication/v1/publickey", pubKeyHandler).Methods(http.MethodGet)
	router.HandleFunc("/authentication/v1/register", registerationHandler).Methods(http.MethodGet)
	//Listen to the defined port and serve
	err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("AUTHAPIPORT")), router)
	if err != nil {
		log.Fatal(err)
	}
}
