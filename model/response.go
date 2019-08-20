package model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response
}

func NewResponse(response *http.Response) *Response {
	return &Response{
		response,
	}
}

func (r *Response) DecodeBody(body interface{}) error {
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	err = json.Unmarshal(raw, body)

	return err
}