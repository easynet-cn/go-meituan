package shipping

type ShippingListItem struct {
	MinPrice              float64 `json:"min_price"`               //描述：该配送区域的起送价，单位是元。
	Area                  string  `json:"area"`                    //描述：配送范围，type=1时：[{"x":29777353,"y":95365224},{"x":29778620,"y":95373893},{"x":29771319,"y":95376553},{"x":29769493,"y":95369944},{"x":29770648,"y":95362735},{"x":29773815,"y":95362563}]，需要对其urlEncode，x代表纬度，y代表经度。 美团使用的是高德坐标系，即火星坐标系，配送范围坐标需要乘以一百万；如果是百度坐标系需要转换。
	AppShippingCode       string  `json:"app_shipping_code"`       //描述：APP方提供的配送范围id；非开放平台接口操作创建的配送范围，无配送范围id。
	ShippingFee           float64 `json:"shipping_fee"`            //描述：该配送区域的配送费，单位是元。 (建议上传这个字段信息设定该配送范围的配送费；如此字段为空，则以门店保存的配送费为准。)
	Type                  int     `json:"type"`                    //描述：配置范围类型，1表示多个配送范围由多个多边形组成。（平台只支持多边形）
	DistanceMarkupFactors string  `json:"distance_markup_factors"` //描述：按距离加价规则
	WeightMarkupFactors   string  `json:"weight_markup_factors"`   //描述：按重量加价规则
	TimeMarkupFactors     string  `json:"time_markup_factors"`     //描述：按时段加价规则
}
