package curl2

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

func PostJson[T1 any, T2 any](apiUrl string, in T1, headers1 map[string]string, proxy1 string) (*T2, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	for k, v := range headers1 {
		if k == "Content-Type" {
			continue
		}
		headers[k] = v
	}
	rsp, err := postX(apiUrl, in, headers, proxy1)
	if err != nil {
		return nil, err
	}
	var out T2
	err1 := json.Unmarshal(rsp, &out)
	if err1 != nil {
		return nil, err1
	}
	return &out, nil
}

func postX(gateway string, data interface{}, headers map[string]string, proxy1 string) ([]byte, error) {
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	// 設置自定義的 http.Transport，禁用證書驗證
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 禁用證書驗證
		},
	}
	if proxy1 != "" {
		customTransport.Proxy = http.ProxyURL(func() *url.URL { url1, _ := url.Parse(proxy1); return url1 }())
	}
	cli1 := &http.Client{Transport: customTransport}
	req1, err := http.NewRequest(http.MethodPost, gateway, bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			req1.Header.Add(k, v)
		}
	}
	res, err := cli1.Do(req1)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			logrus.WithError(err1).Error("fail to close body")
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
