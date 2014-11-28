package enum

type UserStatus int

const (
	UserStatus_Wait     UserStatus = iota // 0 等待
	UserStatus_Approved                   // 1 已批准
	UserStatus_Disabled                   // 2 禁用
)

type PostStatus int

const (
	PostStatus_Draft     PostStatus = 1  // 1 草稿
	PostStatus_Approved  PostStatus = 2  // 2 已批准
	PostStatus_Published PostStatus = 4  // 4 发布
	PostStatus_Highlight PostStatus = 8  // 8 高亮
	PostStatus_Stick     PostStatus = 16 // 16 置顶
	PostStatus_Elite     PostStatus = 32 //32 精华
)

type PostType int

const (
	PostType_Unknow PostType = 0
	// PostType_Home        PostType = 1 // 首页
	PostType_Menu PostType = 2 // 菜单
	// PostType_Page        PostType = 3 // 页面
	PostType_Article PostType = 4 // 文章
	// PostType_ArticleList PostType = 5 // 文章列表
	// PostType_Photo PostType = 6 // 相片
	// PostType_Gallery     PostType = 7 // 相册列表
	PostType_Link PostType = 8 // 链接
	// PostType_Blog PostType = 9 // QA

)

type PostContentType int

const (
	PostContentType_Unknow PostContentType = 0
	PostContentType_Home   PostContentType = 1 // 首页
	// PostContentType_Menu               PostContentType = 2  // 菜单
	PostContentType_Page        PostContentType = 3  // 页面
	PostContentType_Article     PostContentType = 4  // 文章
	PostContentType_ArticleList PostContentType = 5  // 文章列表
	PostContentType_Photo       PostContentType = 6  // 相片
	PostContentType_Gallery     PostContentType = 7  // 相册列表
	PostContentType_QA          PostContentType = 9  // QA
	PostContentType_Blog        PostContentType = 10 // QA
	PostContentType_BlogList    PostContentType = 11 // QA

)

type PostSecurity int

const (
	PostSecurity_NotSet   PostSecurity = 1 // 未设置
	PostSecurity_Visibled PostSecurity = 2 // 可见的
	PostSecurity_Passowrd PostSecurity = 4 // 密码访问
	PostSecurity_Private  PostSecurity = 8 // 私有的

)

type Query int

const (
	Query_All       Query = -1
	Query_IsDeleted Query = 1
)

type RoleType int

const (
	RoleType_Subscriber  RoleType = 1 << iota // 1 订阅者
	RoleType_Contributor                      // 2 投稿者
	RoleType_Author                           // 4 作者
	RoleType_Editor                           // 8 编辑
)

func (s UserStatus) String() string {
	switch s {
	case UserStatus_Wait:
		return "待审核"
	case UserStatus_Approved:
		return "已批准"
	case UserStatus_Disabled:
		return "禁用"

	default:
		return "未知"
	}
}

func (r RoleType) String() string {
	switch r {
	case RoleType_Subscriber:
		return "订阅者"
	case RoleType_Contributor:
		return "投稿者"
	case RoleType_Author:
		return "作者"
	case RoleType_Editor:
		return "编辑"
	default:
		return "未知"
	}
}
