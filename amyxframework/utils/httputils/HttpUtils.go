package httputils

import (
	"io/ioutil"
	"net/http"
)

func Get(url string, headers map[string]string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
