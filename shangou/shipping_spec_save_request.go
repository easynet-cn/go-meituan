package shangou

//新增/更新特殊时段配送范围
//请求地址：https://waimaiopen.meituan.com/api/v1/shipping/spec/save
type ShippingSpecSaveRequest struct {
	AppPoiCode                string  `url:"app_poi_code"`                 //是否必须：是；描述：APP方门店id，传商家中台系统里门店的编码，最长不超过128个字符。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	AppShippingCode           string  `url:"app_shipping_code"`            //是否必须：是；描述：APP方提供的配送范围id；非开放平台接口操作创建的配送范围，无配送范围id。 如一个门店没有配送范围id，且门店只有一个配送范围，此字段可以传1。
	Type                      string  `url:"type"`                         //是否必须：是；描述：配置范围类型，1表示多个配送范围由多个多边形组成。（平台只支持多边形）
	Area                      string  `url:"area"`                         //是否必须：是；描述：配送范围type=1时:[{"x":39941199,"y":116385384},{"x":39926983,"y":116361694},{"x ":39921586,"y":116398430}]，需要对其urlEncode，x代表纬度，y代表经度。 美团使用的是高德坐标系，即火星坐标系，配送范围坐标需要乘以一百万；如果是百度坐标系需要转换。 注意：X坐标要在[18000000,54000000]范围内，Y坐标要在[73800000,135000000]范围内。
	MinPrice                  float64 `url:"min_price"`                    //是否必须：创建时必传，更新时非必传；描述：该配送区域的起送价，单位是元。
	ShippingFee               float64 `url:"shipping_fee"`                 //是否必须：是；描述：该配送区域的配送费，单位是元。 （建议上传这个字段设定配送费，如果此字段为空，则以门店保存的配送费为准）
	TimeRange                 string  `url:"time_range"`                   //是否必须：是；描述：此配送范围的生效时间，最多允许上传1个时间段。允许跨天，但需小于24小时，如时间段开始时间小于结束时间，则认为是跨天。
	DistanceMarkupExecuteType int     `url:"distance_markup_execute_type"` //是否必须：否；描述：按距离加价规则的执行类型：1启用，2不启用，3保持原值
	DistanceMarkupFactors     string  `url:"distance_markup_factors"`      //是否必须：distance_markup_execute_type传递1启用时必传；描述：按距离加价， 1.上传多个规则时，需按从小到大顺序上传 2.最多支持设置5个加价规则 3.传启用则新上传的规则覆盖原有规则 4.传不启用则删除已有规则
	WeightMarkupExecuteType   int     `url:"weight_markup_execute_type"`   //是否必须：否；描述：按重量加价规则的执行类型：1启用，2不启用，3保持原值
	WeightMarkupFactors       string  `url:"weight_markup_factors"`        //是否必须：weight_markup_execute_type传递1启用时必传；描述：按重量加价， 1.上传多个规则时，需按从小到大顺序上传 2.最多支持设置5个加价规则 3.传启用则新上传的规则覆盖原有规则 4.传不启用则删除已有规则
	TimeMarkupExecuteType     int     `url:"time_markup_execute_type"`     //是否必须：否；描述：按时段加价规则的执行类型：1启用，2不启用，3保持原值
	TimeMarkupFactors         string  `url:"time_markup_factors"`          //是否必须：time_markup_execute_type传递1启用时必传；描述：按时段加价， 1.最多支持设置5个加价规则 2.传启用则新上传的规则覆盖原有规则 3.传不启用则删除已有规则
}
