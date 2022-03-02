package comment

//根据门店id批量查询评价信息（新版）
//请求地址：https://waimaiopen.meituan.com/api/v1/comment/query
type CommentQueryRequest struct {
	AppPoiCode  string `url:"app_poi_code"` //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	StartTime   string `url:"start_time"`   //是否必须：是；描述：查询开始日期，单位天。开始日期不能大于结束日期；仅支持查询近30天内的评价信息。
	EndTime     string `url:"end_time"`     //是否必须：是；描述：查询结束日期，单位天。结束日期不能大于等于当前日期。
	Pageoffset  int    `url:"pageoffset"`   //是否必须：是；描述：分页查询的偏移量，表示从第几条评论开始查，0表示第一条。
	Pagesize    int    `url:"pagesize"`     //是否必须：是；描述：每页评价信息展示条数，最大不超过20。
	ReplyStatus int    `url:"replyStatus"`  //是否必须：是；描述：商家回复状态，取值范围：1全部；0未回复；1已回复。
}
