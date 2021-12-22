package creators

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
	table = "TB_CREATORS"
)

// GetAll Creators
func GetAll(ctx context.Context) ([]models.Tb_creators, error) {

	var creators []models.Tb_creators

	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Cant connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By creatorid DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var creator models.Tb_creators

		if err = rowQuery.Scan(&creator.CreatorId,
			&creator.Name); err != nil {
			return nil, err
		}

		creators = append(creators, creator)
	}

	return creators, nil
}

// Insert Creators
func Insert(ctx context.Context, creator models.Tb_creators) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (Creator_id, name) values('%v','%v')", table,
		creator.CreatorId,
		creator.Name,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Creators
func Update(ctx context.Context, creator models.Tb_creators, uname string) error {

	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set name ='%s' where name = %s",
		table,
		creator.Name,
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
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
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
