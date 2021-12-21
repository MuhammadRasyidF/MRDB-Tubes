package movie

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
	table          = "movies"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Movie
func GetAll(ctx context.Context) ([]models.Movie, error) {

	var movies []models.Movie

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By movie_id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var movie models.Movie

		if err = rowQuery.Scan(&movie.Movie_id,
			&movie.Name,
			&movie.Description,
			&movie.Release_date,
			&movie.Image_url,
			&movie.Creator_id,
			&movie.Rate_id,
			&movie.Star_id); err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// Insert Movie
func Insert(ctx context.Context, movie models.Movie) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	// queryText := fmt.Sprintf("INSERT INTO %v (movie_id, name, description, release_date, image_url, creator_id, rate_id, star_id) values('%v',%v, %v, %v, %v, %v, %v, %v)", table,
	// 	movie.Movie_id,
	// 	movie.Name,
	// 	movie.Description,
	// 	movie.Release_date,
	// 	movie.Image_url,
	// 	movie.Creator_id,
	// 	movie.Rate_id,
	// 	movie.Star_id,
	// )

	queryText := fmt.Sprintf("INSERT INTO %v (movie_id, name, description, release_date, image_url) values('%v','%v', '%v', '%v', '%v')", table,
		movie.Movie_id,
		movie.Name,
		movie.Description,
		movie.Release_date,
		movie.Image_url,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Movie
func Update(ctx context.Context, movie models.Movie, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set name ='%s', description ='%s', release_date ='%s', image_url = '%s' where movie_id = %s",
		table,
		movie.Name,
		movie.Description,
		movie.Release_date,
		movie.Image_url,
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
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where movie_id = %s", table, id)

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
