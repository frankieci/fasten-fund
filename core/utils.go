package core

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"runtime"
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
