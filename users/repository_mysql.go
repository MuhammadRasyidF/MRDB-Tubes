package users

import (
	"api-mrdb/config"
	"api-mrdb/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	table          = "users"
	layoutDateTime = "2021-09-27 03:05:05"
)

// GetAll user
func GetAll(ctx context.Context) ([]models.Users, error) {
	var users []models.Users

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By user_id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var user models.Users

		if err = rowQuery.Scan(&user.User_id,
			&user.Username,
			&user.Password,
			&user.Role); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Insert user
func Insert(ctx context.Context, user models.Users) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	var role = "guest"

	queryText := fmt.Sprintf("INSERT INTO %v (user_id, username, password, role) values('%v','%v', '%v', '%v')", table,
		user.User_id,
		user.Username,
		user.Password,
		role,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update user
func Update(ctx context.Context, user models.Users, uname string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set username ='%s', password ='%s', role ='%s' where username = %s",
		table,
		user.Username,
		user.Password,
		user.Role,
		uname,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete user
func Delete(ctx context.Context, uname string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where username = %s", table, uname)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
