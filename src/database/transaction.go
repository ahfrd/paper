package database

import (
	"database/sql"
	"fmt"
	"paper/src/helpers"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// TblBwLimitGroup is a
type TblTrans struct {
	Database
	helpers.ELK
}
type ObjTransactionSum struct {
	Nominal         int    `json:"nominal"`
	TransactionDate string `json:"transaction_date"`
}
type TblDataTransaction struct {
	Id              int    `json:"id"`
	AccountNumber   int    `json:"accountNumber"`
	ProductName     string `json:"productName"`
	ProductCategory string `json:"productCategory"`
	Nominal         string `json:"nominal"`
	TransactionDate string `json:"transaction_date"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeltedAt        string `json:"deleted_at"`
}
type TblDataTransactionRes struct {
	Id              int    `json:"id"`
	AccountNumber   int    `json:"accountNumber"`
	ProductName     string `json:"productName"`
	ProductCategory string `json:"productCategory"`
	Nominal         string `json:"nominal"`
	TransactionDate string `json:"transaction_date"`
}
type countDataTrans struct {
	TotalRow string `json:"total_row"`
}
type FinancialTransaction struct {
	Id              int            `gorm:"column:id:primaryKey"`
	AccountNumber   int            `gorm:"column:accountNumber"`
	ProductName     string         `gorm:"column:productName"`
	ProductCategory string         `gorm:"column:productCategory"`
	Nominal         string         `gorm:"column:nominal"`
	CreatedAt       string         `gorm:"column:created_at"`
	UpdatedAt       string         `gorm:"column:updated_at"`
	DeltedAt        gorm.DeletedAt `gorm:"column:deleted_at"`
}

// SelectRowDataPackage merupakan fungsi untuk get data data_package by id_operator

func (o TblTrans) SelectDataTransaction(accountNumbers int) ([]TblDataTransaction, *sql.DB, error) {
	var result []TblDataTransaction
	var obj TblDataTransaction
	db, err := o.ConnectDB()
	if err != nil {
		defer db.Close()
		return result, db, fmt.Errorf("%s", err)
	}
	defer db.Close()

	var query string = fmt.Sprintf(`
	SELECT id, accountNumber, productName, productCategory,nominal,created_at,updated_at,deleted_at
	from financial_transaction
	where accountNumber = "%d" and deleted_at is NULL`, accountNumbers)

	rows, err := db.Query(query)
	if err != nil {
		return result, db, fmt.Errorf("failed Select Row SQL for table : %v", err)
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(
			&obj.Id, &obj.AccountNumber, &obj.ProductName,
			&obj.ProductCategory, &obj.Nominal,
			&obj.CreatedAt, &obj.UpdatedAt)

		if err != nil {
			return result, db, fmt.Errorf("failed Select Row SQL for assesment_privy : %v", err)
		}

		result = append(result, obj)
	}
	defer rows.Close()

	return result, db, nil
}
func (o TblTrans) InsertTransaction(accountNumber int, productName string, productCategory string, nominal int, currentTime string) (int64, *sql.DB, error) {
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

	queryInsert := "INSERT INTO financial_transaction( accountNumber, productName, productCategory,nominal, transaction_date) values (?,?,?,?,?)"
	prepare, err = db.Prepare(queryInsert)
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert profile SQL : %v", err)
	}
	res, err = prepare.Exec(accountNumber, productName, productCategory, nominal, currentTime)

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

func (o TblTrans) TransactionDelete(id int) (int64, *sql.DB, error) {
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

	queryUpdate := "DELETE FROM financial_transaction WHERE id = ?;"
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
func (o TblTrans) TransactionDeleteGorm(id int) (int64, *gorm.DB, error) {
	var financial_account FinancialAccount
	var err error
	db, err := o.ConnectDBGorm()
	if err != nil {
		return 0, db, err
	}
	result := db.Table("financial_transaction").Where("id", id).Delete(&financial_account)
	if result.Error != nil {
		return 0, db, result.Error
	}
	return 0, db, result.Error
}
func (o TblTrans) UpdateTransactionById(id int, productName string, productCategory string, nominal int, updatedAt string) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE financial_transaction set productName = ?, productCategory = ?,nominal = ?, updated_at = ? where id = ?`
	prepare, err := db.Prepare(queryUpdate)
	res, err := prepare.Exec(productName, productCategory, nominal, updatedAt, id)
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

func (o TblTrans) SelectDataTransactionSummary(periode string) ([]ObjTransactionSum, *sql.DB, error) {
	var result []ObjTransactionSum
	var obj ObjTransactionSum
	db, err := o.ConnectDB()
	if err != nil {
		defer db.Close()
		return result, db, fmt.Errorf("%s", err)
	}
	defer db.Close()
	// transSummaryPeriode := fmt.Sprintf(`DATE_FORMAT(transaction_date,"%Y-%m-%d")`)

	// if periode == "montly" {
	// 	transSummaryPeriode = fmt.Sprintf(`DATE_FORMAT(transaction_date,"%s")`, date)
	// }
	querySummary := "SELECT sum(nominal),date_format(transaction_date,'%Y-%m-%d') from financial_transaction where deleted_at is NULL group by date_format(transaction_date,'%Y-%m-%d')"

	if periode == "montly" {
		querySummary = "SELECT sum(nominal),date_format(transaction_date,'%Y-%m') from financial_transaction where deleted_at is NULL group by date_format(transaction_date,'%Y-%m')"
	}
	var query = querySummary
	rows, err := db.Query(query)
	if err != nil {
		return result, db, fmt.Errorf("failed Select Row SQL for table : %v", err)
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(
			&obj.Nominal,
			&obj.TransactionDate)

		if err != nil {
			return result, db, fmt.Errorf("failed Select Row SQL for assesment_privy : %v", err)
		}

		result = append(result, obj)
	}
	defer rows.Close()

	return result, db, nil
}
func (o TblTrans) CountTrxGeneral() (countDataTrans, *sql.DB, error) {
	var TrxCount countDataTrans

	db, err := o.ConnectDB()

	if err != nil {
		log.Warnf("failed Select SQL for tbl_bw_trx_general : %v", err)
		return TrxCount, db, fmt.Errorf("Failed Connect to Database")
	}

	formatQuery := fmt.Sprintf(`SELECT COUNT(id) AS total_row FROM financial_transaction where deleted_at is NULL;`)

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
func (o TblTrans) SelectDataTransactionPagin(accountNumber int, paging string, firstRecord string) ([]TblDataTransactionRes, *sql.DB, error) {
	var result []TblDataTransactionRes
	var obj TblDataTransactionRes
	db, err := o.ConnectDB()
	if err != nil {
		defer db.Close()
		return result, db, fmt.Errorf("%s", err)
	}
	defer db.Close()

	var query string = fmt.Sprintf(`
	SELECT id, accountNumber, productName, productCategory,nominal,transaction_date
	from financial_transaction
	where accountNumber = %d and deleted_at is NULL order by transaction_date DESC Limit %s,%s`, accountNumber, firstRecord, paging)

	rows, err := db.Query(query)
	if err != nil {
		return result, db, fmt.Errorf("failed Select Row SQL for table : %v", err)
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(
			&obj.Id, &obj.AccountNumber, &obj.ProductName,
			&obj.ProductCategory, &obj.Nominal, &obj.TransactionDate)

		if err != nil {
			return result, db, fmt.Errorf("failed Select Row SQL for table : %v", err)
		}

		result = append(result, obj)
	}
	defer rows.Close()

	return result, db, nil
}

func (o TblTrans) RestoreDataDeleted(id int) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE financial_transaction set deleted_at = NULL where id = ?`
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
