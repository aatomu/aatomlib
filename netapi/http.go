package netapi

import (
	"bytes"
	"io"
	"net/http"
)

func GetRequest(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	return
}

type ReqestMethod string

const (
	Post ReqestMethod = "POST"
	Get  ReqestMethod = "GET"
)

// 複数headerを送る際は map["A"] = "a;b;c"
func Request(method ReqestMethod, uri string, body []byte, headers map[string]string) (resp *http.Response, err error) {
	// リクエストの準備
	req, _ := http.NewRequest(string(method), uri, bytes.NewBuffer(body))
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	// 送信クライアント準備
	client := new(http.Client)
	// Request送信
	resp, err = client.Do(req)
	return
}
