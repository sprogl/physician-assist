package diagnosis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

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
	//Debug
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Err: line 35 of diagnosis")
		return nil, err
	}
	fmt.Println(string(bodyBytes))
	//Parse the posted form and extract it for further process
	jsonDecoder := json.NewDecoder(req.Body)
	form := struct {
		Gender   string `json:"gen"`
		Age      int    `json:"age"`
		Symptoms string `json:"symps"`
	}{}
	//Define the uninitialized patient data
	var p Patient
	err = jsonDecoder.Decode(&form)
	if err != nil {
		fmt.Println("Err: line 50 of diagnosis")
		return nil, err
	}
	//Check the gender input and set in inside the patient struct
	switch form.Gender {
	case "female":
		p.Gender = "Female"
	case "male":
		p.Gender = "Male"
	default:
		fmt.Println("Err: line 60 of diagnosis")
		return nil, errors.New("Wrong gender input format!")
	}
	//Check the age input and set in inside the patient struct
	if form.Age < 0 || form.Age > 100 {
		fmt.Println("Err: line 65 of diagnosis")
		return nil, errors.New("Wrong age input format!")
	} else {
		p.Age = form.Age
	}
	//Check the symptoms input and set in inside the patient struct
	//This is done through splitting the entry by commas
	seperator := regexp.MustCompile(" *(([,;](\r\n|\n)* *)|([,;]*(\r\n|\n) *))")
	p.Symptoms = seperator.Split(form.Symptoms, -1)
	if len(p.Symptoms) == 0 {
		fmt.Println("Err: line 75 of diagnosis")
		return nil, errors.New("Wrong symptom input format!")
	}
	//Return the resulting patient struct and nil as the error
	return &p, nil
}

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
