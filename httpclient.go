package hitbtc

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type client struct {
	apiKey      string
	apiSecret   string
	httpClient  *http.Client
	httpTimeout time.Duration
}

// NewClient return a new Bittrex HTTP client
func NewClient(apiKey, apiSecret string) (c *client) {
	return &client{apiKey, apiSecret, &http.Client{}, 30 * time.Second}
}

// NewClientWithCustomHttpConfig returns a new Bittrex HTTP client using the predefined http client
func NewClientWithCustomHttpConfig(apiKey, apiSecret string, httpClient *http.Client) (c *client) {
	timeout := httpClient.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	return &client{apiKey, apiSecret, httpClient, timeout}
}

// NewClient returns a new Bittrex HTTP client with custom timeout
func NewClientWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) (c *client) {
	return &client{apiKey, apiSecret, &http.Client{}, timeout}
}

func NewClientHttp(h *http.Client, apiKey, apiSecret string) (c *client) {
	return &client{apiKey, apiSecret, h, 30 * time.Second}
}

// doTimeoutRequest do a HTTP request with timeout
func (c *client) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	// Do the request in the background so we can check the timeout
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := c.httpClient.Do(req)
		done <- result{resp, err}
	}()
	// Wait for the read or the timeout
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("timeout on reading data from Hitbtc API")
	}
}

/*
import datetime
import hashlib
import hmac
import random
import string
import time
import unirest as unirest

key = ""
secret = ""
nonce = str(int(time.mktime(datetime.datetime.now().timetuple()) * 1000 + datetime.datetime.now().microsecond / 1000))

clientOrderId = "".join(random.choice(string.digits + string.ascii_lowercase) for _ in range(30))

path = "/api/1/trading/new_order?apikey=" + key + "&nonce=" + nonce

newOrder = "clientOrderId=" + clientOrderId + "&symbol=BTCUSD&side=buy&price=250&quantity=100&type=limit"

signature = hmac.new(secret, path + newOrder, hashlib.sha512).hexdigest()

result = unirest.post("http://api.hitbtc.com" + path, headers={"Api-Signature": signature}, params=newOrder)

print result.body['ExecutionReport']
*/
func (c *client) do(method string, resource string, payload string, authNeeded bool) (response []byte, err error) {
	connectTimer := time.NewTimer(c.httpTimeout)

	var rawurl string
	if strings.HasPrefix(resource, "http") {
		rawurl = resource
	} else {
		rawurl = fmt.Sprintf("%s%s%s", API_BASE, API_VERSION, resource)
	}
	if authNeeded {
		nonce := time.Now().UnixNano()
		rawurl += fmt.Sprintf("?nonce=%d&apikey=%s", nonce, c.apiKey)
	}
	var req *http.Request
	if method == "POST" {
		req, err = http.NewRequest(method, rawurl, strings.NewReader(payload))
	} else {

		if payload != "" {
			rawurl += fmt.Sprintf("&%s", payload)
		}
		req, err = http.NewRequest(method, rawurl, strings.NewReader(""))
		//rawurl += fmt.Sprintf("?nonce=%d&apikey=%s", nonce, c.apiKey)
	}
	if err != nil {
		return
	}

	//if method == "POST" || method == "PUT" {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//}
	req.Header.Add("Accept", "application/json")

	// Auth
	if authNeeded {

		if len(c.apiKey) == 0 || len(c.apiSecret) == 0 {
			err = errors.New("You need to set API Key and API Secret to call this method")
			return
		}

		q := req.URL.Query()
		req.URL.RawQuery = q.Encode()
		mac := hmac.New(sha512.New, []byte(c.apiSecret))
		if method == "POST" {
			_, err = mac.Write([]byte(req.URL.Path + "?" + req.URL.RawQuery + payload))
		} else {
			_, err = mac.Write([]byte(req.URL.Path + "?" + req.URL.RawQuery))
			fmt.Println(req.URL.Path + "?" + req.URL.RawQuery)
		}
		sig := hex.EncodeToString(mac.Sum(nil))
		req.Header.Add("Api-Signature", sig)
	}
	resp, err := c.doTimeoutRequest(connectTimer, req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	//fmt.Println(fmt.Sprintf("reponse %s", response), err)
	if err != nil {
		return response, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}
	return response, err
}
