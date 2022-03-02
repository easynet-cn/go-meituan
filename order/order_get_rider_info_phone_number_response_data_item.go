package order

type OrderGetRiderInfoPhoneNumberResponseDataItem struct {
	OrderId              int64  `json:"order_id"`                //描述：订单号
	WmOrderIdView        int64  `json:"wm_order_id_view"`        //描述：订单展示ID
	RiderName            string `json:"rider_name"`              //描述：骑手名字
	RiderRealPhoneNumber string `json:"rider_real_phone_number"` //描述：骑手真实手机号
	AppPoiCode           string `json:"app_poi_code"`            //描述：门店id
}
