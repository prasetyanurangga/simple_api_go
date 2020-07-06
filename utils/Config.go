package utils

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


const (
	username string = "root"
	password string = ""
	host string = "localhost"
	port string = "3308"
	database string = "data_dummy"
)

var (
	dns = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password,host,port, database)
)

// Choose Driver
func MySQL() (*sql.DB, error){
	db, err := sql.Open("mysql", dns)
	if err != nil{
		return nil, err
	}

	return db, nil
}