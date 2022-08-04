package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"paper/constant"
	"paper/src/helpers"
	"strings"
)

func (o ControllerStructure) ControllerRegister(w http.ResponseWriter, req *http.Request) {
	var bodyReq ControllerUserReq
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&bodyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalln(err)
		return
	}
	email := bodyReq.Email
	password := bodyReq.Password
	username := bodyReq.Username
	validationError := helpers.VerifyPassword(password)
	fmt.Println(bodyReq.Password)
	if validationError != nil {
		res.Body.Code = constant.RCValidationError
		res.Body.Msg = validationError.Error()
		res.Reply(w)
		return
	}
	_, errMail := mail.ParseAddress(email)
	if errMail != nil {
		res.Body.Code = constant.RCValidationError
		res.Body.Msg = errMail.Error()
		res.Reply(w)
		return
	}
	login, db, err := o.SelectByUsername(username)
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		return
	}
	if login.Username != "" {
		res.Body.Code = constant.RCDataAlreadyExist
		res.Body.Msg = "Sorry, please check your data"
		res.Reply(w)
		return
	}
	combine := strings.ToUpper(username) + password
	hash := []byte(combine)
	hash_byte := sha256.Sum256(hash)
	hash_str := hex.EncodeToString(hash_byte[:])

	LastInsertId, db, err := o.Register(username, hash_str, email)
	db.Close()
	fmt.Println(LastInsertId)
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		return
	}
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Reply(w)
	return
}
