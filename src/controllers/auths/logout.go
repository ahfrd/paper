package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"paper/constant"
	"paper/src/helpers"
)

func (o ControllerStructure) ControllerLogout(w http.ResponseWriter, req *http.Request) {
	var bodyReq ControllerLogoutReq
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&bodyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	username := bodyReq.Username

	login, db, err := o.SelectByUsername(username)
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		return
	}
	if login.Username == "" {
		res.Body.Code = constant.NotFoundErrorCode
		res.Body.Msg = "Sorry, please check your data"
		res.Reply(w)
		return
	}
	if login.Status == 0 {
		res.Body.Code = constant.GeneralErrorCode
		res.Body.Msg = "Sorry, Account not login yet"
		res.Reply(w)
		return
	}
	_, db, err = o.UpdateStatusLogin(username, 0)
	db.Close()
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
