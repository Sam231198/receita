package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func ConnectDB() *sql.DB {
	host := viper.GetString("db.host")
	database := viper.GetString("db.database")
	// port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	pass := viper.GetString("db.pass")
	drive := viper.GetString("db.drive")

	// postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable
	dsn_string := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", user, pass, host, database)

	db, err := sql.Open(drive, dsn_string)

	if err != nil {
		fmt.Print("Error ao conectar ao banco de dados: ", err.Error())
	}

	return db
}
