package utils

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"net/http"
	"errors"
)

func GetWithCookie(input *url.URL, cookie *http.Cookie, headers map[string]string) ([]byte, error) {
	fmt.Println("loading url " + input.String() + " with cookie " + cookie.String())
	client := http.Client{}
	req, err := http.NewRequest("GET", input.String(), nil)
	if err != nil {
		return []byte{}, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	if resp.StatusCode != 200 {
		return []byte{}, errors.New(fmt.Sprintf("Server response with invalid status code %d", resp.StatusCode))
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)

}
