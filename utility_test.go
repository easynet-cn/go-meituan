package meituan

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Sign(t *testing.T) {
	url := "https://waimaiopen.meituan.com/api/v1/order/getOrderDaySeq"
	appId := "0000"
	secret := "a00ba58f000001aa697ab000006d52d"
	appPoiCode := "31号测试店"
	timestamp := int64(1389751221)

	urlValues := Sign(url, appId, secret, timestamp, &GetOrderDaySeqRequest{AppPoiCode: appPoiCode})
	sig := urlValues.Get("sig")

	if sig != "dbb4d444d68596d03dbcb79a4e4a4e3f" {
		t.Error(sig)
	}

	fmt.Println(sig)
}

func Test_GetTag(t *testing.T) {
	name := "go"

	s := &A{Id: 1, Tags: []string{"t1", "t2"}, Name: &name}

	e := reflect.TypeOf(s).Elem()

	for i := 0; i < e.NumField(); i++ {
		fmt.Println(e.Field(i).Tag.Get("url")) //将tag输出出来
	}
}

func Test_GetParams(t *testing.T) {
	name := "go"

	s := &A{Id: 1, Tags: []string{"t1", "t2"}, Name: &name}

	fields, params := GetStuctParamMap(s)

	fmt.Println(fields)
	fmt.Println(params)
}

type A struct {
	Id   int64    `url:"id"`
	Tags []string `url:"tags"`
	Name *string  `url:"name"`
}

type GetOrderDaySeqRequest struct {
	AppPoiCode string `url:"app_poi_code"`
}
