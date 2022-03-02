package order

type OrderGetOrderDetailResposneDataItemSgOpenThirdLogisticsInfosItem struct {
	SpCode   string `json:"sp_code"`    //描述：快递公司编码
	SpName   string `json:"sp_name"`    //描述：快递公司名称
	SpPkgNum string `json:"sp_pkg_num"` //描述：快递单号
	SpStatus int    `json:"sp_status"`  //描述：快递单配送状态（来自第三方轨迹）。配送状态值有：-1-未知，0-商家已发货，1-商家发货，10-揽件，20-在途，30-派件，40-签收，50-疑难，60-收件人拒签，70-退回，80-退签，100-待清关，110-清关中，120-清关异常，130-已清关，140-转投，150-待取件。
}
