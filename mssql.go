package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug    = flag.Bool("debug", false, "enable debugging")
	password = flag.String("password", "majid1212@@!!", "the database password")
	port     = flag.Int("port", 1433, "the database port")
	server   = flag.String("server", "192.168.1.8", "the database server")
	user     = flag.String("user", "sa", "the database user")
	database = flag.String("database", "SpGolang", "the database name")
)

//GetDB ...
func GetDB() *sqlx.DB {
	flag.Parse()
	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", *server, *user, *password, *port, *database)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sqlx.Connect("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	fmt.Printf("Connected!\n")
	//defer conn.Close()
	return conn
}
