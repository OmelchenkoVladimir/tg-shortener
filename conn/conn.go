package conn

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Url_response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"url"`
}

func GetShortUrl(u string, server string) (string, error) {
	// Maybe do better error reports after
	var resp Url_response
	params := url.Values{}
	params.Add("url", u)
	res, err := http.PostForm(server, params)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	log.Println("-->", string(body))
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", err
	}
	if resp.Success != true {
		return "", errors.New("Got [success = false] from server")
	}
	return resp.Data.(string), nil
}
