package types

// 缓存相关
const (
	CacheItemKey    = "cache#item"
	CacheItemTime   = 3600
	CacheItemFields = "fields"
	CacheItemStock  = "stock"
	CacheItemStatus = "status"
	CacheStockLock  = "stockLock"
	Luapath         = "./etc/decut_stock.lua"
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

// es同步方式
const (
	KqDel    = iota //删除
	KqUpdate        //更新
)

// Kq补偿类型
const (
	KqCacheAll = iota
	KqCachePart
	KqCacheStock
	KqCacheStatus
	KqCacheDel
)

// 商品状态
const (
	ItemStatusNormal = iota + 1
	ItemStatusRemove
	ItemStatusDeleted
)

// 商品查询忽略字段
const SelOmit = "create_time,update_time,creater,updater,status"
