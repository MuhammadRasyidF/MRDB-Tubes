package movie_creator

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
	table = "TB_GENRES"
)

// GetAll movie-creator
func GetAll(ctx context.Context) ([]models.Movie_creator, error) {

	var movieCreators []models.Movie_creator

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
		var movieCreator models.Movie_creator

		if err = rowQuery.Scan(&movieCreator.MovieId,
			&movieCreator.MovieId); err != nil {
			return nil, err
		}

		movieCreators = append(movieCreators, movieCreator)
	}

	return movieCreators, nil
}

// Insert movie-creator
func Insert(ctx context.Context, movieCreators models.Movie_creator) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (movieid, creatorid) values('%v','%v)", table,
		movieCreators.MovieId,
		movieCreators.CreatorId,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Delete movie-creator
func Delete(ctx context.Context, id int) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where cretorid = %d", table, id)

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
