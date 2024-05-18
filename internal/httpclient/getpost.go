package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func PostV2(path string, postData []byte, header map[string]string, httpClient *http.Client) ([]byte, error) {
	var s []byte
	if httpClient == nil {
		httpClient = defaultHttpClient
	}
	payload := bytes.NewReader(postData)
	req, err := http.NewRequest("POST", path, payload)
	if err != nil {
		return s, err
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return s, err
	}
	if res != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return s, err
	}
	return body, nil
}

func POSTBase(path string, reqStr string, header map[string]string) ([]byte, error) {

	payload := strings.NewReader(reqStr)
	req, _ := http.NewRequest("POST", path, payload)
	req.Header.Add("content-type", "application/json")

	for key, value := range header {
		req.Header.Add(key, value)
	}
	/*
	   req.Header.Add("isupdate", isupdate)
	   req.Header.Add("Api-Key", apiKey)
	   req.Header.Add("accept", "application/json")
	*/

	resp, err := defaultHttpClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respBytes, nil
}

func PostData(path string, data []byte, header map[string]string, cli *http.Client) ([]byte, error) {
	if cli == nil {
		cli = newClient(true)
	}

	payload := bytes.NewReader(data)
	req, err := http.NewRequest("POST", path, payload)
	if err != nil {
		return nil, err
	}

	//req.Header.Add("content-type", "application/json")

	for key, value := range header {
		req.Header.Add(key, value)
	}

	rsp, err := cli.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return []byte(""), err
	}

	return body, nil
}

func GetV2(fullUrlPath string, header map[string]string, httpClient *http.Client) ([]byte, error) {
	var s []byte
	if httpClient == nil {
		httpClient = defaultHttpClient
	}
	req, err := http.NewRequest("GET", fullUrlPath, nil)
	if err != nil {
		return s, err
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return s, err
	}
	if res != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return s, err
	}
	return body, nil
}

func GetV1(fullUrlPath string, header map[string]string, httpClient *http.Client) (*http.Response, error) {
	if httpClient == nil {
		httpClient = defaultHttpClient
	}
	req, err := http.NewRequest("GET", fullUrlPath, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func PutV2(path string, postData []byte, header map[string]string, httpClient *http.Client) ([]byte, error) {
	var s []byte
	if httpClient == nil {
		httpClient = defaultHttpClient
	}
	payload := bytes.NewReader(postData)
	req, err := http.NewRequest("PUT", path, payload)
	if err != nil {
		return s, err
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return s, err
	}
	if res != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return s, err
	}
	return body, nil
}

func POSTHTTPFORM(url string, data url.Values, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return []byte(""), err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := defaultHttpClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respBytes, nil
}

func PostWithAuth(path string, postData []byte, header map[string]string, username, password string, httpClient *http.Client) ([]byte, error) {
	var s []byte
	if httpClient == nil {
		httpClient = defaultHttpClient
	}
	payload := bytes.NewReader(postData)
	req, err := http.NewRequest("POST", path, payload)
	if err != nil {
		return s, err
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}
	req.SetBasicAuth(username, password)
	res, err := httpClient.Do(req)
	if err != nil {
		return s, err
	}
	if res != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return s, err
	}
	return body, nil
}

func PostFormByClient(url string, data string, header map[string]string, httpClient *http.Client) ([]byte, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return []byte(""), err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respBytes, nil
}
