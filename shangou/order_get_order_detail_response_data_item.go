package shangou

type OrderGetOrderDetailResposneDataItem struct {
	OrderId                   int64   `json:"order_id"`                      //描述：订单号（同订单展示ID），数据库中请用bigint(20)存储此字段。
	OrderTagList              []int   `json:"order_tag_list"`                //描述：订单信息，8代表处方药，23代表医保药
	WmOrderIdView             int64   `json:"wm_order_id_view"`              //描述：订单展示ID，与用户端、商家端订单详情中展示的订单号码一致。数据库中请用bigint(20)存储此字段。
	AppPoiCode                string  `json:"app_poi_code"`                  //描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	WmPoiName                 string  `json:"wm_poi_name"`                   //描述：商家门店名称，即美团平台上当前订单所属门店的名称。
	WmPoiAddress              string  `json:"wm_poi_address"`                //描述：商家门店地址，即美团平台上当前订单所属门店的地址。
	RecipientAddress          string  `json:"recipient_address"`             //描述：订单收货人地址，此字段为用户填写的收货地址。商家可在开发者中心->基础设置->订单订阅字段 页面订阅字段“13 收货地址”；若开启订阅，则订单数据会在此字段后追加根据经纬度反查地址的结果，并用“@#”符号分隔，格式为：用户填写地址@#反查结果。 示例：美团·大众点评北京总部 (恒电大厦B座一层)@#北京市北京市朝阳区屏翠东路屏翠东路美团·大众点评北京总部
	RecipientAddressDetail    string  `json:"recipient_address_detail"`      //描述：订单收货人地址，按省、市、区、 街道、详细地址进行拆分。 该字段为医药B2C商家专用，用于商家快递发货。 注意：可能存在查询地址服务时没有查到四级地址信息，此情况下本参数不返回。
	RecipientPhone            string  `json:"recipient_phone"`               //描述：订单收货人联系电话，此字段信息可能推送真实手机号码或隐私号，即需兼容138*****678和13812345678_123456两种号码格式。
	BackupRecipientPhone      string  `json:"backup_recipient_phone"`        //描述：订单收货人姓名，同时包含用户在用户端选择的性别标识信息。(1)若用户没有填写姓名(老版本的用户端可能支持不填写)，此字段默认为空；商家可在开发者中心开启订阅“12 用户名称”字段，则平台会用“美团客人”自动填充此字段。(2)用户填写了收货人姓名时，此字段按用户所填全部内容正常展示，若姓名中包含特殊符号(如emoji表情)，需商家自行解析。
	ShippingFee               float64 `json:"shipping_fee"`                  //描述：门店配送费，单位是元。当前订单产生时该门店的配送费（商家自配送运费或美团配送运费），此字段数据为运费优惠前的原价。
	SgOpenThirdLogisticsInfos string  `json:"sg_open_third_logistics_infos"` //描述：仅适用于医药b2c（快递配送）订单，注意订单下可能存在多条快递信息。
	Total                     float64 `json:"total"`                         //描述：订单的实际在线支付总价，单位是元。此字段数据为用户实际支付的订单总金额，含打包袋、配送费等。
	OriginalPrice             float64 `json:"original_price"`                //描述：订单的总原价，单位是元。此字段数据为未扣减所有优惠前订单的总金额，含打包袋、配送费等。
	Caution                   string  `json:"caution"`                       //描述：订单备注信息。本参数支持提供的信息说明： ①用户下单时自行填写的备注信息。 ②用户下单时“如遇缺货”处的信息。 ③如果订单的取货方式为用户到店自取（pick_type=1），则本参数包含信息“到店自取”。 ④如果商家是美团发票合作商家，支持在线开具电子发票，则本参数包含信息“电子发票已自动开具”。 ⑤如果隐私号服务正常且用户开启了号码保护，本参数包含收货人的脱敏真实号码和隐私号码，示例“收餐人隐私号 18500000000_3473，手机号 185****0000”。 ⑥如果是鲜花绿植品类的订单，在隐私号服务正常且用户开启了号码保护下，本参数支持提供预订人手机号的隐私号码，请开发者支持接收隐私号格式。
	ShipperPhone              string  `json:"shipper_phone"`                 //描述：配送员联系电话，如为美团配送订单，当订单有配送员信息时，查询此字段有值；如为商家自配送订单，商家自行同步了配送员电话，查询则此字段有值。
	Status                    int     `json:"status"`                        //描述：订单状态，返回订单当前的状态。目前平台的订单状态参考值有：1-用户已提交订单；2-向商家推送订单；4-商家已确认；8-订单已完成；9-订单已取消。
	CityId                    int64   `json:"city_id"`                       //描述：城市id，此字段信息为美团定义，商家无需使用。
	HasInvoiced               int     `json:"has_invoiced"`                  //描述：是否支持开发票：0-不支持，1-支持。
	InvoiceTitle              string  `json:"invoice_title"`                 //描述：发票抬头，为用户填写的开发票的抬头。
	TaxpayerId                string  `json:"taxpayer_id"`                   //描述：纳税人识别号，此字段信息默认不返回，如商家支持订单开发票，可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“27 纳税人识别号”。
	CTime                     int64   `json:"ctime"`                         //描述：订单创建时间，为10位秒级的时间戳，此字段为用户提交订单的时间。
	UTime                     int64   `json:"utime"`                         //描述：订单更新时间，为10位秒级的时间戳，此字段信息为当前订单最新订单/配送单状态更新的时间。
	DeliveryTime              int64   `json:"delivery_time"`                 //描述：预计送达时间。若用户选择立即送出（系统自动计算预计送达时间，请参考estimate_arrival_time字段），则本参数推送0。若用户自行选择了预计送达时间，则本参数推送用户所选时间的时间戳（秒级）。
	IsThirdShipping           int     `json:"is_third_shipping"`             //描述：是否是第三方配送平台配送，0-否（含商家自配送和美团配送，目前无法区分，具体配送方式建议参考logistics_code字段。）；1-是（第三方平台配送）。
	PayType                   int     `json:"pay_type"`                      //描述：支付类型：1-货到付款，2-在线支付。目前订单只支持在线支付，此字段推送信息为2。
	PickType                  int     `json:"pick_type"`                     //描述：取货类型：0-普通(配送),1-用户到店自取。此字段的信息默认不推送，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“28 取餐类型订阅字段”。
	Latitude                  float64 `json:"latitude"`                      //描述：订单收货地址的纬度，美团使用的是高德坐标系，也就是火星坐标系，商家如果使用的是百度坐标系需要自行转换，坐标需要乘以一百万。
	Longitude                 float64 `json:"longitude"`                     //描述：订单收货地址的经度，美团使用的是高德坐标系，也就是火星坐标系，商家如果使用的是百度坐标系需要自行转换，坐标需要乘以一百万。
	DaySeq                    *int    `json:"day_seq"`                       //描述：当日订单流水号，门店每日已支付订单的流水号从1开始。未支付的订单，流水号为0。 目前，自提订单的取货码与该订单流水号相同。 此字段信息默认不返回，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“15 订单流水号”。
	IsFavorites               *bool   `json:"is_favorites"`                  //描述：订单用户是否收藏此门店：true-是， false-否。该信息默认不返回，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“17 是否收藏店铺”。
	IsPoiFirstOrder           *bool   `json:"is_poi_first_order"`            //描述：订单用户是否第一次在此门店下单：true-是，false-否。该信息默认不返回，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“16 是否是门店新客”。
	IsPreSaleOrder            bool    `json:"is_pre_sale_order"`             //描述：是否为预售单，true-是，false-否。如果为预售单，可以订阅estimate_arrival_time字段获取预计发货时间
	DinnersNumber             int     `json:"dinners_number"`                //描述：用餐人数（0：用户没有选择用餐人数；1-10：用户选择的用餐人数；-10：10人以上用餐；99：用户不需要餐具）。该信息默认不返回，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“25 用餐人数”。此字段适用于餐饮类订单，闪购品类订单此字段信息默认为0。
	LogisticsCode             *string `json:"logistics_code"`                //描述：订单配送方式，该字段信息默认不返回，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“22 配送方式”。商家可在开放平台的【附录】文档中对照查看不同logistics_code对应的描述，如0000-商家自配、1001-美团加盟、2002-快送、3001-混合送（即美团专送+快送）等。 如商家想了解自己门店的配送方式以及如何区分等情况，请咨询美团品牌经理。
	PoiReceiveDetail          *string `json:"poi_receive_detail"`            //描述：订单维度的商家对账信息，json格式数据。该信息默认不返回，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“19 商家应收款详情”。
	Detail                    string  `json:"detail"`                        //描述：订单商品详情，其值为由list序列化得到的json字符串
	Extras                    string  `json:"extras"`                        //描述：订单优惠信息，其值为由list序列化得到的json字符串。
	SkuBenefitDetail          *string `json:"sku_benefit_detail"`            //描述：商品优惠详情，仅适用于美团闪购品类业务订单，当订单中有商品享有活动优惠分摊时，会在此字段中展示商品分摊明细。该字段默认不返回，家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“31 商品优惠详情”。展示规则：(1)只展示享优惠分摊的活动商品均摊明细；未参与优惠分摊的商品相关信息不在此字段中展示。(2)当订单中同一商品sku参与同一种活动（活动id相同），多件数量时则商品优惠明细会合并展示；同时wmAppOrderActDetails中count字段表示此商品sku参与本活动的次数。
	UserMemberInfo            *string `json:"user_member_info"`              //描述：订单用户会员信息，仅适用于闪购品类商家。当商家已对接美团会员相关接口，且门店已开通会员功能，用户在当前会员门店注册或绑定会员后，如在此会员门店下单，此字段信息为用户的会员信息。 此字段默认不返回，已满足条件（已对接会员API）的商家如有需求可在官网提工单申请开通权限，待权限开通后还需自行在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“32 用户会员信息”。 如用户不是当前会员门店的会员，此字段信息返回{"is_poi_member":false}。
	AvgSendTime               float64 `json:"avg_send_time"`                 //描述：门店平均送货时长，单位是秒。是美团计算的此门店平均送货时长，如商家对此字段数据计算方式有疑问，请商家咨询美团品牌经理。
	OrderSendTime             int64   `json:"order_send_time"`               //描述：用户下单完成支付的时间，为10位秒级的时间戳。
	OrderReceiveTime          int64   `json:"order_receive_time"`            //描述：商家收到订单的时间，为10位秒级的时间戳。
	OrderConfirmTime          int64   `json:"order_confirm_time"`            //描述：商家确认订单(接单)的时间，为10位秒级的时间戳。
	OrderCancelTime           int64   `json:"order_cancel_time"`             //描述：订单取消时间，为10位秒级的时间戳。
	OrderCompletedTime        int64   `json:"order_completed_time"`          //描述：订单完成时间，为10位秒级的时间戳。
	LogisticsStatus           *int    `json:"logistics_status"`              //描述：订单配送状态code。美团配送状态值有：0-配送单发往配送，5-配送侧压单，10-配送单已确认，15-骑手已到店，20-骑手已取货，40-骑手已送达，100-配送单已取消；医药b2c（快递配送）订单配送状态值有：0-初始，1-推送给商家，20-已发货，40-已送达（目前仅建仓门店的仓库相关订单才会有这个值，无仓门店的订单没有这个枚举）。
	LogisticsId               int     `json:"logistics_id"`                  //描述：配送方id，配送方id对应配送方名称：1-商家自配；2-趣活；3-达达；4-美团跑腿；5-加盟；6-自建；7-E代送；8-城市代理；100-角马；101-快送；102-混合送（专送+快送）。
	LogisticsName             string  `json:"logistics_name"`                //描述：配送方名称，配送方id对应配送方名称：1-商家自配；2-趣活；3-达达；4-美团跑腿；5-加盟；6-自建；7-E代送；8-城市代理；100-角马；101-快送；102-混合送（专送+快送）。
	LogisticsSendTime         int     `json:"logistics_send_time"`           //描述：美团配送单发往配送的时间，此字段信息取最新一次发往配送的时间(logistics_status=0)，为10位秒级的时间戳。
	LogisticsConfirmTime      int     `json:"logistics_confirm_time"`        //描述：美团配送单确认时间，此字段信息取最新一次配送单确认的时间(logistics_status=10)，为10位秒级的时间戳。
	LogisticsCancelTime       int     `json:"logistics_cancel_time"`         //描述：美团配送单取消时间，此字段信息取最新一次配送单取消的时间(logistics_status=100)，为10位秒级的时间戳。如最新一次配送单最终未取消，则此字段值为0。
	LogisticsFetchTime        int     `json:"logistics_fetch_time"`          //描述：美团骑手取单时间，此字段信息取最新一次骑手取单的时间(logistics_status=20)，为10位秒级的时间戳。
	LogisticsCompletedTime    int     `json:"logistics_completed_time"`      //描述：美团配送单完成时间，即骑手送达订单商品的时间(logistics_status=40)，为10位秒级的时间戳。
	LogisticsDispatcherName   string  `json:"logistics_dispatcher_name"`     //描述：美团配送骑手的姓名，取最新一次指派的骑手信息。
	LogisticsDispatcherMobile string  `json:"logistics_dispatcher_mobile"`   //描述：美团配送骑手的联系电话，取最新一次指派的骑手信息。
	PackageBagMoney           int     `json:"package_bag_money"`             //描述：订单维度的打包袋金额，单位是分。该信息默认不返回，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“29 打包袋”。（目前仅药品类和部分零售类商家支持打包袋功能） 注意，打包袋费用推送规则为： (1)订单打包袋如由美团提供给商家，费用结算美团平台，则该参数值为0； 注：订阅字段仅是支持推送字段；若开发者再单独向开放平台申请另加白名单支持按应用维度继续推送费用具体金额，则即使费用是结算给美团的，该参数也会推送实际收费金额。 (2)订单打包袋如由商家自己提供，费用结算给商家，则该参数为实际收费金额。
	EstimateArrivalTime       int     `json:"estimate_arrival_time"`         //描述：订单预计送达时间，为10位秒级的时间戳。该字段默认不返回，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“35 订单预计送达时间”。关于订单预计送达时间字段的说明： (1)当用户选择订单立即送出(delivery_time=0)，estimate_arrival_time字段信息则为美团计算的该即时订单预计送达时间。 (2)当用户选择订单在某个特定时间送达，即为预订单。estimate_arrival_time字段与delivery_time字段信息相同，均为用户选择的订单预计送达时间。 (3)当订单为用户到店自取方式，estimate_arrival_time字段与delivery_time字段信息相同，均为用户选择的到店取货时间。
	PackageBagMoneyYuan       string  `json:"package_bag_money_yuan"`        //描述：订单维度的打包袋金额，单位是元。该信息默认不推送，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“39 打包袋(元)”。 与package_bag_money字段相比，package_bag_money_yuan字段的单位为元，其他规则相同；订阅其中一个字段即可。 费用推送规则： (1)订单打包袋如由美团提供给商家，费用结算美团平台，则该参数值为0； 注：订阅字段仅是支持推送字段；若开发者再单独向开放平台申请另加白名单支持按应用维度继续推送费用具体金额，则即使费用是结算给美团的，该参数也会推送实际收费金额。 (2)订单打包袋如由商家自己提供，费用结算给商家，则该参数为实际收费金额。
	PoiReceiveDetailYuan      *string `json:"poi_receive_detail_yuan"`       //描述：订单维度的商家对账信息，json格式数据。该信息默认不推送，商家如有需求可在开发者中心->基础设置->订单订阅字段 页面开启订阅字段“38 商家应收款详情(元)”。 与poi_receive_detail字段相比，poi_receive_detail_yuan字段里的金额类字段单位为元，其他规则相同；订阅其中一个字段即可。
	TotalWeight               int64   `json:"total_weight"`                  //描述：订单重量（该信息默认不返回，可在开发者中心订阅），单位为克/g
	IncmpCode                 int     `json:"incmp_code"`                    //描述：订单数据状态标记。当订单中部分字段的数据因内部交互异常或网络等原因延迟生成（超时），导致开发者当前获取的订单数据不完整，此时平台对订单数据缺失情况进行标记。如不完整，建议尝试重新查询。注意，平台仅对部分模块的数据完整性进行监察标记（参考incmp_modules字段）。参考值： -1：有数据降级 0：无数据降级
	IncmpModules              []int   `json:"incmp_modules"`                 //描述：有降级的数据模块的集合，参考值： 0：订单商品详情 1：订单优惠信息 2：商品优惠详情 3：订单用户会员信息 4：订单维度的商家对账信息 5：订单维度的商家对账信息(元) 6：订单收货人地址 7：订单配送方式 8：开放平台用户id 9：部分退款商品信息 10：退货退款物流信息 11：部分订单基本信息(包括订单优惠信息、订单商品详情、门店信息等) 12：sku信息 13：spu信息 14：商品信息(可能是sku或spu等商品相关信息获取时降级) 15：替换折扣价为原价
	MedicareDetail            string  `json:"medicare_detail"`               //描述：医保报销详情
	PatientCollectionInfo     string  `json:"patient_collection_info"`       //描述：疫情登记信息
	IsWholeCityShip           int     `json:"is_whole_city_ship"`            //描述：是否是全城送订单：1-是，0-否
}
