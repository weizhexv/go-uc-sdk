package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ResToObj(resp *http.Response, target any) error {
	bs, err := read(resp)
	if err != nil {
		return err
	}
	if len(bs) == 0 {
		return nil
	}
	return json.Unmarshal(bs, target)
}

func read(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, nil
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func JoinInt64Arr(int64arr []int64) string {
	var arr []string
	for _, uid := range int64arr {
		arr = append(arr, strconv.FormatInt(uid, 10))
	}
	return strings.Join(arr, ",")
}
