package diagnosis

import (
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

func (p *Patient) IsFemale() bool {
	return (*p).Gender == "Female"
}

//This function processes the post request
//and extracts the sanitized input inside the post request
func FormProcess(req *http.Request) (*Patient, error) {
	//Define the uninitialized patient data
	var p Patient
	//Parse the posted form and extract it for further process
	jsonDecoder := json.NewDecoder(req.Body)
	// form := struct {
	// 	Gender   string `json:"gen"`
	// 	Age      int    `json:"age"`
	// 	Symptoms string `json:"symps"`
	// }{}
	err := jsonDecoder.Decode(&p)
	if err != nil {
		fmt.Println("Err: line 40 of diagnosis.go")
		return nil, err
	}
	//Check the gender input and set in inside the patient struct
	switch p.Gender {
	case "female":
		p.Gender = "Female"
	case "male":
		p.Gender = "Male"
	default:
		fmt.Println("Err: line 52 of diagnosis.go")
		return nil, errors.New("Wrong gender input format!")
	}
	//Check the age input and set in inside the patient struct
	if p.Age < 0 || p.Age > 150 {
		fmt.Println("Err: line 57 of diagnosis.go")
		return nil, errors.New("Wrong age input format!")
	}
	//Check the symptoms input and set in inside the patient struct
	//This is done through splitting the entry by commas
	// seperator := regexp.MustCompile(" *(([,;](\r\n|\n)* *)|([,;]*(\r\n|\n) *))")
	// p.Symptoms = seperator.Split(form.Symptoms, -1)
	if len(p.Symptoms) == 0 {
		fmt.Println("Err: line 65 of diagnosis.go")
		return nil, errors.New("Empty list of symptoms!")
	}
	//Return the resulting patient struct and nil as the error
	return &p, nil
}

//This function searchs the patient's data (symtoms, gender and age) in the database and
//finds the matched disease. It returns the result as list of matched diseases and
//any error occurred in the process.
func (*Patient) Diagnose(dbconn *pgx.Conn) ([]Disease, error) {
	return []Disease{Aids, Cancer}, nil
}
