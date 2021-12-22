package movie_star

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
	table = "MOVIE-STAR"
)

// GetAll movie-star
func GetAll(ctx context.Context) ([]models.Movie_star, error) {

	var movieStars []models.Movie_star

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
		var movieStar models.Movie_star

		if err = rowQuery.Scan(&movieStar.MovieId,
			&movieStar.StarId); err != nil {
			return nil, err
		}

		movieStars = append(movieStars, movieStar)
	}

	return movieStars, nil
}

// Insert movie-star
func Insert(ctx context.Context, movieCreators models.Movie_star) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (movieid, starid) values('%v','%v)", table,
		movieCreators.MovieId,
		movieCreators.StarId,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Delete movie-star
func Delete(ctx context.Context, id int) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where starid = %d", table, id)

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
