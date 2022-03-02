package order

type EcommerceOrderGetOrderRefundDetailResponseDataItem struct {
	WmOrderIdView                     int64                                   `json:"wm_order_id_view"`                  //描述：订单展示ID，与用户端、商家端订单详情中展示的订单号码一致。数据库中请用bigint(20)存储此字段。
	OrderId                           int64                                   `json:"order_id"`                          //描述：订单号（同订单展示ID），数据库中请用bigint(20)存储此字段。
	RefundId                          int64                                   `json:"refund_id"`                         //描述：退款id，每次发起部分退款的退款id不同。
	Ctime                             int                                     `json:"ctime"`                             //描述：退款申请发起时间，为10位秒级的时间戳。
	Utime                             int                                     `json:"utime"`                             //描述：退款申请处理时间，为10位秒级的时间戳。如为商家主动发起的退款，退款申请的发起和处理时间相同；如用户申请后商家还未处理，此字段信息与ctime相同。
	RefundType                        int                                     `json:"refund_type"`                       //描述：退款类型：1-全额退款；2-部分退款；3-退差价。
	ResReason                         string                                  `json:"res_reason"`                        //描述：商家处理退款时答复的内容
	ResType                           int                                     `json:"res_type"`                          //描述：答复类型：0-未处理；1-商家驳回退款请求；2-商家同意退款；3-客服驳回退款请求；4-客服帮商家同意退款；5-超过3小时自动同意；6-系统自动确认；7-用户取消退款申请；8-用户取消退款申诉。
	ApplyType                         int                                     `json:"apply_type"`                        //描述：申请类型：0-订单取消自动确认退款； 1-用户申请退款； 2-客服帮用户申请退款； 3-重复提交而自动申请； 4-支付成功消息在订单取消之后到达而自动申请； 5-支付成功消息在订单被置为无效之后到达而自动申请； 6-用户被商家拒绝后申诉；7-商家申请退款。
	ApplyReason                       string                                  `json:"apply_reason"`                      //描述：申请退款的原因
	Money                             float64                                 `json:"money"`                             //描述：退款金额合计，单位是元，此字段信息为本次退款商品的总金额。如本次退款为全额退款，则此字段信息为订单原始在线实付金额减去已部分退款的金额后，当前剩余的订单总金额。
	Pictures                          string                                  `json:"pictures"`                          //描述：退款图片url的json格式数组，用户在申请退款时上传的退款图片，多张以英文逗号隔开，上限为9张。
	RefundPartialEstimateCharge       RefundPartialEstimateCharge             `json:"refund_partial_estimate_charge"`    //描述：结算预计费信息字段集合。如果在订单完成前用户取消了，则内部所有字段(如total_food_amount)均为null
	WmAppRetailForOrderPartRefundList []WmAppRetailForOrderPartRefundListItem `json:"wmAppRetailForOrderPartRefundList"` //描述：部分退款的商品明细，适用于按件部分退和按克重退差价两种类型。
	ServiceType                       string                                  `json:"service_type"`                      //描述：退款服务类型, 区分是否已开通退货退款售后业务。 未开通的场景： 0-退款流程或申诉流程 已开通场景： 1-仅退款流程 2-退款退货流程
	Status                            string                                  `json:"status"`                            //描述：推送当前售后单的状态类型，仅适用支持退货退款业务的商家： 1-已申请 10-初审已同意 11-初审已驳回 16-初审已申诉 17-初审申诉已同意 18-初审申诉已驳回 20-终审已发起（用户已发货） 21-终审已同意 22-终审已驳回 26-终审已申诉 27-终审申诉已同意 28-终审申诉已驳回 30-已取消
	ApplyOpUserType                   string                                  `json:"apply_op_user_type"`                //描述：推送当前退款或退货退款流程的发起方，仅适用于支持退货退款的商家。
	LogisticsInfo                     string                                  `json:"logistics_info"`                    //描述：物流信息集合，仅适用支持退货退款业务的品类；在用户提交物流信息以及商家终审的推送消息中展示。
	IncmpCode                         int                                     `json:"incmp_code"`                        //描述：订单数据状态标记。当订单中部分字段的数据因内部交互异常或网络等原因延迟生成（超时），导致开发者当前获取的订单数据不完整，此时平台对订单数据缺失情况进行标记。如不完整，建议尝试重新查询。注意，平台仅对部分模块的数据完整性进行监察标记（参考incmp_modules字段）。参考值： -1：有数据降级 0：无数据降级
	IncmpModules                      []int                                   `json:"incmp_modules"`                     //描述：有降级的数据模块的集合，参考值： 0：订单商品详情 1：订单优惠信息 2：商品优惠详情 3：订单用户会员信息 4：订单纬度的商家对账信息 5：订单纬度的商家对账信息(元) 6：订单收货人地址 7：订单配送方式 8：开放平台用户id 9：部分退款商品信息 10：退货退款物流信息 11：部分订单基本信息(包括订单优惠信息、订单商品详情、门店信息等) 12：sku信息 13：spu信息 14：商品信息(可能是sku或spu等商品相关信息获取时降级)
}
