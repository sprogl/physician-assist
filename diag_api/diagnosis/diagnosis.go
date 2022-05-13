package diagnosis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v4"
)

var Cancer = Disease{
	Name: "Cancer",
	Symptoms: []string{
		"symp1",
		"symp2",
		"symp3",
	},
}

var Aids = Disease{
	Name: "Aids",
	Symptoms: []string{
		"symp3",
		"symp4",
		"symp5",
	},
}

//Introduce the struct Disease and some method to export its content
type Disease struct {
	Name     string   `json:"name"`
	Symptoms []string `json:"symptoms"`
}

//Introduce the struct Patient and some method to export its content
type Patient struct {
	Gender   string   `json:"gen"`
	Age      int      `json:"age"`
	Symptoms []string `json:"symps"`
}

func (pat *Patient) IsFemale() bool {
	return (*pat).Gender == "Female"
}

func (pat *Patient) SympsCode(dbconn *pgx.Conn) ([]byte, error) {
	codeLength, err := SympsCodeLength(dbconn)
	if err != nil {
		fmt.Println("Err: line 139 of diagnosis.go")
		return nil, err
	}
	code := make([]byte, codeLength)
	rows, err := dbconn.Query(context.Background(), "SELECT id FROM symptoms_table WHERE symp IN $1;", []interface{}{(*pat).Symptoms})
	if err != nil {
		fmt.Println("Err: line 145 of diagnosis.go")
		return nil, err
	}
	defer rows.Close()
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("Err: line 153 of diagnosis.go")
			return nil, err
		}
		codeIndex := id / 8
		code[codeIndex] |= (1 << (id % 8))
	}
	return code, nil
}

//This function searchs the patient's data (symtoms, gender and age) in the database and
//finds the matched disease. It returns the result as list of matched diseases and
//any error occurred in the process.
func (pat *Patient) Diagnose(dbconn *pgx.Conn) ([]Disease, error) {
	//Query to the diseases database
	rows, err := dbconn.Query(context.Background(),
		"SELECT name, gen, age_min, age_max, code FROM disease_table	WHERE (gen IS NULL OR gen = $1) AND (age_min <= $2 AND age_max > $2) AND ();",
		[]interface{}{3}...) //ToDO
	if err != nil {
		fmt.Println("Err: line 94 of diagnosis.go")
		return nil, err
	}
	defer rows.Close()
	//Define the uninitialized disease (d) and disease slice (ds)
	var ds []Disease
	var d Disease
	//Loop over the rows ofdatabase response and populate the disease slice with the return data
	for rows.Next() {
		err := rows.Scan(&d.Name) //TODO
		if err != nil {
			fmt.Println("Err: line 105 of diagnosis.go")
			return nil, err
		}
		ds = append(ds, d)
	}
	//Return the result
	return []Disease{Aids, Cancer}, nil //TODO
}

//This function processes the post request
//and extracts the sanitized input inside the post request
func FormProcess(req *http.Request) (*Patient, error) {
	//Define the uninitialized patient data
	var p Patient
	//Parse the posted form and extract it for further process
	jsonDecoder := json.NewDecoder(req.Body)
	err := jsonDecoder.Decode(&p)
	if err != nil {
		fmt.Println("Err: line 57 of diagnosis.go")
		return nil, err
	}
	//Check the gender input and set it inside the patient struct
	switch p.Gender {
	case "female":
		p.Gender = "Female"
	case "male":
		p.Gender = "Male"
	default:
		fmt.Println("Err: line 67 of diagnosis.go")
		return nil, errors.New("wrong gender input format")
	}
	//Check the age input and set in inside the patient struct
	if p.Age < 0 || p.Age > 150 {
		fmt.Println("Err: line 72 of diagnosis.go")
		return nil, errors.New("wrong age input format")
	}
	//Check the symptoms input and set in inside the patient struct
	//This is done through splitting the entry by commas
	// seperator := regexp.MustCompile(" *(([,;](\r\n|\n)* *)|([,;]*(\r\n|\n) *))")
	// p.Symptoms = seperator.Split(form.Symptoms, -1)
	if len(p.Symptoms) == 0 {
		fmt.Println("Err: line 80 of diagnosis.go")
		return nil, errors.New("empty list of symptoms")
	}
	//Return the resulting patient struct and nil as the error
	return &p, nil
}

func SympsCodeLength(dbconn *pgx.Conn) (int, error) {
	var MaxID int
	rows, err := dbconn.Query(context.Background(), "SELECT MAX(id) FROM symptoms_table;", []interface{}{})
	if err != nil {
		fmt.Println("Err: line 120 of diagnosis.go")
		return 0, err
	}
	defer rows.Close()
	err = rows.Scan(&MaxID)
	if err != nil {
		fmt.Println("Err: line 126 of diagnosis.go")
		return 0, err
	}
	codeLength := (MaxID + 1) / 8
	if MaxID-(codeLength*8)+1 > 0 {
		codeLength += 1
	}
	return codeLength, nil
}
