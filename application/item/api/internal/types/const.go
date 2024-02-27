package types

const (
	CacheItemKey      = "cache#item"
	CacheItemStockKey = "cache#item#stock"
	CacheOrderTime    = 3600
)

// 分页查询默认值
const (
	SortBy   = "id"
	IsAsc    = "desc"
	Page     = 1
	PageSize = 5
)

// 同步es和缓存 忽略查询item的字段
var Field = []string{"create_time", "update_time", "creater", "updater"}

const (
	KqDel    = iota //删除
	KqUpdate        //更新
)
