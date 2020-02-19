package mango_request

import "net/http"

type Request struct {
	ApiKey string `json:"vpbx_api_key"`
	Sign   string `json:"sign"`
	Json   string `json:"json"`
}

func ParseRequest(r *http.Request) *Request {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	return &Request{ApiKey: r.PostFormValue("vpbx_api_key"), Sign: r.PostFormValue("sign"), Json: r.PostFormValue("json")}
}
