package main

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UsError interface {
	error
	Status() int
}

type ReqError struct {
	Code int
	Err  error
}

func (re ReqError) Error() string {
	return re.Err.Error()
}

func (re ReqError) Status() int {
	return re.Code
}

func NewNotFindErr(err error) ReqError {
	return ReqError{
		Code: http.StatusNotFound,
		Err:  err,
	}
}

func NewBadReqErr(err error) ReqError {
	return ReqError{
		Code: http.StatusBadRequest,
		Err:  err,
	}
}
func NewTimeoutErr() ReqError {
	return ReqError{
		Code: http.StatusRequestTimeout,
		Err:  errors.New(http.StatusText(http.StatusRequestTimeout)),
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func responseWithError(w http.ResponseWriter, err error, payload interface{}) {
	switch e := err.(type) {
	case UsError:
		logrus.Warnf("http %d - %s", e.Status(), e)
		resp, _ := json.Marshal(Response{Code: e.Status(), Message: e.Error(), Content: payload})
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(resp)
	default:
		logrus.Error(err)
		responseWithJson(w, http.StatusInternalServerError, payload)
	}
}

func responseWithJson(w http.ResponseWriter, status int, payload interface{}) {
	msg := http.StatusText(status)
	if status == http.StatusOK {
		status = 1
	}
	resp, _ := json.Marshal(Response{Code: status, Message: msg, Content: payload})
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resp)
}
