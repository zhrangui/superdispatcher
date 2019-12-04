package db

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	user          = flag.String("user", "DevLogin", "the database user")
	password      = flag.String("password", "LoginDev1!", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "CTV-PARALLEL5.Production.CTV.ca", "the database server")
	database      = flag.String("database", "VideoWeb", "database name")
)

func TestInvalidConnectionString(t *testing.T) {
	flag.Parse()

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", *server, *user, *password, *database)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT COUNT(*) FROM AtlasUsers")
	if err != nil {
		t.Fatal("Prepare failed:", err.Error())
	}
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	if count <= 0 {
		t.Fail()
	}
}
