package httpx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

/**
 * Created by frankieci on 2022/3/16 17:34
 */

var client = http.Client{Transport: http.DefaultTransport}

func ClientDo(method, urlPath string, req interface{}, headers map[string]interface{}, body io.Reader) ([]byte, error) {
	httpReq, err := http.NewRequest(method, urlPath, body)
	if err != nil {
		return nil, err
	}

	rawQuery, err := reflectReq(req)
	if err != nil {
		return nil, err
	}
	httpReq.URL.RawQuery = rawQuery

	for k, v := range headers {
		httpReq.Header.Add(k, fmt.Sprintf("%v", v))
	}

	response, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func reflectReq(v interface{}) (string, error) {
	var empty string
	if v == nil {
		return empty, nil
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
