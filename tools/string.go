package tools

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
)

func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

func GetBodyString(c *gin.Context) (string, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("read boy err, %v\n", err)
		return string(body), nil
	} else {
		return "", err
	}
}

func JsonStrToMap(e string) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	} else {
		return nil, err
	}
}
