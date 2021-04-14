package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type DisbursementRequestData struct {
	BizMerchantID int64 `json:"bizMerchantId"`
	DisbursementID int64 `json:"disbursementId"`
	AccountHolderName string `json:"accountHolderName"`
	AccountNumber string `json:"accountNumber"`
	BankCode string `json:"bankCode"`
	Description string `json:"description"`
	Currency string `json:"currency"`
	DisbursementAmount float64 `json:"disbursementAmount"`
}

var (
	globalCounter = time.Now().UnixNano()
)

type RequestFunc string

func (rf RequestFunc) RequestFunc(req *http.Request) *http.Request {
	r2 := new(http.Request)
	*r2 = *req
	r2.Header = make(http.Header, len(req.Header))
	for k,s := range req.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	disbursementRequest := &DisbursementRequestData{
		AccountHolderName: "test",
		AccountNumber:     "1231231231",
		DisbursementAmount: 100010,
		BankCode:          "014",
		Description:       "test description",
		Currency: "IDR",
		BizMerchantID: 1321739445694734338,
	}
	disbursementRequest.DisbursementID = globalCounter
	globalCounter += 1
	disbursementBytes, _ := json.Marshal(disbursementRequest)
	r2.Body = ioutil.NopCloser(bytes.NewReader(disbursementBytes))
	// 重置ContentLength，否则会直接被拒绝
	r2.ContentLength = int64(len(disbursementBytes))
	return r2
}


// exported as symbol named "Adapter"
var Adapter = RequestFunc("PaymentEngine For Disbursement")
