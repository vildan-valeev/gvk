package gvk

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func get[T Response](base, endpoint string, vals url.Values) (res T, err error) {
	url, err := url.JoinPath(base, endpoint)
	if err != nil {
		return res, err
	}

	if vals != nil {
		if queries := vals.Encode(); queries != "" {
			url = fmt.Sprintf("%s?%s", url, queries)
		}
	}
	log.Printf("Request URL: %s", url)
	cnt, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	return
}

func getUpdates(base, endpoint string, vals url.Values) (res ResponseUpdate, err error) {
	url, err := url.JoinPath(base, endpoint)
	if err != nil {
		return res, err
	}

	if vals != nil {
		if queries := vals.Encode(); queries != "" {
			url = fmt.Sprintf("%s?%s", url, queries)
		}
	}

	cnt, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	return
}

// sendGetRequest is used to send an HTTP GET request.
func SendGetRequest(url string) ([]byte, error) {
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

func check(r APIError) error {
	//switch r.failed {
	//case 0:
	//	d.Ts = r.Ts
	//case 1:
	//	d.Ts = r.Ts
	//case 2:
	//	err = d.updateServer(false)
	//case 3:
	//	err = d.updateServer(true)
	//default:
	//	err = &LongPoolError{failed: r.failed}
	//}
	//return nil
	e := r.Base()
	log.Println(e.Code, e.Message)
	return nil
}
