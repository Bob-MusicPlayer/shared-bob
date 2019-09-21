package shared

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"shared-bob/helper"
	"strings"
)

type ResponseHelper struct {
	w http.ResponseWriter
	req *http.Request
}

func NewResponseHelper(w http.ResponseWriter, req *http.Request) *ResponseHelper {
	return &ResponseHelper{
		w:   w,
		req: req,
	}
}

func (rh *ResponseHelper) ReturnOptionsOrNotAllowed(methods ...string) bool {
	if !helper.StringInSlice(rh.req.Method, methods) && rh.req.Method != http.MethodOptions {
		rh.w.WriteHeader(http.StatusMethodNotAllowed)
		return true
	}

	rh.w.Header().Set("Access-Control-Allow-Origin", "*")
	rh.w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ", "))
	rh.w.Header().Set("Content-Type","application/json; charset=utf-8")

	if rh.req.Method == http.MethodOptions {
		rh.w.WriteHeader(http.StatusNoContent)
		return true
	}

	return false
}

func (rh *ResponseHelper) ReturnHasError(err error) bool {
	if err != nil {
		rh.w.WriteHeader(901)
		_, err := rh.w.Write([]byte(err.Error()))
		if rh.ReturnHasError(err) {
			return true
		}
		return true
	}
	return false
}

func (rh *ResponseHelper) ReturnOk(payload interface{}) {
	if payload != nil {
		data, err := json.Marshal(&payload)
		if rh.ReturnHasError(err) {
			return
		}
		_, _ = rh.w.Write(data)
	} else {
		rh.w.WriteHeader(200)
	}
}

func (rh *ResponseHelper) DecodeBody(body interface{}) error {
	raw, err := ioutil.ReadAll(rh.req.Body)
	if err != nil {
		return err
	}

	defer rh.req.Body.Close()

	err = json.Unmarshal(raw, body)

	return err
}

func (rh *ResponseHelper) ReturnError(err error) {
		rh.w.WriteHeader(901)
		_, err = rh.w.Write([]byte(err.Error()))
		if rh.ReturnHasError(err) {}
}
