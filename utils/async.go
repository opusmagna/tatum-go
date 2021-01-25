package utils

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"os"
)

type Async struct {
}

var DefaultClient = &http.Client{}

func (a *Async) SendPut(url string, body []byte) (string, error) {

	var baseUrl = TATUM_API_URL
	if len(baseUrl) == 0 {
		baseUrl = os.Getenv("TATUM_API_URL")
	}
	requestUrl := baseUrl + url

	asyncChan := make(chan *http.Response, 1)

	errGrp, _ := errgroup.WithContext(context.Background())

	errGrp.Go(func() error { return put(requestUrl, body, asyncChan) })

	err := errGrp.Wait()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error with submitting the order, try again later...")
		return "", err
	}

	asyncResponse := <-asyncChan
	defer asyncResponse.Body.Close()
	bytes, _ := ioutil.ReadAll(asyncResponse.Body)

	fmt.Println(string(bytes))

	return string(bytes), nil
}

func put(url string, body []byte, rc chan *http.Response) error {

	req, err := http.NewRequest("PUT", url, bytes.NewReader(body))

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", os.Getenv("TATUM_API_KEY"))

	response, err := DefaultClient.Do(req)
	if err == nil {
		rc <- response
	}
	fmt.Println(response)

	return err
}

func (a *Async) SendDel(url string, body []byte) (string, error) {

	var baseUrl = TATUM_API_URL
	if len(baseUrl) == 0 {
		baseUrl = os.Getenv("TATUM_API_URL")
	}
	requestUrl := baseUrl + url

	asyncChan := make(chan *http.Response, 1)

	errGrp, _ := errgroup.WithContext(context.Background())

	errGrp.Go(func() error { return delete(requestUrl, body, asyncChan) })

	err := errGrp.Wait()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error with submitting the order, try again later...")
		return "", err
	}

	asyncResponse := <-asyncChan
	defer asyncResponse.Body.Close()
	bytes, _ := ioutil.ReadAll(asyncResponse.Body)

	fmt.Println(string(bytes))

	return string(bytes), nil
}

func delete(url string, body []byte, rc chan *http.Response) error {

	req, err := http.NewRequest("DELETE", url, bytes.NewReader(body))

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", os.Getenv("TATUM_API_KEY"))

	response, err := DefaultClient.Do(req)
	if err == nil {
		rc <- response
	}
	fmt.Println(response)

	return err
}

func (a *Async) SendPost(url string, body []byte) (string, error) {

	var baseUrl = TATUM_API_URL
	if len(baseUrl) == 0 {
		baseUrl = os.Getenv("TATUM_API_URL")
	}
	requestUrl := baseUrl + url

	asyncChan := make(chan *http.Response, 1)

	errGrp, _ := errgroup.WithContext(context.Background())

	errGrp.Go(func() error { return post(requestUrl, body, asyncChan) })

	err := errGrp.Wait()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error with submitting the order, try again later...")
		return "", err
	}

	asyncResponse := <-asyncChan
	defer asyncResponse.Body.Close()
	bytes, _ := ioutil.ReadAll(asyncResponse.Body)

	fmt.Println("============================")
	fmt.Println(string(bytes))
	fmt.Println("============================")

	return string(bytes), nil
}

func post(url string, body []byte, rc chan *http.Response) error {
	fmt.Println(string(body))
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", os.Getenv("TATUM_API_KEY"))

	response, err := DefaultClient.Do(req)
	if err == nil {
		rc <- response
	}
	fmt.Println(response)

	return err
}

func (a *Async) SendGet(url string, body []byte) (string, error) {

	var baseUrl = TATUM_API_URL
	if len(baseUrl) == 0 {
		baseUrl = os.Getenv("TATUM_API_URL")
	}
	requestUrl := baseUrl + url

	asyncChan := make(chan *http.Response, 1)

	errGrp, _ := errgroup.WithContext(context.Background())

	errGrp.Go(func() error { return get(requestUrl, body, asyncChan) })

	err := errGrp.Wait()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error with submitting the order, try again later...")
		return "", err
	}

	asyncResponse := <-asyncChan
	defer asyncResponse.Body.Close()
	bytes, _ := ioutil.ReadAll(asyncResponse.Body)

	fmt.Println(string(bytes))

	return string(bytes), nil
}

func get(url string, body []byte, rc chan *http.Response) error {

	req, err := http.NewRequest("GET", url, bytes.NewReader(body))

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", os.Getenv("TATUM_API_KEY"))

	response, err := DefaultClient.Do(req)
	if err == nil {
		rc <- response
	}
	fmt.Println(response)

	return err
}
