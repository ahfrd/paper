package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"paper/constant"
	"paper/src/helpers"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ControllerNamaFungsi is a
func (o ControllerStructure) GetAccount(w http.ResponseWriter, req *http.Request) {
	var bodyReq ControllerShowDataReq
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&bodyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	countTrx, db, err := o.CountTrxGeneralAccount()
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		return
	}
	if countTrx.TotalRow == "" {
		res.Body.Code = constant.NotFoundErrorCode
		res.Body.Msg = constant.NotFoundErrorCodeDesc
		res.Reply(w)
		return
	}
	page := bodyReq.Page
	recordPerPage := bodyReq.RecordPerPage
	floTotalRow, _ := strconv.ParseFloat(countTrx.TotalRow, 64)
	floRecordPerPage, _ := strconv.ParseFloat(recordPerPage, 64)
	floCurrentPage, _ := strconv.ParseFloat(page, 64)
	totalPage := math.Ceil(floTotalRow / floRecordPerPage)
	currentPage := page
	firstRecord := (floCurrentPage - 1) * floRecordPerPage
	startRecord := firstRecord + 1
	countData := CountData{
		TotalRecord:   countTrx.TotalRow,
		TotalPage:     strconv.FormatFloat(totalPage, 'f', 0, 64),
		RecordPerPage: recordPerPage,
		CurrentPage:   currentPage,
		StartRecord:   strconv.FormatFloat(startRecord, 'f', 0, 64),
	}
	listTransaction, db, err := o.SelectDataAccountPagin(recordPerPage, strconv.FormatFloat(firstRecord, 'f', 0, 64))
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		return
	}
	if len(listTransaction) == 0 {
		res.Body.Code = constant.NotFoundErrorCode
		res.Body.Msg = constant.NotFoundErrorCodeDesc
		res.Reply(w)
		return
	}

	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = DataPagination{
		Pagination: countData,
		RecordData: listTransaction,
	}
	res.Reply(w)
	return
}
func (o ControllerStructure) ControllerCreateAccount(w http.ResponseWriter, req *http.Request) {
	var request ControllerAccountReq
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	rand.Seed(time.Now().UnixNano())

	s := fmt.Sprintf("%016d", rand.Int63n(1e16))
	numberRandom := s[0:11]

	// numberRandom := s[0:11]
	firstname := request.FirstName
	lastname := request.LastName
	address := request.Address
	username := request.Username
	phonenumber := request.PhoneNumber
	createAccountId, db, err := o.InsertAccount(numberRandom, firstname, lastname, address, username, phonenumber)
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		log.Fatalln(err)
		return
	}
	resData := AccountStructRes{}
	resData.Id = int(createAccountId)
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)
	return
}
func (o ControllerStructure) ControllerDeleteAccount(w http.ResponseWriter, req *http.Request) {
	res := helpers.Response{}

	keys, ok := req.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
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
	id, _ := strconv.Atoi(keys[0])
	_, _, err = o.AccountDeleteGorm(id)
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Body.Data = id
		log.Fatalln(err)

		return
	}

	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Reply(w)

	return
}
func (o ControllerStructure) ControllerUpdateAccount(w http.ResponseWriter, req *http.Request) {
	var request ControllerAccountReq
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	firstname := request.FirstName
	lastname := request.LastName
	address := request.Address
	phoneNumber := request.PhoneNumber
	id := request.Id
	currentTime := time.Now()
	_, db, err := o.UpdateAccountById(firstname, lastname, address, phoneNumber, id, currentTime.String())
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		return
	}
	resData := AccountStructRes{}
	resData.Id = request.Id
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)
	return
}
func (o ControllerStructure) RestoreDataAccount(w http.ResponseWriter, req *http.Request) {
	res := helpers.Response{}
	keys, ok := req.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
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
	id, _ := strconv.Atoi(keys[0])
	result, _, err := o.RestoreAccountDeleted(id)
	if err != nil {
		res.Body.Code = constant.GeneralErrorCode
		res.Body.Msg = err.Error()
		res.Reply(w)
		return
	}

	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = result
	res.Reply(w)
	return
}
