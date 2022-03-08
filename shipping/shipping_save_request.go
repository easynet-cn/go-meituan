package shipping

//创建/更新门店配送范围
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/shipping/save
//文档地址：https://open-shangou.meituan.com/home/docDetail/42
type ShippingSaveRequest struct {
	AppPoiCode                string   `url:"app_poi_code"`                 //是否必须：是；描述：APP方门店id，传商家中台系统里门店的编码，最长不超过128个字符。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	AppShippingCode           string   `url:"app_shipping_code"`            //是否必须：是；描述：APP方提供的配送范围id，创建配送范围时由商家自行赋值。 (1)仅支持上传一个配送范围id。 (2)非开放平台接口创建的配送范围没有配送范围id，因此无法通过接口更新此配送范围的信息。商家如需通过接口更新配送范围，建议重新通过接口创建，可在商家端操作删除原配送范围。 (3)通过接口创建的配送范围，建议不要在其他端(例如商家端)操作修改，否则将无法再使用接口更新。
	Type                      string   `url:"type"`                         //是否必须：是；描述：配置范围类型，1表示多个配送范围由多个多边形组成。（平台只支持多边形）
	Area                      string   `url:"area"`                         //是否必须：是；描述：配送范围，type=1时：[{"x":29777353,"y":95365224},{"x":29778620,"y":95373893},{"x":29771319,"y":95376553},{"x":29769493,"y":95369944},{"x":29770648,"y":95362735},{"x":29773815,"y":95362563}]，需要对其urlEncode，x代表纬度，y代表经度。 美团使用的是高德坐标系，即火星坐标系，配送范围坐标需要乘以一百万；如果是百度坐标系需要转换。 注意：X坐标要在[18000000,54000000]范围内，Y坐标要在[73800000,135000000]范围内。
	MinPrice                  float64  `url:"min_price"`                    //是否必须：是；描述：该配送区域的起送价，单位是元。
	ShippingFee               *float64 `url:"shipping_fee"`                 //是否必须：否；描述：该配送区域的配送费，单位是元。建议上传这个字段设定该区域配送费。 如果不传此字段，则该区域配送费是以门店当前已保存的所有配送范围配送费的最大值生效；如无配送范围或其配送费为0元，则该区域配送费为0元。
	DistanceMarkupExecuteType *int     `url:"distance_markup_execute_type"` //是否必须：否；描述：按距离加价规则的执行类型：1-启用，2-不启用，3-保持原值 如果门店没有初始规则且不需要，则不传即可
	DistanceMarkupFactors     string   `url:"distance_markup_factors"`      //是否必须：是；描述：按距离加价， 1.上传多个规则时，需按从小到大顺序上传 2.最多支持设置5个加价规则 3.传启用则新上传的规则覆盖原有规则 4.传不启用则删除已有规则
	WeightMarkupExecuteType   *int     `url:"weight_markup_execute_type"`   //是否必须：否；描述：按重量加价规则的执行类型：1-启用，2-不启用，3-保持原值 如果门店没有初始规则且不需要，则不传即可
	WeightMarkupFactors       string   `url:"weight_markup_factors"`        //是否必须：是；描述：按重量加价， 1.上传多个规则时，需按从小到大顺序上传 2.最多支持设置5个加价规则 3.传启用则新上传的规则覆盖原有规则 4.传不启用则删除已有规则
	TimeMarkupExecuteType     *int     `url:"time_markup_execute_type"`     //是否必须：否；描述：按时段加价规则的执行类型：1-启用，2-不启用，3-保持原值 如果门店没有初始规则且不需要，则不传即可
	TimeMarkupFactors         string   `url:"time_markup_factors"`          //是否必须：是；描述：按时段加价， 1.最多支持设置5个加价规则 2.传启用则新上传的规则覆盖原有规则 3.传不启用则删除已有规则
}
