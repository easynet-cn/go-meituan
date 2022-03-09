package shipping

type ShippingFetchItem struct {
	MinPrice        float64 `json:"min_price"`         //描述：该配送区域的起送价，单位是元。
	Area            string  `json:"area"`              //描述：配送范围type=1时:[{"x":39941199,"y":116385384},{"x":39926983,"y":116361694},{"x ":39921586,"y":116398430}]，type 为 2 时: {"r":1000,"x":40029380,"y":116418390}，需要对其urlEncode，x代表纬度，y代表经度。 美团使用的是高德坐标系，即火星坐标系，配送范围坐标需要乘以一百万；如果是百度坐标系需要转换。
	AppShippingCode string  `json:"app_shipping_code"` //描述：APP方提供的配送范围id；非开放平台接口操作创建的配送范围，无配送范围id。
	ShippingFee     float64 `json:"shipping_fee"`      //描述：该配送区域的配送费，单位是元。 (建议上传这个字段信息设定该配送范围的配送费；如此字段为空，则以门店保存的配送费为准。)
	Type            int     `json:"type"`              //描述：配置范围类型，1表示多个配送范围由多个多边形组成。（平台只支持多边形）
	LogisticsCode   string  `json:"logistics_code"`    //描述：此配送范围的配送方式，参考值：2002-快送。
}
