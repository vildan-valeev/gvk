package gvk

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func check(r APIResponse) error {
	if b := r.Base(); !b.Ok {
		return &APIError{code: b.ErrorCode, desc: b.Description}
	}
	return nil
}

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
