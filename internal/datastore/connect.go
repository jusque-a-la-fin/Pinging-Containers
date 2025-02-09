package datastore

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func CreateNewDB() (*sql.DB, error) {
	username := viper.GetString("postgre.username")
	password := viper.GetString("postgre.password")
	host := viper.GetString("postgre.host")
	port := viper.GetString("postgre.port")
	database := viper.GetString("postgre.database")
	sslmode := viper.GetString("postgre.sslmode")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, database, sslmode)
	dtb, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to database: %v", err)
	}
	return dtb, nil
}
