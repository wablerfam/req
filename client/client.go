package client

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

func Request(url string) (int, string, time.Duration) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: transport,
	}

	stime := time.Now()

	req, err := http.NewRequest("GET", url, nil)
  	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	etime := time.Now().Sub(stime)

	code := int(res.StatusCode)

	byteBody, _ := ioutil.ReadAll(res.Body)
	body := string(byteBody)

	return code, body, etime
}
