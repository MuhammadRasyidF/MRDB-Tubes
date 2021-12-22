package rating

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
	table = "RATING"
)

// GetAll rating
func GetAll(ctx context.Context) ([]models.Rating, error) {

	var ratings []models.Rating

	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Cant connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By movieid DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var rating models.Rating

		if err = rowQuery.Scan(&rating.MovieId,
			&rating.UserId,
			&rating.Rate); err != nil {
			return nil, err
		}

		ratings = append(ratings, rating)
	}

	return ratings, nil
}

// Insert rating
func Insert(ctx context.Context, rating models.Rating) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (movieid, userid, rate) values('%v','%v','%v')", table,
		rating.MovieId,
		rating.UserId,
		rating.Rate,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update rating
func Update(ctx context.Context, rating models.Rating, movieid int, userid int) error {

	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set rate ='%v' where movieid = %d AND userid = %d",
		table,
		rating.Rate,
		movieid,
		userid,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete rating
func Delete(ctx context.Context, movieid int, userid int) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where movieid = %d AND userid = %d", table, movieid, userid)

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
