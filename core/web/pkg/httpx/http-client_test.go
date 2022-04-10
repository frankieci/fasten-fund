package httpx

import (
	"log"
	"net/url"
	"testing"
)

/**
 * Created by frankieci on 2022/3/23 10:27 pm
 */

func TestClientDo(t *testing.T) {
	method := "GET"
	urlPath := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "v1/custom/test",
	}

	paramReq := &map[string]interface{}{}
	bearer := ""
	headers := map[string]interface{}{
		"Authorization": bearer,
	}

	log.Println(paramReq)

	resp, err := ClientDo(method, urlPath.String(), paramReq, headers, nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(resp))
}
