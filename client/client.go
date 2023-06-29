package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var nilBinding = Binding{}

type PartyMQClient interface {
	Bind(hostname string) (Binding, error)
	Unbind(hostname string) error
	Ready(hostname string) error
	HeartBeat(hostname string) error
}

type srvCtx struct {
	httpClient *http.Client
	baseUrl    string
}

func New(httpClient *http.Client, baseUrl string) PartyMQClient {
	return &srvCtx{httpClient: httpClient, baseUrl: baseUrl}
}

func (sCtx *srvCtx) Bind(hostname string) (Binding, error) {
	resp, err := sCtx.httpClient.Post(fmt.Sprintf("%s/%s/bind", sCtx.baseUrl, hostname), "json/application", nil)
	if err != nil {
		return nilBinding, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var binding Binding
	err = json.Unmarshal(bodyBytes, &binding)
	if err != nil {
		return nilBinding, err
	}

	return binding, nil
}

func (sCtx *srvCtx) Unbind(hostname string) error {
	resp, err := sCtx.httpClient.Post(fmt.Sprintf("%s/%s/unbind", sCtx.baseUrl, hostname), "json/application", nil)
	if err != nil {
		return err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		return fmt.Errorf("recieved non sucessfull status code: %s", resp.StatusCode)
	}

	return nil
}

func (sCtx *srvCtx) Ready(hostname string) error {
	resp, err := sCtx.httpClient.Post(fmt.Sprintf("%s/%s/ready", sCtx.baseUrl, hostname), "json/application", nil)
	if err != nil {
		return err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		return fmt.Errorf("recieved non sucessfull status code: %s", resp.StatusCode)
	}

	return nil
}

func (sCtx *srvCtx) HeartBeat(hostname string) error {
	resp, err := sCtx.httpClient.Post(fmt.Sprintf("%s/%s/heartbeat", sCtx.baseUrl, hostname), "json/application", nil)
	if err != nil {
		return err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		return fmt.Errorf("recieved non sucessfull status code: %s", resp.StatusCode)
	}

	return nil
}
