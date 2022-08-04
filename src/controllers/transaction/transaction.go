package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"paper/constant"
	"paper/src/helpers"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ControllerNamaFungsi is a
func (o ControllerStructure) GetTransaction(w http.ResponseWriter, req *http.Request) {
	var bodyReq ControllerShowDataReq
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&bodyReq)
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
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
	countTrx, db, err := o.CountTrxGeneral()
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
	accountNumber := bodyReq.AccountNumber
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
	// accountNumber := bodyReq.AccountNumber
	listTransaction, db, err := o.SelectDataTransactionPagin(accountNumber, recordPerPage, strconv.FormatFloat(firstRecord, 'f', 0, 64))
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
func (o ControllerStructure) ControllerInsertTransaction(w http.ResponseWriter, req *http.Request) {
	var request ControllerCreateTransactionReq
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
	accountNumber := request.AccountNumber
	productName := request.ProductName
	productCategory := request.ProductCategory
	nominal := request.Nominal
	currentTime := time.Now()
	createAccountId, db, err := o.InsertTransaction(accountNumber, productName, productCategory, nominal, currentTime.String())
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		log.Fatalln(err)
		return
	}
	resData := TransactionShowIdStructRes{}
	resData.Id = int(createAccountId)
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)
	return
}
func (o ControllerStructure) ControllerDeleteTransaction(w http.ResponseWriter, req *http.Request) {
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
	_, _, err = o.TransactionDeleteGorm(id)
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
func (o ControllerStructure) ControllerUpdateTransaction(w http.ResponseWriter, req *http.Request) {
	var request ControllerUpdateTransReq
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := request.Id
	productName := request.ProductName
	productCategory := request.ProductCategory
	nominal := request.Nominal
	currentTime := time.Now()
	_, db, err := o.UpdateTransactionById(id, productName, productCategory, nominal, currentTime.String())
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		log.Fatalln(err)
		return
	}
	resData := TransactionShowIdStructRes{}
	resData.Id = request.Id
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)
	return
}
func (o ControllerStructure) ShowDeleteData(w http.ResponseWriter, req *http.Request) {
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
	result, _, err := o.RestoreDataDeleted(id)
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
func (o ControllerStructure) GetTransactionSummary(w http.ResponseWriter, req *http.Request) {
	var bodyReq ControllerTransactionSummary
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&bodyReq)
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
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
	periode := bodyReq.Periode
	if strings.ToLower(periode) != "montly" && strings.ToLower(periode) != "daily" {
		res.Body.Code = constant.GeneralErrorCode
		res.Body.Msg = "Maaf, periode hanya boleh daily dan montly"
		res.Reply(w)
		return
	}
	resultSummaryTrans, db, err := o.SelectDataTransactionSummary(periode)
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Code = constant.RCDatabaseErrorDesc
		res.Reply(w)
		return
	}
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resultSummaryTrans
	res.Reply(w)
	return
}
