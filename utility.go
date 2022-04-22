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

func Sign(requestUrl string, appId string, secret string, accessToken string, timestamp int64, o interface{}) (string, url.Values) {
	urlValues := url.Values{}
	sb := new(strings.Builder)
	fields, params := GetStuctParamMap(o)

	fields = append(fields, "app_id", "timestamp")

	if accessToken != "" {
		fields = append(fields, "access_token")
	}

	sort.Strings(fields)

	params["app_id"] = appId
	params["timestamp"] = timestamp

	if accessToken != "" {
		params["access_token"] = accessToken
	}

	sb.WriteString(requestUrl)
	sb.WriteString("?")

	for i, field := range fields {
		if i > 0 {
			sb.WriteString("&")
		}

		sb.WriteString(field)
		sb.WriteString("=")

		v := reflect.TypeOf(params[field])

		if v.Kind() != reflect.Struct {
			fieldValue := fmt.Sprintf("%v", params[field])
			sb.WriteString(fieldValue)

			if field != "app_id" && field != "timestamp" && field != "access_token" {
				urlValues.Set(field, fieldValue)
			}
		} else {
			if bytes, err := json.Marshal(params[field]); err == nil {
				fieldValue := string(bytes)
				sb.WriteString(fieldValue)

				if field != "app_id" && field != "timestamp" && field != "access_token" {
					urlValues.Set(field, fieldValue)
				}
			}
		}
	}

	sb.WriteString(secret)

	sig := fmt.Sprintf("%x", md5.Sum([]byte(sb.String())))

	if accessToken != "" {
		return fmt.Sprintf("%s?access_token=%v&app_id=%s&timestamp=%v&sig=%s", requestUrl, accessToken, appId, timestamp, sig), urlValues
	}

	return fmt.Sprintf("%s?app_id=%s&timestamp=%v&sig=%s", requestUrl, appId, timestamp, sig), urlValues
}

func GetStuctParamMap(o interface{}) ([]string, map[string]interface{}) {
	if o == nil {
		return make([]string, 0), make(map[string]interface{})
	}

	params := make(map[string]interface{})

	t := reflect.TypeOf(o)

	if t.Kind() == reflect.Map {
		m := o.(map[string]interface{})
		fields := make([]string, 0, len(m))

		for k, v := range m {
			params[k] = v
			fields = append(fields, k)
		}

		return fields, params
	}

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
