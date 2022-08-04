package helpers

import (
	"encoding/json"
	"net/http"
	"time"
)

// Response is a Object Response Helpers
type Response struct {
	Header headerRes
	Body   bodyRes
}

type headerRes struct {
	RefNum string
}

type bodyRes struct {
	Code string      `json:"response_code"`
	Msg  string      `json:"response_msg"`
	Data interface{} `json:"response_data"`
}

type Log struct {
	URILog      string
	ProcessTime int64
	RefNum      string
	PayloadLog  interface{}
	HeaderLog   interface{}
	StatusLog   string
	MicroName   string
	ConnectTo   string
	ConnectFrom string
	Phase       string
	Platform    string
}

// ResJSON is method for json response
func (o Response) ResJSON(res http.ResponseWriter, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Content-Security-Policy", "default-src 'self'")
	_ = json.NewEncoder(res).Encode(data)
}
func (r *Response) Reply(res http.ResponseWriter) {
	var logRes Log
	var timez time.Time

	logRes.Phase = "Response"
	logRes.ConnectTo = "self"

	js, _ := json.Marshal(r.Body)

	logRes.PayloadLog = r.Body.Data

	// write to log
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("X-Reference-Number", r.Header.RefNum)
	logRes.HeaderLog = res.Header()
	processTime := time.Since(timez).Milliseconds()
	logRes.ProcessTime = processTime
	res.Write(js)
}

// Data is a

// Data is a
type Data struct {
	ReferenceNum string      `json:"reference_num"`
	Data         interface{} `json:"data"`
}

// OneRes is single response a api
type OneRes struct {
	ResponseCode string `json:"response_code"`
	ResponseMsg  string `json:"response_msg"`
	ResponseData *Data  `json:"response_data"`
}

// Datas is a
type Datas struct {
	ReferenceNum string      `json:"reference_num"`
	Data         interface{} `json:"data"`
	Pagination   *Pagin      `json:"pagination"`
}

// Pagin is a
type Pagin struct {
	Limit int `json:"limit"`
	Start int `json:"start"`
}

// ManyRes is multiple response a api
type ManyRes struct {
	ResponseCode string `json:"response_code"`
	ResponseMsg  string `json:"response_msg"`
	ResponseData *Datas `json:"response_data"`
}
