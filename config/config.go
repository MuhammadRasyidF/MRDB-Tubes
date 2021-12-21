package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/godror/godror"
)

const (
	localUsername string = "TIF420_A10"
	localPassword string = "mercury"
	localHost     string = "103.157.96.115"
)

// HubToMySQL
func OracleSQL() (*sql.DB, error) {
	var (
		username string = os.Getenv("DB_USERNAME")
		password string = os.Getenv("DB_PASSWORD")
		host     string = os.Getenv("DB_HOST")
	)

	var dsn string
	fmt.Println(host)

	if host == "" {
		dsn = fmt.Sprintf("%v/%v@%v", localUsername, localPassword, localHost)
		fmt.Println("host kosong")
	} else {
		dsn = fmt.Sprintf("%v/%v@%v", username, password, host)
	}
	db, err := sql.Open("godror", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
