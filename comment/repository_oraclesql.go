package comment

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
	table = "TB_COMMENT"
)

// GetAll comment
func GetAll(ctx context.Context) ([]models.Tb_comment, error) {

	var comments []models.Tb_comment

	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Cant connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By commentid DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var comment models.Tb_comment

		if err = rowQuery.Scan(&comment.CommentId,
			comment.Comment,
			comment.UserId,
			comment.MovieId); err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

// Insert comment
func Insert(ctx context.Context, comment models.Tb_comment) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (commentid, comment) values('%v','%v')", table,
		comment.CommentId,
		comment.Comment,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update comment
func Update(ctx context.Context, comment models.Tb_comment, id string) error {

	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set comment ='%s' where commentid = %s",
		table,
		comment.Comment,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete comment
func Delete(ctx context.Context, id string) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where commentid = %s", table, id)

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
