package database

import (
	"database/sql"
	"fmt"
	"paper/src/helpers"
)

// TblBwLimitGroup is a
type TblUser struct {
	Database
	helpers.ELK
}

type StructureUser struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	LoginRetry    int    `json:"login_retry"`
	NextLogindate string `json:"next_login_date"`
	LastLogin     string `json:"last_login"`
	Status        int    `json:"status"`
}

func (o TblUser) Register(username string, password string, email string) (int64, *sql.DB, error) {
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

	queryInsert := "INSERT INTO tbl_user (username,password,email) values (?,?,?)"
	prepare, err = db.Prepare(queryInsert)
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl_user SQL : %v", err)
	}
	res, err = prepare.Exec(username, password, email)

	if err != nil {
		return 0, db, fmt.Errorf("failed to insert error_general on tbl_user SQL : %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return count, db, fmt.Errorf("failed to populate status inserted : %v", err)
	}
	LastinsertId, err := res.LastInsertId()
	return LastinsertId, db, err

}
func (o TblUser) SelectByUsername(username string) (StructureUser, *sql.DB, error) {
	var listuser StructureUser
	var id sql.NullInt64
	var uname sql.NullString
	var rowspas sql.NullString
	var loginretry sql.NullInt64
	var nextlogindate sql.NullString
	var email sql.NullString
	var lastLogin sql.NullString
	var status sql.NullInt64
	db, err := o.ConnectDB()
	defer db.Close()
	var query string = fmt.Sprintf(`select id, username, password, login_retry, next_login_date,email,last_login,status  from tbl_user WHERE username ="%s"`, username)
	err = db.QueryRow(query).Scan(
		&id,
		&uname,
		&rowspas,
		&loginretry,
		&nextlogindate,
		&email,
		&lastLogin,
		&status,
	)
	listuser.Id = int(id.Int64)
	listuser.Username = uname.String
	listuser.Password = rowspas.String
	listuser.LoginRetry = int(loginretry.Int64)
	listuser.NextLogindate = nextlogindate.String
	listuser.Email = email.String
	listuser.LastLogin = lastLogin.String
	listuser.Status = int(status.Int64)
	defer db.Close()
	if err != nil && err != sql.ErrNoRows {
		return listuser, db, fmt.Errorf("failed Select SQL for tbl_user : %v", err)
	}

	return listuser, db, nil
}
func (o TblUser) UpdateUsernameLoginRetry(count int, login_again string, username string) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE tbl_user set login_retry = ?, next_login_date = ? where username = ?`
	prepare, err := db.Prepare(queryUpdate)
	res, err := prepare.Exec(count, login_again, username)
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
func (o TblUser) UpdateUsernameLoginRetrySetCount(count int, username string) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE tbl_user set login_retry = ? where username = ?`
	prepare, err := db.Prepare(queryUpdate)
	res, err := prepare.Exec(count, username)
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
func (o TblUser) UpdateLastLoginAndStatus(username string, loginTime string, status int) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE tbl_user set last_login = ?,status = ? where username = ?`
	prepare, err := db.Prepare(queryUpdate)
	res, err := prepare.Exec(loginTime, status, username)
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
func (o TblUser) UpdateSessionId(username string, loginTime string) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE tbl_user set session_id = ? where username = ?`
	prepare, err := db.Prepare(queryUpdate)
	res, err := prepare.Exec(loginTime, username)
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
func (o TblUser) UpdateStatusLogin(username string, status int) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE tbl_user set status = ? where username = ?`
	prepare, err := db.Prepare(queryUpdate)
	res, err := prepare.Exec(status, username)
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
