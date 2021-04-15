package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type LoginRequestData struct {
	BatchUUID string `json:"batchUuid"`
	Channel string `json:"channel"`
	FBAvatarUrl string `json:"fbAvatarUrl"`
	FBName string `json:"fbName"`
	FBUserID string `json:"fbUserId"`
	Token string `json:"token"`
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
	loginRequest := &LoginRequestData{
		BatchUUID: "1df3663e",
		Channel: "",
		FBAvatarUrl: "https://www.baidu.com",
		FBName: "st-heart",
		Token: "salkhglkasdnflkjdsahng",
	}
	loginRequest.FBUserID = strconv.FormatInt(globalCounter, 10)
	globalCounter += 1
	loginBytes, _ := json.Marshal(loginRequest)
	r2.Body = ioutil.NopCloser(bytes.NewReader(loginBytes))
	// 重置ContentLength，否则会直接被拒绝
	r2.ContentLength = int64(len(loginBytes))
	return r2
}

// exported as symbol named "Adapter"
var Adapter = RequestFunc("Login For OpenPO")
