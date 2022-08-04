package database

import (
	"database/sql"
	"fmt"
	"paper/src/helpers"
)

// TblBwLimitGroup is a
type TblWallet struct {
	Database
	helpers.ELK
}
type StructurWallet struct {
	Id       int    `json:"id"`
	UserId   string `json:"user_id"`
	Ballance int    `json:"ballance"`
}

func (o TblWallet) InsertWallet(user_id int) (int64, *sql.DB, error) {
	var err error
	var res sql.Result
	var prepare *sql.Stmt
	db, err := o.ConnectDB()
	if err != nil {
		defer db.Close()
		return 0, nil, fmt.Errorf("%s", err)
	}
	defer db.Close()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	queryInsert := "INSERT INTO tbl_wallet (user_id) values (?)"
	prepare, err = db.Prepare(queryInsert)
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl_user SQL : %v", err)
	}
	res, err = prepare.Exec(user_id)

	if err != nil {
		return 0, db, fmt.Errorf("failed to insert error_general on tbl_user SQL : %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, db, fmt.Errorf("failed to populate status inserted : %v", err)
	}
	return count, db, err

}
func (o TblWallet) SelectWalletByUserId(userId int) (StructurWallet, *sql.DB, error) {
	var listWallet StructurWallet
	var user_id sql.NullString
	var ballance sql.NullInt64
	var id sql.NullInt64
	db, err := o.ConnectDB()
	defer db.Close()
	var query string = fmt.Sprintf(`select id, user_id, ballance  from tbl_wallet WHERE user_id ="%d"`, userId)
	err = db.QueryRow(query).Scan(
		&id,
		&user_id,
		&ballance,
	)
	listWallet.Id = int(id.Int64)
	listWallet.UserId = user_id.String
	listWallet.Ballance = int(ballance.Int64)

	defer db.Close()
	if err != nil && err != sql.ErrNoRows {
		return listWallet, db, fmt.Errorf("failed Select SQL for tbl_user : %v", err)
	}

	return listWallet, db, nil
}
func (o TblWallet) UpdateBallance(user_id int, sisahBallance int, updateAt string) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE tbl_wallet set ballance = ?,updated_at = ? where user_id = ?`
	prepare, err := db.Prepare(queryUpdate)
	res, err := prepare.Exec(sisahBallance, updateAt, user_id)
	defer db.Close()
	if err != nil {
		return 0, db, fmt.Errorf("failed to tbl_user SQL : %v", err)
	}

	if err != nil {
		return 0, db, fmt.Errorf("failed to tbl_user SQL : %v", err)
	}

	counter, err := res.RowsAffected()
	if err != nil {
		return 0, db, fmt.Errorf("failed to populate status updated : %v", err)
	}
	return counter, db, err
}
