package core

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"reflect"
	"runtime"
	"strconv"

	"github.com/xuri/excelize/v2"
)

/**
 * Created by frankieci on 2021/12/28 9:56 pm
 */

// DeepCopy vs copy(built-in), its diff
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func GetFunctionName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func WriteExcelFile(sheetName string, sheetHeader []string, data [][]interface{}) *excelize.File {
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

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
