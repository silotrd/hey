package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type DisbursementRequestData struct {
	AccountHolderName string `json:"accountHolderName"`
	AccountNumber     string `json:"accountNumber"`
	Amount	          int64  `json:"amount"`
	BankCode          string `json:"bankCode"`
	Description       string `json:"description"`
	ExternalID        string `json:"externalId"`
	IdempotencyKey    string `json:"idempotencyKey"`
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
		AccountHolderName: "miaolinjie",
		AccountNumber:     "123123123123",
		Amount:            100010,
		BankCode:          "002",
		Description:       "test description",
	}
	disbursementRequest.ExternalID = strconv.FormatInt(globalCounter, 10)
	disbursementRequest.IdempotencyKey = disbursementRequest.ExternalID
	globalCounter += 1
	disbursementBytes, _ := json.Marshal(disbursementRequest)
	r2.Body = ioutil.NopCloser(bytes.NewReader(disbursementBytes))
	// 重置ContentLength，否则会直接被拒绝
	r2.ContentLength = int64(len(disbursementBytes))
	return r2
}

// exported as symbol named "Adapter"
var Adapter = RequestFunc("Xendit Adapter For Disbursement")
