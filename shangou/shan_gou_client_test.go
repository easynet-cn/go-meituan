package shangou

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_ShanGouClient(t *testing.T) {
	httpClient := &http.Client{}

	shanGouClient := &ShanGouClient{
		HttpClient: httpClient,
		AppId:      "0000",
		AppSecret:  "a00ba58f000001aa697ab000006d52d",
	}

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/order/getOrderDaySeq", "GET", "", &GetOrderDaySeqRequest{AppPoiCode: "31号测试店"})

	fmt.Println(string(bytes), err)
}

type GetOrderDaySeqRequest struct {
	AppPoiCode string `url:"app_poi_code"`
}
