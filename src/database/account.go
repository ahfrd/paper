package database

import (
	"database/sql"
	"fmt"
	"paper/src/helpers"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// TblBwLimitGroup is a
type TblAccount struct {
	Database
	helpers.ELK
}

type TblDatAccount struct {
	Id            int    `json:"id"`
	AccountNumber int    `json:"accountNumber"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Address       string `json:"address"`
	Username      string `json:"username"`
	PhoneNumber   string `json:"phoneNumber"`
}

type FinancialAccount struct {
	Id            int    `gorm:"column:id:primaryKey"`
	AccountNumber int    `gorm:"column:accountNumber"`
	FirstName     string `gorm:"column:firstname"`
	LastName      string `gorm:"column:lastname"`
	Address       string `gorm:"column:address"`
	Username      string `gorm:"column:username"`
	PhoneNumber   string `gorm:"column:phoneNumber"`
	DeletedAt     gorm.DeletedAt
}

// SelectRowDataPackage merupakan fungsi untuk get data data_package by id_operator

func (o TblAccount) SelectDataAccount(uname string) (TblDatAccount, *sql.DB, error) {
	var listuser TblDatAccount
	var id sql.NullInt64
	var accountNumber sql.NullInt64
	var firstname sql.NullString
	var lastname sql.NullString
	var address sql.NullString
	var username sql.NullString
	var phonenumber sql.NullString
	db, err := o.ConnectDB()
	defer db.Close()
	var query string = fmt.Sprintf(`select id, accountNumber, firstname, lastname, address,username,phonenumber  from financial_account where username = "%s"`, uname)
	err = db.QueryRow(query).Scan(
		&id,
		&accountNumber,
		&firstname,
		&lastname,
		&address,
		&username,
		&phonenumber,
	)
	listuser.Id = int(id.Int64)
	listuser.AccountNumber = int(accountNumber.Int64)
	listuser.FirstName = firstname.String
	listuser.LastName = lastname.String
	listuser.Address = address.String
	listuser.Username = username.String
	listuser.PhoneNumber = phonenumber.String
	defer db.Close()
	if err != nil && err != sql.ErrNoRows {
		return listuser, db, fmt.Errorf("failed Select SQL for tbl_user : %v", err)
	}

	return listuser, db, nil
}
func (o TblAccount) InsertAccount(accountNumber string, firstname string, lastname string, address string, username string, phonenumber string) (int64, *sql.DB, error) {
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

	queryInsert := "INSERT INTO financial_account(accountNumber, firstname, lastname, address,username,phonenumber) values (?,?,?,?,?,?)"
	prepare, err = db.Prepare(queryInsert)
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert profile SQL : %v", err)
	}
	res, err = prepare.Exec(accountNumber, firstname, lastname, address, username, phonenumber)

	if err != nil {
		return 0, db, fmt.Errorf("failed to insert error_general on profile SQL : %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return count, db, fmt.Errorf("failed to populate status inserted : %v", err)
	}
	LastinsertId, err := res.LastInsertId()
	return LastinsertId, db, err
}

func (o TblAccount) AccountDelete(id int) (int64, *sql.DB, error) {
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

	queryUpdate := "DELETE FROM financial_account WHERE id = ?;"
	prepare, err = db.Prepare(queryUpdate)
	if err != nil {
		log.Warnf("failed to delete TblAccount SQL : %v", err)
		return 0, db, fmt.Errorf("failed to delete TblAccount SQL : %v", err)
	}
	res, err = prepare.Exec(id)

	if err != nil {
		log.Warnf("failed to delete limit on TblAccount SQL : %v", err)
		return 0, db, fmt.Errorf("failed to delete limit on TblAccount SQL : %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Warnf("failed to populate status delete : %v", err)
		return 0, db, fmt.Errorf("failed to populate status delete : %v", err)
	}
	return count, db, err
}
func (o TblAccount) AccountDeleteGorm(id int) (int64, *gorm.DB, error) {
	var financial_account FinancialAccount
	var err error
	db, err := o.ConnectDBGorm()
	if err != nil {
		panic("gagal")
	}
	result := db.Table("financial_account").Where("id", id).Delete(&financial_account)
	if result.Error != nil {
		return 0, db, result.Error
	}
	return 0, db, result.Error
}
func (o TblAccount) UpdateAccountById(firstname string, lastname string, address string, phoneNumber string, id int, updatedAt string) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE financial_account set firstname = ?, lastname = ?,address = ?, phoneNumber = ?, updated_at = ? where id = ?`
	prepare, err := db.Prepare(queryUpdate)
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl profile SQL : %v", err)

	}
	res, err := prepare.Exec(firstname, lastname, address, phoneNumber, updatedAt, id)
	defer db.Close()
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl profile SQL : %v", err)
	}

	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl profile SQL : %v", err)
	}

	counter, err := res.RowsAffected()
	if err != nil {
		return 0, db, fmt.Errorf("failed to populate status updated : %v", err)
	}
	return counter, db, err
}
func (o TblAccount) CountTrxGeneralAccount() (countDataTrans, *sql.DB, error) {
	var TrxCount countDataTrans

	db, err := o.ConnectDB()

	if err != nil {
		log.Warnf("failed Select SQL for account : %v", err)
		return TrxCount, db, fmt.Errorf("Failed Connect to Database")
	}

	formatQuery := fmt.Sprintf(`SELECT COUNT(id) AS total_row FROM financial_account where deleted_at is NULL;`)

	var query string = formatQuery
	db.QueryRow(query).Scan(
		&TrxCount.TotalRow,
	)

	defer db.Close()
	if err != nil {
		log.Warnf("failed Select SQL for tbl_bw_trx_general : %v", err)
		return TrxCount, db, fmt.Errorf("failed Select SQL for tbl_bw_trx_general : %v", err)
	}

	return TrxCount, db, nil
}
func (o TblAccount) SelectDataAccountPagin(paging string, firstRecord string) ([]TblDatAccount, *sql.DB, error) {
	var result []TblDatAccount
	var obj TblDatAccount
	db, err := o.ConnectDB()
	if err != nil {
		defer db.Close()
		return result, db, fmt.Errorf("%s", err)
	}
	defer db.Close()

	var query string = fmt.Sprintf(`
	SELECT id, accountNumber, firstname, lastname,address,username,phoneNumber
	from financial_account
	where deleted_at is NULL order by id DESC Limit %s,%s`, firstRecord, paging)

	rows, err := db.Query(query)
	if err != nil {
		return result, db, fmt.Errorf("failed Select Row SQL for table : %v", err)
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(
			&obj.Id, &obj.AccountNumber, &obj.FirstName,
			&obj.LastName, &obj.Address, &obj.Username, &obj.PhoneNumber)

		if err != nil {
			return result, db, fmt.Errorf("failed Select Row SQL for table : %v", err)
		}

		result = append(result, obj)
	}
	defer rows.Close()

	return result, db, nil
}
func (o TblAccount) RestoreAccountDeleted(id int) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE financial_account set deleted_at = NULL where id = ?`
	prepare, err := db.Prepare(queryUpdate)
	res, err := prepare.Exec(id)
	defer db.Close()
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl profile SQL : %v", err)
	}

	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl profile SQL : %v", err)
	}

	counter, err := res.RowsAffected()
	if err != nil {
		return 0, db, fmt.Errorf("failed to populate status updated : %v", err)
	}
	return counter, db, err
}
