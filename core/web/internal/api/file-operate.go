package api

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/fasten-fund/core"

	"github.com/fasten-fund/core/web/pkg/httpx"
	"github.com/gin-gonic/gin"
)

type uploadFileReq struct {
	Uuid string `form:"uuid"`
}

// UploadFile  上传文件接口
// @Summary  上传文件
// @Description  上传文件接口
// @Tags  文件操作
// @Produce  application/json
// @Produce  plain
// @Param uuid       query string false "uuid"
// @Param  Authorization  header  string  true  "Authentication header"
// @Param file   formData file true "upload file"
// @Success  200  {object}  resp{}  "code: 0"
// @Failure  400  {object}  resp{}
// @Failure  404  {object}  resp{}
// @Failure  500  {object}  resp{}
// @Router   /v1/custom/file/upload [post]
func UploadFile(ctx *gin.Context) {
	var req uploadFileReq
	if err := ctx.BindQuery(&req); err != nil {
		return
	}

	filePath, err := fileHandler(ctx, req.Uuid)
	if err != nil {
		return
	}

	// record
	_ = filePath
}

func fileHandler(ctx *gin.Context, uuid string) (string, error) {
	var filePath string
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return filePath, err
	}

	savePath := filepath.Join("/data/files")
	if ok, _ := core.PathExists(savePath); !ok {
		if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
			return filePath, err
		}
	}

	filePath = filepath.Join(savePath, fmt.Sprintf("file-record-%s.txt", uuid))
	if err := ctx.SaveUploadedFile(fileHeader, filePath); err != nil {
		return filePath, err
	}

	return filePath, nil
}

type fileDownloadReq struct {
	FromTime int64 `form:"from,omitempty,string" desc:"起始时间"`
	EndTime  int64 `form:"end,omitempty,string" desc:"结束时间"`
}

// FileDownload 文件下载接口
// @Summary  文件下载
// @Description 文件下载接口
// @Tags 文件操作
// @Produce application/json
// @Produce      plain
// @Param from  query string false "起始时间"
// @Param end   query string false "结束时间"
// @Param        Authorization  header    string  true  "Authentication header"
// @Success      200 {object}   resp{data=}    "code: 0"
// @Failure      400  {object}  resp{}
// @Failure      404  {object}  resp{}
// @Failure      500  {object}  resp{}
// @Router  /v1/custom/file/download [get]
func FileDownload(ctx *gin.Context) {
	var req fileDownloadReq
	if err := ctx.BindQuery(&req); err != nil {
		return
	}

	data := make([][]interface{}, 0)
	// Assemble data
	filename := fmt.Sprintf("file-record-%d.xlsx", time.Now().Unix())
	httpx.SetHttpHeaders(ctx.Writer, httpx.GetFileHeaders(filename))

	sheetHeader := []string{"A1"}
	f := core.WriteExcelFile("Sheet1", sheetHeader, data)

	buf, err := f.WriteToBuffer()
	if err != nil {
		return
	}

	http.ServeContent(ctx.Writer, ctx.Request, filename, time.Now(), bytes.NewReader(buf.Bytes()))
}
