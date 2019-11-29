package db

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestInvalidConnectionString(t *testing.T) {
	var (
		userid   = flag.String("U", "WebSite", "login_id")
		password = flag.String("P", "LoginDev1!", "password")
		server   = flag.String("S", "CTV-PARALLEL5.Production.CTV.ca", "server_name[\\instance_name]")
		database = flag.String("d", "VideoWeb", "db_name")
	)
	flag.Parse()

	dsn := "server=" + *server + ";user id=" + *userid + ";password=" + *password + ";database=" + *database
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return
	}
	r := bufio.NewReader(os.Stdin)
	for {
		_, err = os.Stdout.Write([]byte("> "))
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println()
				return
			}
			fmt.Println(err)
			return
		}
		err = exec(db, cmd)
		if err != nil {
			fmt.Println(err)
		}
	}
}
