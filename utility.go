package meituan

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

func Sign(requestUrl string, appId string, secret string, timestamp int64, o interface{}) (string, url.Values) {
	urlValues := url.Values{}
	sb := new(strings.Builder)
	fields, params := GetStuctParamMap(o)

	fields = append(fields, "app_id", "timestamp")

	sort.Strings(fields)

	params["app_id"] = appId
	params["timestamp"] = timestamp

	sb.WriteString(requestUrl)
	sb.WriteString("?")

	for _, field := range fields {
		sb.WriteString(field)
		sb.WriteString("=")

		v := reflect.TypeOf(params[field])

		if v.Kind() != reflect.Struct {
			filedValue := fmt.Sprintf("%v", params[field])
			sb.WriteString(filedValue)
			urlValues.Set(field, filedValue)
		} else {
			if bytes, err := json.Marshal(params[field]); err == nil {
				fieldValue := string(bytes)
				sb.WriteString(fieldValue)
				urlValues.Set(field, fieldValue)
			}
		}
	}

	sb.WriteString(secret)

	sig := fmt.Sprintf("%x", md5.Sum([]byte(sb.String())))

	return fmt.Sprintf("%s?app_id=%s&timestamp=%v&sig=%s", requestUrl, appId, timestamp, sig), urlValues
}

func GetStuctParamMap(o interface{}) ([]string, map[string]interface{}) {
	params := make(map[string]interface{})

	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	elem := t.Elem()
	fields := make([]string, 0, elem.NumField())

	for i := 0; i < elem.NumField(); i++ {
		if p := elem.Field(i).Tag.Get("url"); len(p) > 0 {
			field := v.Elem().FieldByName(elem.Field(i).Name)

			if field.Kind() != reflect.Ptr || (field.Kind() == reflect.Ptr && !field.IsNil()) {
				if field.Kind() == reflect.Ptr {
					params[p] = field.Elem().Interface()
				} else {
					params[p] = field.Interface()
				}

				fields = append(fields, p)
			}
		}
	}

	return fields, params
}
