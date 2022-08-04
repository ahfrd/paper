package controllers

import (
	"paper/src/database"
	"paper/src/helpers"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

// ControllerNamaFungsiObjectRes is a

type ControllerUserReq struct {
	Username        string `son:"username"`
	Password        string `json:"password"`
	Login_Retry     int    `json:"login_retry"`
	Next_Login_date string `json:"next_login_date"`
	Email           string ` json:"email"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type ControllerLogoutReq struct {
	Username string `son:"username"`
}
type LoginStructRes struct {
	Username string `json:"username"`
	Email    string ` json:"email"`
}
type ShowAllDataRes struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	LoginRetry    int    `json:"login_retry"`
	NextLogindate string `json:"next_login_date"`
	LastLogin     string `json:"last_login"`
	Status        int    `json:"status"`
}
type ControllerStructure struct {
	helpers.ELK
	helpers.Response
	database.TblUser
	database.TblWallet
}
