package routes

import v1 "gin-blog-hufeng/api/v1"

// 后台接口
var (
	userAuthAPI v1.UserAuth
	blogInfoAPI v1.BlogInfo
	userAPI     v1.User
	categoryAPI v1.Category
	tagAPI      v1.Tag
	menuAPI     v1.Menu
)

// 前台接口