package controllers

import (
	"paper/src/database"
	"paper/src/helpers"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

// ControllerNamaFungsiObjectRes is a

type ControllerCreateTransactionReq struct {
	AccountNumber   int    `json:"accountNumber"`
	ProductName     string `json:"productName"`
	ProductCategory string `json:"productCategory"`
	Nominal         int    `json:"nominal"`
}
type ControllerShowDataReq struct {
	AccountNumber int    `json:"accountNumber"`
	Page          string `json:"page"`
	RecordPerPage string `json:"record_per_page"`
}
type ControllerTransactionSummary struct {
	Periode string `json:"periode" validate:"enum=daily,montly"`
}
type TransactionShowIdStructRes struct {
	Id int `json:"id"`
}

type SummaryTransaction struct {
	Total int    `json:"total_nominal"`
	Date  string `json:"date"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type ControllerUpdateTransReq struct {
	Id              int    `json:"id"`
	AccountNumber   int    `json:"accountNumber"`
	ProductName     string `json:"productName"`
	ProductCategory string `json:"productCategory"`
	Nominal         int    `json:"nominal"`
}
type CountData struct {
	TotalRecord   string `json:"total_record"`
	TotalPage     string `json:"total_page"`
	RecordPerPage string `json:"record_per_page"`
	CurrentPage   string `json:"current_page"`
	StartRecord   string `json:"start_record"`
}
type DataPagination struct {
	Pagination CountData   `json:"pagination"`
	RecordData interface{} `json:"record_data"`
}
type ControllerStructure struct {
	helpers.ELK
	helpers.Response
	database.TblTrans
}
