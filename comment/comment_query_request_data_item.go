package comment

type CommentQueryRequestDataItem struct {
	AddComment           string `json:"add_comment"`            //描述：用户追加评论的内容
	AddCommentTime       string `json:"add_comment_time"`       //描述：用户追加评论的时间，单位天。
	CommentTime          string `json:"comment_time"`           //描述：用户首次评价的时间，单位天。
	CommentContent       string `json:"comment_content"`        //描述：用户首次评价的内容
	CommentLables        string `json:"comment_lables"`         //描述：评价用户对配送的评价标签，多个标签以英文逗号隔开。标签
	CommentId            int64  `json:"comment_id"`             //描述：评价id，用户如追加评价，评价id不变。
	DeliveryCommentScore int    `json:"delivery_comment_score"` //描述：配送评价分数，满分为5分。对应前端的“配送”评价。
	FoodCommentScore     int    `json:"food_comment_score"`     //描述：商家评价分数，满分为5分。对应前端的“总体”评价。
	OrderCommentScore    int    `json:"order_comment_score"`    //描述：质量评价分数，满分为5分。对应前端的“质量”评价。
	PackingScore         int    `json:"packing_score"`          //描述：包装评价分数，满分为5分。对应前端的“包装”评价。
	ShipTime             int    `json:"ship_time"`              //描述：配送时间，单位是分钟。
	CommentPictures      string `json:"comment_pictures"`       //描述：用户提交的评价图片，多个图片url以英文逗号隔开。
	PraiseFoodList       string `json:"praise_food_list"`       //描述：用户点赞的商品，多个商品以英文逗号分隔。
	CriticFoodList       string `json:"critic_food_list"`       //描述：用户点踩的商品，多个商品以英文逗号分隔。
	ReplyStatus          int    `json:"reply_status"`           //描述：商家回复该评价的状态：0未回复，1已回复
	EReplyTime           string `json:"e_reply_time"`           //描述：商家回复时间，精确到天。
	EReplyContent        string `json:"e_reply_content"`        //描述：商家回复内容
	Result               string `json:"result"`
}
