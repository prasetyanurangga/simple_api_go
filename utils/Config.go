package utils

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)


const (
	username string = "root"
	password string = ""
	host string = "localhost"
	port string = "3308"
	database string = "data_dummy"
)


// Choose Driver
func MySQL() (*sql.DB, error){

	viper.SetConfigType("json")
    viper.AddConfigPath(".")
    viper.SetConfigName("app.config")

    err := viper.ReadInConfig()
    if err != nil {
        log.Fatal(err)
	}
	
	dns := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v", 
		viper.GetString("database.user"), 
		viper.GetString("database.pass"),
		viper.GetString("database.host"),
		viper.GetString("database.port"), 
		viper.GetString("database.dbname"),
	)

	db, err := sql.Open("mysql", dns)
	if err != nil{
		return nil, err
	}

	return db, nil
}