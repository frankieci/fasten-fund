package httpx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"
)

/**
 * Created by frankieci on 2022/3/16 17:34
 */

func ClientDo(req interface{}, httpReq *http.Request, headers map[string]interface{}) ([]byte, error) {
	if strings.ToUpper(httpReq.Method) == "GET" {
		rawQuery, err := reflectReq(req)
		if err != nil {
			return nil, err
		}
		httpReq.URL.RawQuery = rawQuery
	} else {
		b, err := json.Marshal(&req)
		if err != nil {
			return nil, err
		}
		httpReq, err = http.NewRequest(httpReq.Method, httpReq.URL.String(), bytes.NewReader(b))
	}

	for k, v := range headers {
		httpReq.Header.Add(k, fmt.Sprintf("%v", v))
	}

	log.Println(httpReq.URL.String())

	client := http.Client{Transport: http.DefaultTransport}
	response, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func reflectReq(v interface{}) (string, error) {
	var empty string
	if v == nil {
		return empty, fmt.Errorf("%w: %+v", errors.New("unsupported data type"), v)
	}

	reflectValueType := reflect.ValueOf(v).Type()
	isPtr := reflectValueType.Kind() == reflect.Ptr
	if !isPtr {
		return empty, errors.New("reflecting the dest is not pointer type")
	}

	for reflectValueType.Kind() == reflect.Ptr {
		reflectValueType = reflectValueType.Elem()
	}

	if reflectValueType.Kind() == reflect.Struct || reflectValueType.Kind() == reflect.Map {
		bm, err := json.Marshal(&v)
		if err != nil {
			return empty, err
		}

		var mr map[string]interface{}
		if err := json.Unmarshal(bm, &mr); err != nil {
			return empty, err
		}

		values := make(url.Values)
		for k, val := range mr {
			values.Set(k, fmt.Sprintf("%v", val))
		}
		return values.Encode(), nil
	}
	return empty, nil
}
