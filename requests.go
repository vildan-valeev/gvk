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
	//log.Printf("Request URL: %s", url)
	cnt, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	//log.Printf("Response Error: %v", err)
	//log.Printf("Response BODY: %s", string(cnt))
	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	err = check(res)

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

func check(r Response) error {
	e := r.Base()
	switch e.Code {
	case 0:
		return nil
	// TODO: add another codes https://vk.com/dev/errors
	default:
		log.Println(e.Error())
		return e
	}
}
