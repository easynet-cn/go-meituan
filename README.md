# go-meituan 
# 美团闪购开放平台-医疗健康 glang sdk

---

- [act] 活动类接口 [接口文档地址](https://open-shangou.meituan.com/home/doc/medical/25) 
- [app] APP方接口 [订单推送接口文档地址](https://open-shangou.meituan.com/home/doc/medical/27)  [取消订单推送接口文档地址](https://open-shangou.meituan.com/home/doc/medical/28) [退款信息推送接口文档地址](https://open-shangou.meituan.com/home/doc/medical/29) 
- [comment] 评论评价接口 [接口文档地址](https://open-shangou.meituan.com/home/doc/medical/24) 
- [image] 图片接口 [接口文档地址](https://open-shangou.meituan.com/home/doc/medical/30) 
- [isv] 服务商接口 [接口文档地址](https://open-shangou.meituan.com/home/guide/medical/10686) 
- [medicine] 药品接口 [接口文档地址](https://open-shangou.meituan.com/home/doc/medical/23) 
- [order] 订单接口 [接口文档地址](https://open-shangou.meituan.com/home/doc/medical/26) 
- [poi] 门店接口 [接口文档地址](https://open-shangou.meituan.com/home/doc/medical/21) 
- [shippping] 配送范围接口 [接口文档地址](https://open-shangou.meituan.com/home/doc/medical/22) 

## 例子

```go
package shangou

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/easynet-cn/go-meituan/isv"
	"github.com/easynet-cn/go-meituan/medicine"
)

const (
	AppId        = ""
	AppSecret    = ""
	AppPoiCode   = ""
	AccessToken  = ""
	RefreshToken = ""
)

func Test_GetAccessToken(t *testing.T) {
	shanGouClient := getClient()

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/oauth/authorize", "GET", "", &isv.GetAccessTokenRequest{AppPoiCode: AppPoiCode, ResponseType: "token"})

	fmt.Println(string(bytes), err)
}

func Test_RefreshAccessToken(t *testing.T) {
	shanGouClient := getClient()

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/oauth/token", "POST", "", &isv.RefreshAccessTokenRequest{GrantType: "refresh_token", RefreshToken: RefreshToken})

	fmt.Println(string(bytes), err)
}

func Test_GetStoreCategories(t *testing.T) {
	shanGouClient := getClient()

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/medicineCat/list", "GET", "", &medicine.MedicineCatListRequest{AppPoiCode: AppPoiCode})

	fmt.Println(string(bytes), err)
}

func Test_GetStoreCategoriesByMapParam(t *testing.T) {
	shanGouClient := getClient()

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/medicineCat/list", "GET", "", map[string]interface{}{"app_poi_code": AppPoiCode})

	fmt.Println(string(bytes), err)
}

func Test_GetMedicineList(t *testing.T) {
	shanGouClient := getClient()

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/medicine/list", "GET", AccessToken, &medicine.MedicineListRequest{AppPoiCode: AppPoiCode, Offset: 0, Limit: 200})

	fmt.Println(string(bytes), err)
}

func Test_CateMedicineCategory(t *testing.T) {
	shanGouClient := getClient()

	request := &medicine.MedicineCatRequest{
		AppPoiCode:   AppPoiCode,
		CategoryCode: "092400",
		CategoryName: "其他",
		Sequence:     24,
	}

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/medicineCat/save", "POST", AccessToken, request)

	fmt.Println(string(bytes), err)
}

func getClient() *ShanGouClient {
	return &ShanGouClient{
		HttpClient: &http.Client{},
		AppId:      AppId,
		AppSecret:  AppSecret,
	}
}

type GetOrderDaySeqRequest struct {
	AppPoiCode string `url:"app_poi_code"`
}
```

