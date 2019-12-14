package shared

import (
	"bytes"
	"encoding/json"
	"github.com/Bob-MusicPlayer/shared-bob/model"
	"net/http"
	"net/url"
	"strings"
)

type RequestHelper struct {
	baseUrl string
}

func NewRequestHelper(baseUrl string) *RequestHelper {
	return &RequestHelper{
		baseUrl: baseUrl,
	}
}

func (rh *RequestHelper) buildUrl(endpoint string) (string, error) {
	rh.baseUrl = strings.TrimPrefix(rh.baseUrl, "/")
	fullUrl, err := url.Parse(rh.baseUrl + endpoint)
	if err != nil {
		return "", err
	}

	return fullUrl.String(), nil
}

func (rh *RequestHelper) Get(endpoint string, header http.Header) (*model.Response, error) {
	uri, err := rh.buildUrl(endpoint)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	if header != nil {
		request.Header = header
	}

	response, err := http.DefaultClient.Do(request)

	return model.NewResponse(response), err
}

func (rh *RequestHelper) Post(endpoint string, payload interface{}, header http.Header) (*model.Response, error) {
	uri, err := rh.buildUrl(endpoint)
	if err != nil {
		return nil, err
	}

	var data []byte

	if payload != nil {
		data, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	if header != nil {
		request.Header = header
	}

	response, err := http.DefaultClient.Do(request)

	return model.NewResponse(response), err
}
