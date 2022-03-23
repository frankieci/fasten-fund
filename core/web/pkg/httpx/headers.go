package httpx

import (
	"github.com/gin-gonic/gin"
)

/**
 * Created by frankieci on 2022/3/15 9:44
 */

type (
	HttpHeaderKey = string
	HttpHeaderVal = string
)

const (
	ContentDisposition      HttpHeaderKey = "Content-Disposition"
	ContentType             HttpHeaderKey = "Content-Type"
	ContentTransferEncoding HttpHeaderKey = "Content-Transfer-Encoding"
	ResponseType            HttpHeaderKey = "response-type"
)

const (
	OctetStream HttpHeaderVal = "application/octet-stream"
	Binary      HttpHeaderVal = "binary"
	Excel       HttpHeaderVal = "application/vnd.ms-excel"
	Blob        HttpHeaderVal = "blob"
)

// SetHttpHeaders specifies http headers for user-defined
func SetHttpHeaders(w gin.ResponseWriter, headers map[string]string) {
	for key, val := range headers {
		w.Header().Add(key, val)
	}
}

func GetFileHeaders(filename string) map[string]string {
	return map[string]string{
		ContentDisposition:      "attachment; filename=" + filename,
		ContentType:             OctetStream,
		ContentTransferEncoding: Binary,
		ResponseType:            Blob,
	}
}
