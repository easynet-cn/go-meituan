package medicine

//查询门店药品分类列表请求
//请求方式：POST
//请求地址：https://waimaiopen.meituan.com/api/v1/medicineCat/save
//文档地址：https://open-shangou.meituan.com/home/docDetail/80
type MedicineCatRequest struct {
	AppPoiCode         string  `url:"app_poi_code"`         //是否必须：是；描述：APP方门店id，即商家中台系统里门店的编码。如商家在操作绑定门店至开放平台应用中时，未绑定三方门店id信息，则默认APP方门店id与美团门店id相同。
	CategoryCode       string  `url:"category_code"`        //是否必须：否；描述：一级分类code：(1)门店内一级、二级分类code都不允许重复，且与美团标库标准的21个分类code也不能重复。(2)美团标库标准分类code限定长度不超过6个字符；自定义分类的分类code长度暂无限制。(3)不允许在美团标准分类下创建二级分类。
	CategoryName       string  `url:"category_name"`        //是否必须：创建一级分类时必填；描述：一级分类名称：(1)门店内一级、二级分类名称都不允许重复，且与美团标库标准的21个分类名称也不能重复。(2)字段信息限定长度不超过8个字符。(3)不允许在美团标准分类下创建二级分类。
	Sequence           int     `url:"sequence"`             //是否必须：创建一级分类时必填；描述：一级分类在当前门店内的排序序号，数字越小，前端排位越靠前。
	SecondCategoryCode *string `url:"second_category_code"` //是否必须：否；描述：二级分类code：(1)门店内一级、二级分类code都不允许重复，且与美团标库标准的21个分类code也不能重复。(2)自定义分类的分类code长度暂无限制。(3)不允许在美团标准分类下创建二级分类。(4)允许美团标准分类被创建为二级分类。
	SecondCategoryName *string `url:"second_category_name"` //是否必须：创建二级分类时必填；描述：二级分类名称：(1)门店内一级、二级分类名称都不允许重复，且与美团标库标准的21个分类名称也不能重复。(2)字段信息限定长度不超过8个字符。(3)不允许在美团标准分类下创建二级分类。(4)允许美团标准分类被创建为二级分类。
	SecondSequence     *int    `url:"second_sequence"`      //是否必须：创建二级分类时必填；描述：二级分类在所属一级分类下的排序序号，数字越小，前端排位越靠前。
}
