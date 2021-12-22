package movie

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
	table          = "TB_MOVIES"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Movie
func GetAll(ctx context.Context) ([]models.Tb_movies, error) {

	var movies []models.Tb_movies

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
		var movie models.Tb_movies

		if err = rowQuery.Scan(&movie.MovieId,
			&movie.Name,
			&movie.Description,
			&movie.ReleaseDate,
			&movie.ImageUrl); err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// Insert Movie
func Insert(ctx context.Context, movie models.Tb_movies) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (name, description, releasedate, imageurl) values('%v', '%v', '%v', '%v')", table,
		movie.Name,
		movie.Description,
		movie.ReleaseDate,
		movie.ImageUrl,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Movie
func Update(ctx context.Context, movie models.Tb_movies, id string) error {

	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set name ='%s', description ='%s', releasedate ='%s', imageurl = '%s' where movieid = %s",
		table,
		movie.Name,
		movie.Description,
		movie.ReleaseDate,
		movie.ImageUrl,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Movie
func Delete(ctx context.Context, id string) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where movieid = %s", table, id)

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
