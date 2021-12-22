package likecomment

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
	table = "LIKECOMMENT"
)

// GetAll Likecomment
func GetAll(ctx context.Context) ([]models.LikeComment, error) {

	var likecomments []models.LikeComment

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
		var likecomment models.LikeComment

		if err = rowQuery.Scan(&likecomment.CommentId,
			&likecomment.UserId); err != nil {
			return nil, err
		}

		likecomments = append(likecomments, likecomment)
	}

	return likecomments, nil
}

// Insert Likecomment
func Insert(ctx context.Context, likecomment models.LikeComment) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (commentid, userid) values('%v','%v')", table,
		likecomment.CommentId,
		likecomment.UserId,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Delete Likecomment
func Delete(ctx context.Context, commentid int, userid int) error {
	db, err := config.OracleSQL()

	if err != nil {
		log.Fatal("Can't connect to OracleSQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where commentid = %d AND userid = %d", table, commentid, userid)

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
