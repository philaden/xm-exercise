package helpers

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	IRestClient interface {
		Get(endPointUrl string, headers map[string]string) RestResponse
		Post(endPointUrl string, headers map[string]string, payLoad interface{}) RestResponse
		Patch(endPointUrl string, headers map[string]string, payLoad interface{}) RestResponse
		Put(endPointUrl string, headers map[string]string, payLoad interface{}) RestResponse
	}

	RestClient struct {
		HttpClient *http.Client
	}
)

func (client RestClient) Get(endPointUrl string, headers map[string]string) RestResponse {

	req, err := http.NewRequest(http.MethodGet, endPointUrl, nil)
	for k := range req.Header {
		delete(req.Header, k)
	}

	if headers != nil {
		addHeaders(req, headers)
	}

	resp, err := client.HttpClient.Do(req)

	if err != nil {
		fmt.Printf("Unsuccessful response from %s", endPointUrl)
		headrs, _ := json.Marshal(req.Header)
		fmt.Printf("Headers sent %s", string(headrs))
		return BaseRestResponse(false, resp.StatusCode, "", err)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Unsuccessful response from %s", endPointUrl)
		headrs, _ := json.Marshal(req.Header)
		fmt.Printf("Headers sent %s", string(headrs))
		return BaseRestResponse(false, resp.StatusCode, "", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("unable to read http response data")
		return BaseRestResponse(false, 400, "", err)
	}

	return BaseRestResponse(true, resp.StatusCode, string(body), nil)
}

func (client RestClient) Post(endPointUrl string, headers map[string]string, payLoad interface{}) RestResponse {

	var network bytes.Buffer
	encoder := gob.NewEncoder(&network)
	err := encoder.Encode(payLoad)

	if err != nil {
		return BaseRestResponse(false, http.StatusBadRequest, "", err)
	}

	postData := bytes.NewBuffer(network.Bytes())

	req, err := http.NewRequest(http.MethodPost, endPointUrl, postData)
	for k := range req.Header {
		delete(req.Header, k)
	}

	if headers != nil {
		addHeaders(req, headers)
	}

	resp, err := client.HttpClient.Do(req)

	if err != nil {
		fmt.Printf("Unsuccessful response from %s", endPointUrl)
		headrs, _ := json.Marshal(req.Header)
		fmt.Printf("Headers sent %s", string(headrs))
		return BaseRestResponse(false, resp.StatusCode, "", err)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Unsuccessful response from %s", endPointUrl)
		headrs, _ := json.Marshal(req.Header)
		fmt.Printf("Headers sent %s", string(headrs))
		return BaseRestResponse(false, resp.StatusCode, "", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("unable to read http response data")
		return BaseRestResponse(false, 400, "", err)
	}

	return BaseRestResponse(true, resp.StatusCode, string(body), nil)
}

func (client RestClient) Patch(endPointUrl string, headers map[string]string, payLoad interface{}) RestResponse {

	var network bytes.Buffer
	encoder := gob.NewEncoder(&network)
	err := encoder.Encode(payLoad)

	if err != nil {
		return BaseRestResponse(false, http.StatusBadRequest, "", err)
	}

	postData := bytes.NewBuffer(network.Bytes())

	req, err := http.NewRequest(http.MethodPatch, endPointUrl, postData)
	for k := range req.Header {
		delete(req.Header, k)
	}

	if headers != nil {
		addHeaders(req, headers)
	}

	resp, err := client.HttpClient.Do(req)

	if err != nil {
		fmt.Printf("Unsuccessful response from %s", endPointUrl)
		headrs, _ := json.Marshal(req.Header)
		fmt.Printf("Headers sent %s", string(headrs))
		return BaseRestResponse(false, resp.StatusCode, "", err)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Unsuccessful response from %s", endPointUrl)
		headrs, _ := json.Marshal(req.Header)
		fmt.Printf("Headers sent %s", string(headrs))
		return BaseRestResponse(false, resp.StatusCode, "", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("unable to read http response data")
		return BaseRestResponse(false, 400, "", err)
	}

	return BaseRestResponse(true, resp.StatusCode, string(body), nil)
}

func (client RestClient) Put(endPointUrl string, headers map[string]string, payLoad interface{}) RestResponse {

	var network bytes.Buffer
	encoder := gob.NewEncoder(&network)
	err := encoder.Encode(payLoad)

	if err != nil {
		return BaseRestResponse(false, http.StatusBadRequest, "", err)
	}

	postData := bytes.NewBuffer(network.Bytes())

	req, err := http.NewRequest(http.MethodPut, endPointUrl, postData)
	for k := range req.Header {
		delete(req.Header, k)
	}

	if headers != nil {
		addHeaders(req, headers)
	}

	resp, err := client.HttpClient.Do(req)

	if err != nil {
		fmt.Printf("Unsuccessful response from %s", endPointUrl)
		headrs, _ := json.Marshal(req.Header)
		fmt.Printf("Headers sent %s", string(headrs))
		return BaseRestResponse(false, resp.StatusCode, "", err)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Unsuccessful response from %s", endPointUrl)
		headrs, _ := json.Marshal(req.Header)
		fmt.Printf("Headers sent %s", string(headrs))
		return BaseRestResponse(false, resp.StatusCode, "", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("unable to read http response data")
		return BaseRestResponse(false, 400, "", err)
	}

	return BaseRestResponse(true, resp.StatusCode, string(body), nil)
}

func addHeaders(request *http.Request, headers map[string]string) {
	headers["Content-Type"] = "application/json"
	for key, value := range headers {

		request.Header.Add(key, value)
	}
}
