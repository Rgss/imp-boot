package crequest

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func SimplePost(link string, data []byte) (string, error) {
	return Request("POST", link, data, map[string]interface{}{})
}

func SimpleGet(link string, data []byte) (string, error) {
	return Request("POST", link, data, map[string]interface{}{})
}

func Post(link string, data []byte, header map[string] interface{}) (r string, err error) {
	return Request("POST", link, data, header)
}

func Get(link string, data []byte, header map[string] interface{}) (r string, err error) {
	return Request("GET", link, data, header)
}

func Request(method string, link string, data []byte, header map[string] interface{}) (r string, err error) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequest(method, link, strings.NewReader(string(data)))
	if err != nil {
		log.Println("err: ", err)
		return "", err
	}

	// header
	for key := range header {
		req.Header.Set(key, header[key].(string))
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(body), err
}