package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"paper/constant"
	"paper/src/helpers"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (o ControllerStructure) ControllerLogin(w http.ResponseWriter, req *http.Request) {

	var bodyReq ControllerUserReq
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&bodyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	password := bodyReq.Password
	username := bodyReq.Username
	combine := strings.ToUpper(username) + password
	hash := []byte(combine)
	hash_byte := sha256.Sum256(hash)
	hash_str := hex.EncodeToString(hash_byte[:])

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
		res.Body.Msg = "Username / Password salah"
		res.Reply(w)
		return
	}

	currentTime := time.Now()
	if currentTime.String() <= login.NextLogindate {
		res.Body.Code = constant.RCUserCannotLogin
		res.Body.Msg = fmt.Sprint("User anda dapat login pada ", login.NextLogindate)
		res.Reply(w)
		return
	}
	if login.Password != hash_str {
		i := login.LoginRetry
		count := i + 1
		if count > 3 {
			math := int(math.Pow((float64(count)-3), 2) * 600)
			login_again := time.Now().Add(time.Second * time.Duration(math))
			_, db, errs := o.UpdateUsernameLoginRetry(count, login_again.String(), username)
			db.Close()
			if errs != nil {
				res.Body.Code = constant.RCDatabaseError
				res.Body.Msg = fmt.Sprintf("%v", err.Error())
				res.Reply(w)
				return
			}
		} else {
			_, db, err := o.UpdateUsernameLoginRetrySetCount(count, username)
			db.Close()
			if err != nil {
				res.Body.Code = constant.RCDatabaseError
				res.Body.Msg = fmt.Sprintf("%v", err.Error())
				res.Reply(w)
				return
			}
		}
		res.Body.Code = constant.NotFoundErrorCode
		res.Body.Msg = "Username / Password salah"
		res.Reply(w)
		return
	} else {
		_, db, err := o.UpdateUsernameLoginRetrySetCount(0, username)
		db.Close()
		if err != nil {
			res.Body.Code = constant.RCDatabaseError
			res.Body.Msg = fmt.Sprintf("%v", err.Error())
			res.Reply(w)
			return
		}
		_, db, err = o.UpdateLastLoginAndStatus(username, currentTime.String(), 1)
		db.Close()
		if err != nil {
			res.Body.Code = constant.RCDatabaseError
			res.Body.Msg = fmt.Sprintf("%v", err.Error())
			res.Reply(w)
			return
		}

		combineSessionId := strings.ToUpper(username) + currentTime.String()
		hashSessionId := []byte(combineSessionId)
		hash_byteSessionId := sha256.Sum256(hashSessionId)
		hash_strSessionId := hex.EncodeToString(hash_byteSessionId[:])

		layoutFormat := "2006-01-02 15:04:05"
		strtime, _ := time.Parse(layoutFormat, login.LastLogin)
		sessionMaxTime := strtime.Add(time.Minute * 30)
		if currentTime.String() > sessionMaxTime.String() {
			_, db, err = o.UpdateSessionId(username, hash_strSessionId)
			db.Close()
			if err != nil {
				res.Body.Code = constant.RCDatabaseError
				res.Body.Msg = fmt.Sprintf("%v", err.Error())
				res.Reply(w)
				return
			}
		}
		claims := &Claims{
			Username: login.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: sessionMaxTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(helpers.JwtKey())
		if err != nil {
			res.Body.Code = constant.GeneralErrorCode
			res.Body.Msg = "test"
			res.Reply(w)
			return
		}
		http.SetCookie(w,
			&http.Cookie{
				Name:    "token",
				Path:    "/",
				Value:   tokenString,
				Expires: sessionMaxTime,
			})
		fmt.Println(tokenString)
	}
	resData := LoginStructRes{}
	resData.Username = login.Username
	resData.Email = login.Email
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)
	return

}
