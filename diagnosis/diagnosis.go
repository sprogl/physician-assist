package diagnosis

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
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
	//Parse the posted form and extract it for further process
	req.ParseForm()
	form := req.PostForm
	//Define the uninitialized patient data
	var p Patient
	//Check the gender input and set in inside the patient struct
	switch form.Get("gen") {
	case "female":
		p.Gender = "Female"
	case "male":
		p.Gender = "Male"
	default:
		return nil, errors.New("Wrong gender input format!")
	}
	//Check the age input and set in inside the patient struct
	if i, err := strconv.Atoi(form.Get("age")); err != nil {
		return nil, errors.New("Wrong age input format!")
	} else {
		p.Age = i
	}
	//Check the symptoms input and set in inside the patient struct
	//This is done through splitting the entry by commas
	seperator := regexp.MustCompile(" *(([,;](\r\n|\n)* *)|([,;]*(\r\n|\n) *))")
	p.Symptoms = seperator.Split(form.Get("symps"), -1)
	if len(p.Symptoms) == 0 {
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
