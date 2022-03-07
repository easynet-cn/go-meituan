package poi

//已授权商家账号所关联门店的绑定与解绑
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/ecommerce/poi/bind/app/by/account
//文档地址：https://open-shangou.meituan.com/home/docDetail/481
type EcommercePoiBindAppByAccountRequest struct {
	Type       int    `url:"type"`         //是否必须：是；描述：操作类型。仅支持传枚举值： 1-授权绑定；2-解除关联；3-修改三方门店ID。
	WmPoiId    int64  `url:"wm_poi_id"`    //是否必须：是；描述：美团门店ID，该门店的商家账号必须已绑定当前应用的前提下，才能通过此接口完成门店相关操作。 若上传的门店是下线状态，不支持绑定应用，也不支持修改三方门店ID。 1.type=1（绑定）时，上传门店（非下线）如已绑定，报错。 2.type=2（解绑）时，上传门店（非下线）如未绑定，报错。 3.type=3（修改三方ID）时，上传门店（非下线）如未绑定，报错。 4.type=2（解绑）时，上传的wm_poi_id与app_poi_code之间如没有关联关系，则报错。
	AppPoiCode string `url:"app_poi_code"` //是否必须：是；描述：三方门店ID，与当前应用内已绑定门店的三方ID不能重复。
}
