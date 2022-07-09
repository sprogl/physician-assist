package diagnosis

import (
	"context"
	"encoding/json"
	"errors"
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

//This function searchs the patient's data (symtoms, gender and age) in the database and
//finds the matched disease. It returns the result as list of matched diseases and
//any error occurred in the process.
func (pat *Patient) Diagnose(conn *pgx.Conn) ([]Disease, error) {
	//Query to the diseases database
	mainQ := `
	SELECT name AS Disease, Symptom
	FROM (
		SELECT dis_id, name AS Symptom
		FROM (
			SELECT symp_id, dis_id
			FROM ((
				SELECT diag_table.dis_id AS disease_id
				FROM symps_table INNER JOIN diag_table ON symps_table.id=diag_table.symp_id
				WHERE (diag_table.gen IS NULL OR diag_table.gen=$1) AND (diag_table.age_min <= $2 AND diag_table.age_max >= $2) AND (symps_table.name = ANY($3))
			) As q_table1 INNER JOIN diag_table ON diag_table.dis_id=q_table1.disease_id)
		) As q_table2 INNER JOIN symps_table ON q_table2.symp_id=symps_table.id
	) As q_table3 INNER JOIN dis_table ON q_table3.dis_id=dis_table.id
	ORDER By Disease;
	`
	rows, err := conn.Query(context.Background(), mainQ, (*pat).Gender, (*pat).Age, (*pat).Symptoms)
	if err != nil {
<<<<<<< HEAD
=======
		fmt.Println("Err: line 70 of diagnosis.go")
>>>>>>> ab318a64e67f4eab5c00a6a37e4a2c30f1f318eb
		return nil, err
	}
	defer rows.Close()
	//Define the uninitialized disease (d) and disease slice (ds)
	var ds []Disease
	var d Disease
	var s1 string
	var s2 string
	//Loop over the rows ofdatabase response and populate the disease slice with the return data
	for rows.Next() {
		err := rows.Scan(&s1, &s2)
		if err != nil {
<<<<<<< HEAD
=======
			fmt.Println("Err: line 83 of diagnosis.go")
>>>>>>> ab318a64e67f4eab5c00a6a37e4a2c30f1f318eb
			return nil, err
		}
		if s1 == d.Name {
			d.Symptoms = append(d.Symptoms, s2)
<<<<<<< HEAD
		} else {
			if d.Name != "" {
				ds = append(ds, d)
			}
			d.Name = s1
			d.Symptoms = []string{s2}
		}
	}
	if d.Name != "" {
		ds = append(ds, d)
=======
		} else if d.Name == "" {
			d.Name = s1
			d.Symptoms = []string{s2}
		} else {
			ds = append(ds, d)
			d.Name = s1
			d.Symptoms = []string{s2}
		}
>>>>>>> ab318a64e67f4eab5c00a6a37e4a2c30f1f318eb
	}
	ds = append(ds, d)
	//Return the result
<<<<<<< HEAD
=======
	// return ds, nil
>>>>>>> ab318a64e67f4eab5c00a6a37e4a2c30f1f318eb
	return ds, nil
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
<<<<<<< HEAD
=======
		fmt.Println("Err: line 109 of diagnosis.go")
>>>>>>> ab318a64e67f4eab5c00a6a37e4a2c30f1f318eb
		return nil, err
	}
	//Check the gender input and set it inside the patient struct
	switch p.Gender {
	case "female":
		p.Gender = "Female"
	case "male":
		p.Gender = "Male"
	default:
<<<<<<< HEAD
=======
		fmt.Println("Err: line 119 of diagnosis.go")
>>>>>>> ab318a64e67f4eab5c00a6a37e4a2c30f1f318eb
		return nil, errors.New("wrong gender input format")
	}
	//Check the age input and set in inside the patient struct
	if p.Age < 0 || p.Age > 150 {
<<<<<<< HEAD
=======
		fmt.Println("Err: line 124 of diagnosis.go")
>>>>>>> ab318a64e67f4eab5c00a6a37e4a2c30f1f318eb
		return nil, errors.New("wrong age input format")
	}
	//Check the symptoms input and set in inside the patient struct
	//This is done through splitting the entry by commas
	// seperator := regexp.MustCompile(" *(([,;](\r\n|\n)* *)|([,;]*(\r\n|\n) *))")
	// p.Symptoms = seperator.Split(form.Symptoms, -1)
	if len(p.Symptoms) == 0 {
<<<<<<< HEAD
=======
		fmt.Println("Err: line 132 of diagnosis.go")
>>>>>>> ab318a64e67f4eab5c00a6a37e4a2c30f1f318eb
		return nil, errors.New("empty list of symptoms")
	}
	//Return the resulting patient struct and nil as the error
	return &p, nil
}
