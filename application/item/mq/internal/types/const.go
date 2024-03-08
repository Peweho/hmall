package types

// 缓存相关
const (
	CacheItemKey = "cache#item"
)

// 商品缓存字段
const (
	CacheItemLockUils     = "lockUtil"
	CacheItemLockUilsTime = 5   //ms
	CacheItemDeadLine     = "0" //过期
	CacheItemOwner        = "owner"
	CacheItemFields       = "fields"
	CacheItemStock        = "stock"
	CacheItemStatus       = "status"
)
