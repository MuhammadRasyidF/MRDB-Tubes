package movie_genre

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
	table = "MOVIE-GENRE"
)

// GetAll movie-genre
func GetAll(ctx context.Context) ([]models.Movie_genre, error) {

	var movieGenres []models.Movie_genre

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
		var movieGenre models.Movie_genre

		if err = rowQuery.Scan(&movieGenre.MovieId,
			&movieGenre.GenreId); err != nil {
			return nil, err
		}

		movieGenres = append(movieGenres, movieGenre)
	}

	return movieGenres, nil
}

// Insert movie-genre
func Insert(ctx context.Context, movieGenre models.Movie_genre) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (movieid, genreid) values('%v','%v')", table,
		movieGenre.MovieId,
		movieGenre.GenreId,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update movie-genre
func Update(ctx context.Context, movieGenre models.Movie_genre, movieId int) error {

	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set genreid = %d where movieid = %d",
		table,
		movieGenre.GenreId,
		movieId,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete movie-genre
func Delete(ctx context.Context, genreId int, movieId int) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where genreid = %v AND movieid = %v", table, genreId, movieId)

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
