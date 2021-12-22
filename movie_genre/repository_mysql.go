package genre_category

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
	table          = "genre_category"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll genre_category
func GetAll(ctx context.Context) ([]models.Genre_category, error) {

	var genre_categorys []models.Genre_category

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
		var genre_category models.Genre_category

		if err = rowQuery.Scan(&genre_category.Movie_id,
			&genre_category.Genre_id); err != nil {
			return nil, err
		}

		genre_categorys = append(genre_categorys, genre_category)
	}

	return genre_categorys, nil
}

// Insert genre_category
func Insert(ctx context.Context, genre_category models.Genre_category) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (movie_id, genre_id) values('%v','%v')", table,
		genre_category.Movie_id,
		genre_category.Genre_id,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update genre_category
func Update(ctx context.Context, genre_category models.Genre_category, movie_id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set genre_id =%d where movie_id = %s",
		table,
		genre_category.Genre_id,
		movie_id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete genre_category
func Delete(ctx context.Context, genre_id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where genre_id = %s", table, genre_id)

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
