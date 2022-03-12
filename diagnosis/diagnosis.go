package diagnosis

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
)

//Introduce the struct Disease and some method to export its content
type Disease struct {
	name     string
	symptoms []string
}

func (d *Disease) Name() string {
	return (*d).name
}

func (d *Disease) Symptoms() []string {
	return (*d).symptoms
}

//Introduce the struct Patient and some method to export its content
type Patient struct {
	male     bool
	age      int
	symptoms []string
}

func (p *Patient) Age() int {
	return (*p).age
}

func (p *Patient) Symptoms() []string {
	return (*p).symptoms
}

func (p *Patient) Gender() string {
	if (*p).male {
		return "Male"
	}
	return "Female"
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
		p.male = false
	case "male":
		p.male = true
	default:
		return nil, errors.New("Wrong gender input format!")
	}
	//Check the age input and set in inside the patient struct
	if i, err := strconv.Atoi(form.Get("age")); err != nil {
		return nil, errors.New("Wrong age input format!")
	} else {
		p.age = i
	}
	//Check the symptoms input and set in inside the patient struct
	//This is done through splitting the entry by commas
	seperator := regexp.MustCompile(" *(([,;](\r\n|\n)* *)|([,;]*(\r\n|\n) *))")
	p.symptoms = seperator.Split(form.Get("symps"), -1)
	if len(p.symptoms) == 0 {
		return nil, errors.New("Wrong symptom input format!")
	}
	//Return the resulting patient struct and nil as the error
	return &p, nil
}

var Cancer = Disease{
	name: "Cancer",
	symptoms: []string{
		"symp1",
		"symp2",
		"symp3",
	},
}

var Aids = Disease{
	name: "Aids",
	symptoms: []string{
		"symp3",
		"symp4",
		"symp5",
	},
}
