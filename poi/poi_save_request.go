package poi

//更新门店信息
//请求地址：https://waimaiopen.meituan.com/api/v1/poi/save
type PoiSaveRequest struct {
	AppPoiCode         string  `url:"app_poi_code"`        //是否必须：是；描述：APP方门店id，传商家中台系统里门店的编码，最长不超过128个字符。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	Name               string  `url:"name"`                //是否必须：是；描述：门店名称
	Address            string  `url:"address"`             //是否必须：是；描述：门店地址
	Latitude           float64 `url:"latitude"`            //是否必须：是；描述：门店纬度，美团使用的是高德坐标系，即火星坐标系。 开发者上传参数请参考示例值，平台系统会自动乘以一百万。
	Longitude          float64 `url:"longitude"`           //是否必须：是；描述：门店经度 ，美团使用的是高德坐标系，即火星坐标系。 开发者上传参数请参考示例值，平台系统会自动乘以一百万。
	PicUrl             string  `url:"pic_url"`             //是否必须：否；描述：门店图片地址（默认图片为：http://p1.meituan.net/crm/__37375183__1582979.jpg）。需要为JPG/JPEG格式。可自行适配前端4:3的展示比例上传. 若更新pic_url，则需要人工审核，不会立即生效。
	PicUrlLarge        string  `url:"pic_url_large"`       //是否必须：否；描述：门店图片地址（默认图片为：http://p1.meituan.net/crm/__37375183__1582979.jpg ），需要为JPG/JPEG格式。 （目前此参数无效，使用pic_url参数即可。）
	Phone              string  `url:"phone"`               //是否必须：新增时必填；描述：客服电话号码，多个号码之间以/分隔。 （注意：此号码留客服号码）
	StandbyTel         string  `url:"standby_tel"`         //是否必须：否；描述：门店电话号码 （注意：此号码留商家电话）(已废弃)
	ShippingFee        float64 `url:"shipping_fee"`        //是否必须：新增时必填；描述：每个订单的配送费，单位是元。 更新时，如果不传配送费，平台会查询门店当前所有配送范围的配送费取最大的值。如果门店当前没有配送范围或其配送费均为0，则门店配送费取值为0。
	ShippingTime       string  `url:"shipping_time"`       //是否必须：新增时必填；描述：门店营业时间，多个时间段之间用英文逗号分隔。B2C医药门店，暂不支持修改门店营业时间。注意格式:每个时段之间不能存在交集。
	PromotionInfo      string  `url:"promotion_info"`      //是否必须：否；描述：门店公告信息
	OpenLevel          int     `url:"open_level"`          //是否必须：新增时必填；描述：门店的营业状态，取值范围：1可配送；3休息中。 设置门店恢复营业状态的角色的权限需要与上一次设置门店置休的角色的权限一致。所以，即使当前门店是营业状态，但是这个门店上次置休的时候是总部（商家总账号）操作的，所以再次设置营业状态时仍需要总部设置才可以。如使用接口操作，会返回没有权限的提示。B2C医药门店，暂不支持设置休息中。
	IsOnline           int     `url:"is_online"`           //是否必须：新增时必填；描述：门店上下线状态，取值范围：1上线；0下线。 注意：创建门店时，门店必须在商品、配送范围和门店信息都齐全后，才能提交上线，即创建门店时此字段须传0。 B2C医药门店，暂不支持设置为下线。 提交门店基本信息后，使用商品和配送类接口分别上传商品和配送信息，然后再通过本接口修改门店状态为上线，即表示提交审核。（注意：此字段不为1时门店不会提交审核）
	InvoiceSupport     int     `url:"invoice_support"`     //是否必须：否；描述：门店是否支持发票，取值范围：1支持；0不支持。
	InvoiceMinPrice    int     `url:"invoice_min_price"`   //是否必须：否；描述：门店支持开发票的最小订单价，单位是元。当invoice_support=1时，本字段可用。 注意：根据当地法律法规，发票最低金额应为0。
	InvoiceDescription string  `url:"invoice_description"` //是否必须：否；描述：发票相关说明，当invoice_support=1时，本字段可用。
	ThirdTagName       string  `url:"third_tag_name"`      //是否必须：新增时必填；描述：目前最多支持上传两个门店品类：一个主营品类（上传的第一个品类为主营品类），一个辅营品类，以英文逗号分隔。 部分门店品类只支持上传一个品类，例如：火锅，特色菜，地方菜，东南亚菜，日韩料理，生活超市。 通过接口【poiTag/list】可获取门店品类信息列表。 对闪购品类应用 门店，当调用此接口时，不允许商家对该字段进行修改。如商家入参则直接忽略。
	PreBook            int     `url:"pre_book"`            //是否必须：否；描述：是否支持营业时间范围外预下单，取值范围：1支持；0不支持。B2C医药门店，暂不支持设置营业时间外用户预下单。
	TimeSelect         int     `url:"time_select"`         //是否必须：否；描述：是否支持营业时间范围内预下单，取值范围：1支持；0不支持。此字段开启后不支持关闭，新建门店将自动开启。
	AppBrandCode       string  `url:"app_brand_code"`      //是否必须：否；描述：第三方品牌code值（对接的三方需要提前将该值告诉美团技术同学）
	PreBookMinDays     int     `url:"pre_book_min_days"`   //是否必须：否；描述：商家接受预订日期的最早日期，取值范围：07。最早日期是指用户最少需要提前下单的天数，如果不选“当天”（传0），则用户不能下要求今日送达的订单。例如：最早日期为“隔天”（传1），则用户今日最快只能下明天备货配送的订单。
	PreBookMaxDays     int     `url:"pre_book_max_days"`   //是否必须：否；描述：商家接受预订日期的最长日期，取值范围：07。最长日期是指用户可要求送达的最多天数。例如：最长日期为“隔天”（传1），则用户今日可下要求明天配送的订单，不可下要求后天配送的订单。
}
