package creators

import (
	"api-mysql/config"
	"api-mysql/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	table          = "creators"
	layoutDateTime = "2021-09-27 03:05:05"
)

// GetAll Creators
func GetAll(ctx context.Context) ([]models.Creators, error) {

	var Creators []models.Creators

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By Creator_id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var Creator models.Creators

		if err = rowQuery.Scan(&Creator.Creator_id,
			&Creator.Name); err != nil {
			return nil, err
		}

		Creators = append(Creators, Creator)
	}

	return Creators, nil
}

// Insert Creators
func Insert(ctx context.Context, Creator models.Creators) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (Creator_id, name) values('%v','%v')", table,
		Creator.Creator_id,
		Creator.Name,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Creators
func Update(ctx context.Context, Creator models.Creators, uname string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set name ='%s' where name = %s",
		table,
		Creator.Name,
		uname,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Creators
func Delete(ctx context.Context, uname string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where name = %s", table, uname)

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
