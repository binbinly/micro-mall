package constvar

import "time"

// 用来定义一些公共的常量
const (
	// DefaultLimit 默认分页数
	DefaultLimit = 20

	// MaxID 最大id
	MaxID = 0xffffffffffff

	// CacheExpireTime 缓存过期时间
	CacheExpireTime = time.Hour * 24

	//HotSearchKey 搜索热词
	HotSearchKey = "search_hot"
)

// GetPageOffset 获取分页起始偏移量
func GetPageOffset(page int32) int {
	var offset int32
	if page > 0 {
		offset = (page - 1) * DefaultLimit
	}
	return int(offset)
}
