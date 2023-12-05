package gvk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func get[T APIResponse](base, endpoint string, vals url.Values) (res T, err error) {
	url, err := url.JoinPath(base, endpoint)
	if err != nil {
		return res, err
	}

	if vals != nil {
		if queries := vals.Encode(); queries != "" {
			url = fmt.Sprintf("%s?%s", url, queries)
		}
	}

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}
	err = check(res)
	return
}

// sendGetRequest is used to send an HTTP GET request.
func sendGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func check(r APIResponse) error {
	if b := r.Base(); b.Failed == 1 {
		return &APIError{code: b.ErrorCode, desc: b.Description}
	}
	return nil
}
