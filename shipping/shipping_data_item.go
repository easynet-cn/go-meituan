package shipping

type ShippingDataItem struct {
	MinPrice                  float64  `json:"min_price"`                    //是否必须：创建时必传，更新时非必传；描述：该配送区域的起送价，单位是元。
	Area                      string   `json:"area"`                         //是否必须：是；描述：配送范围type=1时:[{"x":39941199,"y":116385384},{"x":39926983,"y":116361694},{"x ":39921586,"y":116398430}]，type 为 2 时: {"r":1000,"x":40029380,"y":116418390}，需要对其urlEncode，x代表纬度，y代表经度。 美团使用的是高德坐标系，即火星坐标系，配送范围坐标需要乘以一百万；如果是百度坐标系需要转换。 注意：X坐标要在[18000000,54000000]范围内，Y坐标要在[73800000,135000000]范围内。
	AppShippingCode           string   `json:"app_shipping_code"`            //是否必须：是；描述：APP方提供的配送范围id，创建配送范围时由商家自行赋值。 (1)支持上传多个配送范围id。 (2)如门店有多个配送范围，更新时未上传的配送范围将会被删除。注意：接口不支持删除无配送范围id的配送范围。 (3)非开放平台接口创建的配送范围没有配送范围id，因此无法通过接口更新此配送范围的信息。商家如需通过接口更新配送范围，建议重新通过接口创建，可在商家端操作删除原配送范围。 (4)通过接口创建的配送范围，建议不要在其他端(例如商家端)操作修改，否则将无法再使用接口更新。
	Type                      int      `json:"type"`                         //是否必须：是；描述：配置范围类型，1表示多个配送范围由多个多边形组成。（平台只支持多边形）
	ShippingFee               *float64 `json:"shipping_fee"`                 //是否必须：否；描述：该配送区域的配送费，单位是元。建议上传这个字段设定该区域配送费。 如果不传此字段，则该区域配送费是以门店当前已保存的所有配送范围配送费的最大值生效；如无配送范围或其配送费为0元，则该区域配送费为0元。
	DistanceMarkupExecuteType *int     `json:"distance_markup_execute_type"` //是否必须：否；描述：按距离加价规则的执行类型：1-启用，2-不启用，3-保持原值 如果门店没有初始规则且不需要，则不传即可
	DistanceMarkupFactors     string   `json:"distance_markup_factors"`      //是否必须：distance_markup_execute_type传递1-启用时必传；描述：按距离加价， 1.上传多个规则时，需按从小到大顺序上传 2.最多支持设置5个加价规则 3.传启用则新上传的规则覆盖原有规则 4.传不启用则删除已有规则
	WeightMarkupExecuteType   *int     `json:"weight_markup_execute_type"`   //是否必须：否；描述：按重量加价规则的执行类型：1-启用，2-不启用，3-保持原值 如果门店没有初始规则且不需要，则不传即可
	WeightMarkupFactors       string   `json:"weight_markup_factors"`        //是否必须：weight_markup_execute_type传递1-启用时必传；描述：按重量加价， 1.上传多个规则时，需按从小到大顺序上传 2.最多支持设置5个加价规则 3.传启用则新上传的规则覆盖原有规则 4.传不启用则删除已有规则
	TimeMarkupExecuteType     *int     `json:"time_markup_execute_type"`     //是否必须：否；描述：按时段加价规则的执行类型：1-启用，2-不启用，3-保持原值 如果门店没有初始规则且不需要，则不传即可
	TimeMarkupFactors         string   `json:"time_markup_factors"`          //是否必须：time_markup_execute_type传递1-启用时必传；描述：按时段加价， 1.最多支持设置5个加价规则 2.传启用则新上传的规则覆盖原有规则 3.传不启用则删除已有规则
}
