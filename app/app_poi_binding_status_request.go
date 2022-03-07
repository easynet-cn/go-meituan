package app

//推送门店绑定和解绑消息
//请求方式：POST
//文档地址：https://open-shangou.meituan.com/home/docDetail/439
type AppPoiBindingStatusRequest struct {
	OpType        int    `url:"op_type"`        //是否必须：是；描述：操作类型：1-绑定门店；2-解绑门店。
	PoiInfo       string `url:"poi_info"`       //是否必须：是；描述：门店相关信息
	OpTime        string `url:"op_time"`        //是否必须：是；描述：操作时间
	OpSource      string `url:"op_source"`      //是否必须：是；描述：当前操作来源：1.商家侧通过开放平台门店绑定授权工具操作，推送具体的商家账号信息。2.商家侧申请美团运营协助操作，推送“美团运营”。
	SettleRebind  string `url:"settle_rebind"`  //是否必须：否；描述：该字段仅适用于通过接口实现门店自入驻绑定冲突时更改app_poi_code重新绑定消息推送。
	AuthorizeCode string `url:"authorize_code"` //是否必须：否；描述：ISV类型的app适用的独有推送字段，当门店绑定至ISV的app时会推送门店授权码。 注意，针对第三方软件服务商申请的 ISV 类型的 app，需要确认前期是否已接入“绑定门店后OAuth相关数据推送”接口，该接口与本接口不可兼容。背景为“绑定门店后OAuth相关数据推送”更名为“推送门店绑定和解绑消息”，数据结构有修改，并且两种结构是共存的。所以，如之前对接的本推送与此文档的数据结构不一致，需要在代码改造后，提工单申请进行切换。
}
