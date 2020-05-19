// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package httputil

import (
	"encoding/base64"
	"io"
	"net"
	"net/http"
	"time"
)

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func HttpGetWithCredential(url, credential string) (*http.Response, error) {
	var netTransport = &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
		Dial: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 30,
		Transport: netTransport,
	}
	//response, err := netClient.Get(url)
	response, err := func(credential string) (*http.Response, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		if credential != "" {
			req.Header.Set("Authorization", "Basic "+credential)
		}
		return netClient.Do(req)
	}(credential)

	if err != nil {
		return response, err
	}
	return response, err
}

func HttpGet(url string) (*http.Response, error) {
	var netTransport = &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
		Dial: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 30,
		Transport: netTransport,
	}
	response, err := netClient.Get(url)
	if err != nil {
		return response, err
	}
	return response, err
}

func HttpPost(url, contentType string, body io.Reader) (*http.Response, error) {
	var netTransport = &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
		Dial: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 30,
		Transport: netTransport,
	}

	response, err := netClient.Post(url, contentType, body)
	if err != nil {
		return response, err
	}
	return response, err
}
