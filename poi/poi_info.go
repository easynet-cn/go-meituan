package poi

type PoiInfo struct {
	AppPoiCode         string  `json:"app_poi_code"`        //描述：APP方门店ID
	Name               string  `json:"name"`                //描述：门店名称
	Address            string  `json:"address"`             //描述：门店地址
	Latitude           float64 `json:"latitude"`            //描述：门店纬度，美团使用的是高德坐标系。 注：如需使用获取的纬度信息在poi/save接口中上传，需用此返回信息再除以一百万。 例如本接口查询到的纬度值为2.9774566E7，那么在poi/save接口上传时，需转为29.774566（上传时平台系统会自动乘以一百万）。
	Longitude          float64 `json:"longitude"`           //描述：门店经度，美团使用的是高德坐标系。 注：如需使用获取的经度信息在poi/save接口中上传，需用此返回信息再除以一百万。 例如本接口查询到的经度值为9.5369278E7，那么在poi/save接口上传时，需转为95.369278（上传时平台系统会自动乘以一百万）。
	PicUrl             string  `json:"pic_url"`             //描述：门店图片地址
	PicUrlLarge        string  `json:"pic_url_large"`       //描述：门店图片地址
	Phone              string  `json:"phone"`               //描述：客服电话号码
	StandbyTel         string  `json:"standby_tel"`         //描述：门店电话号码(已废弃)
	ShippingFee        float64 `json:"shipping_fee"`        //描述：每个订单的配送费，单位是元。
	ShippingTime       string  `json:"shipping_time"`       //描述：门店营业时间，多个时间段之间以英文逗号分隔。
	PromotionInfo      string  `json:"promotion_info"`      //描述：门店公告信息
	InvoiceSupport     int     `json:"invoice_support"`     //描述：门店是否支持发票，参考值：1-支持；0-不支持。
	InvoiceMinPrice    int     `json:"invoice_min_price"`   //描述：门店支持开发票的最小订单价，单位是元。
	InvoiceDescription string  `json:"invoice_description"` //描述：发票相关说明
	OpenLevel          int     `json:"open_level"`          //描述：门店的营业状态，参考值：1-可配送；3-休息中。
	IsOnline           int     `json:"is_online"`           //描述：门店上下线状态，参考值：0-下线；1-上线；2-上单中；3-审核通过可上线。
	CTime              int     `json:"ctime"`               //描述：门店创建时间，返回10位秒级时间戳。（当前距离Epoch（1970年1月1日） 以秒计算的时间，即unix-timestamp）
	UTime              int     `json:"utime"`               //描述：门店最近一次更新时间，返回10位秒级的时间戳。（当前距离Epoch（1970年1月1日） 以秒计算的时间，即unix-timestamp）
	ThirdTagName       string  `json:"third_tag_name"`      //描述：门店经营品类，多个品类信息以英文逗号分隔。
	PreBook            int     `json:"pre_book"`            //描述：是否支持营业时间范围外预下单，参考值：1-支持；0-不支持。
	TimeSelect         int     `json:"time_select"`         //描述：是否支持营业时间范围内预下单，参考值：1-支持；0-不支持。
	LogisticsCodes     string  `json:"logistics_codes"`     //描述：门店的配送方式,参考值： 门店的配送方式,参考值： 1003-美团跑腿（众包） 1001-专送（加盟） 1002-专送（自建） 1004-城市代理 2002-快送 2010-全城送 0000-商家自配 3001-混合送（专送+快送) 30011002-混合自建 30011001-混合加盟 30012002-混合快送 0002-趣生活美食配送 0016-达达快递 0033-E_代送
	PreBookMinDays     int     `json:"pre_book_min_days"`   //描述：商家接受预订日期的最早日期，范围：0-7。最早日期是指用户最少需要提前下单的天数，如果不选“当天”（传0），则用户不能下要求今日送达的订单。例如：最早日期为“隔天”（传1），则用户今日最快只能下明天备货配送的订单。
	PreBookMaxDays     int     `json:"pre_book_max_days"`   //描述：商家接受预订日期的最长日期，取值范围：0-7。最长日期是指用户可要求送达的最多天数。例如：最长日期为“隔天”（传1），则用户今日可下要求明天配送的订单，不可下要求后天配送的订单。
}
