package mango_request

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func RequestToMango(jsonMessage, url, apiKey, apiSing, apiUrl string) string {

	var sign = fmt.Sprintf("%x", sha256.Sum256([]byte(apiKey+jsonMessage+apiSing)))

	request := []byte("vpbx_api_key=" + apiKey + "&sign=" + sign + "&json=" + jsonMessage)

	resp, err := http.Post(apiUrl+url, "application/x-www-form-urlencoded", bytes.NewBuffer(request))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}
