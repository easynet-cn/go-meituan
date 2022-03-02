package order

//自配送商家同步发货状态和配送信息（推荐）
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/ecommerce/order/logistics/sync
//文档地址：https://open-shangou.meituan.com/home/docDetail/355
type EcommerceOrderLogisticsSyncRequest struct {
	OrderId               string `url:"order_id"`                //是否必须：是；描述：订单号（同订单展示ID）
	ThirdCarrierOrderId   string `url:"third_carrier_order_id"`  //是否必须：是；描述：第三方配送商物流单号
	CourierName           string `url:"courier_name"`            //是否必须：是；描述：配送员姓名。 姓名属于隐私信息，在用户端订单详情的“配送信息”中默认展示为“商家/第三方骑手”。 在商家端后台会展示骑手真实姓名。
	CourierPhone          string `url:"courier_phone"`           //是否必须：是；描述：配送员手机号码，此字段信息会同步展示在用户端订单详情的“配送信息”中。
	LogisticsProviderCode string `url:"logistics_provider_code"` //是否必须：是；描述：配送此订单商品的物流平台。参考取值： 10001-顺丰, 10002-达达, 10003-闪送, 10004-蜂鸟, 10005 UU跑腿,10006 快跑者, 10007 极客快送,10008-点我达,10009 同达, 10010-生活半径,10011 邻趣,10012 趣送, 10013 快服务 10014 菜鸟新配盟 10015 商家自建配送 10016 风先生,10017-其他,10032-美团跑腿。
	LogisticsStatus       int    `url:"logistics_status"`        //是否必须：是；描述：配送状态code，如下提供配送状态枚举值，以及各配送状态对应在C端（用户端）和B端（商家PC端）后台展示的配送状态信息。 未同步配送状态时（C端：商家已接单；B端：待发配送） 0-配送单发往配送（C端：商家已接单；B端：待骑手接单） 1-已创建配送包裹（C端：商家已接单；B端：待骑手接单） 5-已分配骑手（C端：商家已接单；B端：已分配骑手） 10-骑手已接单（C端：骑手正在赶往商家；B端：待骑手取货） 15-骑手已到店（C端：骑手到店取货中；B端：骑手已到店） 20-骑手已取货（C端：商品配送中/骑手正在送货；B端：骑手已取货） 40-骑手已送达（C端：商品已送达；B端：骑手已送达） 100-配送单已取消（C端：商家已接单；B端：配送已取消） 注：若同步配送状态为“配送单已取消”，接口仍支持继续同步配送状态。 说明：商家如未上传此信息，则平台默认值为20。
	Latitude              string `url:"latitude"`                //是否必须：是；描述：骑手当前的纬度，美团使用的是高德坐标系。
	Longitude             string `url:"longitude"`               //是否必须：是；描述：骑手当前的经度，美团使用的是高德坐标系。
}
