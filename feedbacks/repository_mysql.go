package feedbacks

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
	table          = "Feedbacks"
	layoutDateTime = "2021-09-27 03:05:05"
)

// GetAll Feedbacks
func GetAll(ctx context.Context) ([]models.Feedback, error) {

	var Feedbacks []models.Feedback

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By feedback_id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var Feedback models.Feedback

		if err = rowQuery.Scan(&Feedback.Feedback_id,
			&Feedback.Feedback,
			&Feedback.User_id,
			&Feedback.Reviews_id); err != nil {
			return nil, err
		}

		Feedbacks = append(Feedbacks, Feedback)
	}

	return Feedbacks, nil
}

// Insert Feedbacks
func Insert(ctx context.Context, Feedback models.Feedback) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (feedback_id, feedback) values('%v','%v')", table,
		Feedback.Feedback_id,
		Feedback.Feedback,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Feedbacks
func Update(ctx context.Context, Feedback models.Feedback, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set feedback ='%s' where feedback_id = %s",
		table,
		Feedback.Feedback,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Feedbacks
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where feedback_id = %s", table, id)

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
