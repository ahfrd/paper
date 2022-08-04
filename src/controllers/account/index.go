package controllers

import (
	"paper/src/database"
	"paper/src/helpers"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

// ControllerNamaFungsiObjectRes is a

type ControllerAccountReq struct {
	Id            int    `json:"id"`
	AccountNumber string `json:"accountNumber"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Address       string `json:"address"`
	Username      string `json:"username"`
	SessionId     string `json:"session_id"`

	PhoneNumber string `json:"phoneNumber"`
}
type ControllerLogoutReq struct {
	Username string `son:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type AccountStructRes struct {
	Id int `json:"id"`
}
type ControllerShowDataReq struct {
	Page          string `json:"page"`
	RecordPerPage string `json:"record_per_page"`
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
	database.TblAccount
}
