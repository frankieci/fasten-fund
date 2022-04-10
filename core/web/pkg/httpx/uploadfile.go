package httpx

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"os"
)

/**
 * Created by frankieci on 2022/4/10 10:08 am
 */

func UploadFile(filePath string, urlPath url.URL, req interface{}) ([]byte, error) {
	method := "POST"

	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	fileWriter, _ := bodyWriter.CreateFormFile("file", "file.txt")

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if _, err := io.Copy(fileWriter, file); err != nil {
		return nil, err
	}

	headers := map[string]interface{}{"Content-Type": bodyWriter.FormDataContentType()}

	if err := bodyWriter.Close(); err != nil {
		return nil, err
	}

	resp, err := ClientDo(method, urlPath.String(), req, headers, bodyBuffer)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
