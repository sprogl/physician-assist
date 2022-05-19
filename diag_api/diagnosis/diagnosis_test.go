package diagnosis

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v4"
)

var pat1 = Patient{
	Gender:   "Male",
	Age:      25,
	Symptoms: []string{"Goh Gije", "Pare"},
}

var ds1 = []Disease{
	{
		Name:     "Aids",
		Symptoms: []string{"Pare"},
	},
	{
		Name:     "Cancer",
		Symptoms: []string{"Goh Gije", "Kap khoshkak"},
	},
	{
		Name:     "Kachali",
		Symptoms: []string{"Chet", "Pare"},
	},
}

var pat2 = Patient{
	Gender:   "Female",
	Age:      25,
	Symptoms: []string{"Goh Gije", "Pare"},
}

var ds2 = []Disease{
	{
		Name:     "Cancer",
		Symptoms: []string{"Goh Gije", "Kap khoshkak"},
	},
}

var pat3 = Patient{
	Gender:   "Male",
	Age:      40,
	Symptoms: []string{"Goh Gije", "Pare"},
}

var ds3 = []Disease{
	{
		Name:     "Cancer",
		Symptoms: []string{"Goh Gije", "Kap khoshkak"},
	},
	{
		Name:     "Kachali",
		Symptoms: []string{"Chet", "Pare"},
	},
}

func TestDiagnose(t *testing.T) {
	DBAdrress := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DBUSER_TEST"), os.Getenv("DBPASS_TEST"), os.Getenv("DBIP_TEST"), os.Getenv("DBPORT_TEST"), os.Getenv("DATABASE_TEST"))
	dbconn, err := pgx.Connect(context.Background(), DBAdrress)
	if err != nil {
		t.Log("Err: line 38 of diagnosis_test.go")
		t.Fatal(err)
	}
	defer dbconn.Close(context.Background())
	initQs := []string{
		"DROP TABLE IF EXISTS symps_table;",
		"DROP TABLE IF EXISTS dis_table;",
		"DROP TABLE IF EXISTS diag_table;",
		`
		CREATE TABLE symps_table(
			id int not null unique,
			name varchar(64) not null,
			primary key(id)
		);`,
		`
		CREATE TABLE dis_table(
			id int not null unique,
			name varchar(64) not null,
			primary key(id)
		);`,
		`
		CREATE TABLE diag_table(
			symp_id int not null,
			dis_id int not null,
			gen varchar(6),
			age_max int not null,
			age_min int not null
		);`,
		`	
		INSERT INTO symps_table
		VALUES (
			0,
			'Goh Gije'
		);`,
		`	
		INSERT INTO symps_table
		VALUES (
			1,
			'Chet'
		);`,
		`		
		INSERT INTO symps_table
		VALUES (
			2,
			'Pare'
		);`,
		`	
		INSERT INTO symps_table
		VALUES (
			3,
			'Kap khoshkak'
		);`,
		`	
		INSERT INTO dis_table
		VALUES (
			0,
			'Aids'
		);`,
		`	
		INSERT INTO dis_table
		VALUES (
			1,
			'Cancer'
		);`,
		`	
		INSERT INTO dis_table
		VALUES (
			2,
			'Koor'
		);`,
		`	
		INSERT INTO dis_table
		VALUES (
			3,
			'Kachali'
		);`,
		`	
		INSERT INTO diag_table
		VALUES (
			0,
			1,
			null,
			60,
			20
		);`,
		`	
		INSERT INTO diag_table
		VALUES (
			3,
			1,
			null,
			60,
			20
		);`,
		`	
		INSERT INTO diag_table
		VALUES (
			2,
			0,
			'Male',
			30,
			20
		);`,
		`	
		INSERT INTO diag_table
		VALUES (
			1,
			3,
			'Male',
			100,
			18
		);`,
		`	
		INSERT INTO diag_table
		VALUES (
			2,
			3,
			'Male',
			100,
			18
		);`,
	}
	for _, initQ := range initQs {
		rows, err := dbconn.Query(context.Background(), initQ)
		if err != nil {
			t.Fatal(err)
		}
		rows.Close()
	}
	//Test 1
	ds, err := (&pat1).Diagnose(dbconn)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
	if !cmp.Equal(ds, ds1) {
		t.Fatal("the first test failed")
	}
	//Test 2
	ds, err = (&pat2).Diagnose(dbconn)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
	if !cmp.Equal(ds, ds2) {
		t.Fatal("the first test failed")
	}
	//Test 3
	ds, err = (&pat3).Diagnose(dbconn)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ds)
	if !cmp.Equal(ds, ds3) {
		t.Fatal("the first test failed")
	}
}
