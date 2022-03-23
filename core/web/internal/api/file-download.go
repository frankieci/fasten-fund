package api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fasten-fund/core/web/pkg/httpx"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type Req struct {
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
	var req Req
	if err := ctx.BindQuery(&req); err != nil {
		return
	}

	data := make([][]interface{}, 0)
	// Assemble data
	filename := fmt.Sprintf("file-record-%d.xlsx", time.Now().Unix())
	httpx.SetHttpHeaders(ctx.Writer, httpx.GetFileHeaders(filename))

	sheetHeader := []string{"A1"}
	f := writeExcelFile("Sheet1", sheetHeader, data)

	buf, err := f.WriteToBuffer()
	if err != nil {
		return
	}

	http.ServeContent(ctx.Writer, ctx.Request, filename, time.Now(), bytes.NewReader(buf.Bytes()))
}

func writeExcelFile(sheetName string, sheetHeader []string, data [][]interface{}) *excelize.File {
	// The sheet name sets 'Sheet1' for default Excel files
	f := excelize.NewFile()
	sheet := getSheet()
	for k, v := range sheetHeader {
		if err := f.SetCellValue(sheetName, sheet[k]+"1", v); err != nil {
			log.Println("cell value error:", err)
		}
	}

	// Write data in a loop
	for line, d := range data {
		for col := range d {
			if err := f.SetCellValue(sheetName, sheet[col]+strconv.Itoa(line+2), d[col]); err != nil {
				log.Println("cell value error:", err)
			}
		}
	}
	return f
}

func getSheet() []string {
	return []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z",
	}
}
