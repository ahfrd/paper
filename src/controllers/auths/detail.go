package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"paper/constant"
	"paper/src/helpers"

	"github.com/dgrijalva/jwt-go"
)

// ControllerNamaFungsi is a
func (o ControllerStructure) GetDetailUser(w http.ResponseWriter, req *http.Request) {
	var bodyReq ControllerUserReq
	res := helpers.Response{}

	err := json.NewDecoder(req.Body).Decode(&bodyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	tokenJwt, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			res.Body.Code = constant.RCNoCookie
			res.Body.Msg = "No Cookie detect"
			res.Reply(w)
			return
		}
		res.Body.Code = constant.GeneralErrorCode
		res.Body.Msg = constant.GeneralErrorDesc
		res.Reply(w)
		return
	}
	tokenStr := tokenJwt.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return helpers.JwtKey(), nil
		})
	if err != nil {
		res.Body.Code = constant.GeneralErrorCode
		res.Body.Msg = constant.GeneralErrorDesc
		res.Reply(w)
		return
	}
	if !tkn.Valid {
		res.Body.Code = constant.RCTokenNotValid
		res.Body.Msg = "Token Not Valid"
		res.Reply(w)
		return
	}
	user, db, err := o.SelectByUsername(bodyReq.Username)
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = err.Error()
		res.Reply(w)
		return
	}
	resData := ShowAllDataRes{}
	resData.Id = user.Id
	resData.Email = user.Email
	resData.LastLogin = user.LastLogin
	resData.Status = user.Status
	resData.Username = bodyReq.Username
	resData.LoginRetry = user.LoginRetry
	resData.NextLogindate = user.NextLogindate
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)
	return
}
