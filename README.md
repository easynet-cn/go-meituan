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
```

