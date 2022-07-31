package diagnosis

import (
	"context"
	"fmt"
	"os"
	"testing"

	pgx "github.com/jackc/pgx/v4"
)

func TestDiagnose(t *testing.T) {
	DBAdrress := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DBUSER_TEST"), os.Getenv("DBPASS_TEST"), os.Getenv("DBIP_TEST"), os.Getenv("DBPORT_TEST"), os.Getenv("DATABASE_TEST"))
	dbconn, err := pgx.Connect(context.Background(), DBAdrress)
	if err != nil {
		t.Log("Err: line 19 of diagnosis_test.go")
		t.Fatal(err)
	}
	defer dbconn.Close(context.Background())
	initQs := []string{
		"DROP TABLE IF EXISTS user_table;",
		`
		CREATE TABLE user_table(
			id int not null unique,
			name varchar(64) not null,
			email varchar(320) not null unique,
			passhash char(44),
			primary key(id)
		);`,
		`	
		INSERT INTO user_table
		VALUES (
			1,
			'Amir Golparvar',
			'golp@gmail.com',
			'YWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWE='
		);`,
		`	
		INSERT INTO user_table
		VALUES (
			2,
			'Vahid Toomani',
			'tmn@gmail',
			'YWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWE='
		);`,
		`	
		INSERT INTO user_table
		VALUES (
			3,
			'Arash HoseiniMoghadam',
			'ahm.gmail.com',
			'YWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWE='
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
}
