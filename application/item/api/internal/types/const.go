package types

// 缓存相关
const (
	CacheItemKey           = "cache#item"
	CacheItemTime          = 3600
	CacheItemFields        = "fields"
	CacheItemStock         = "stock"
	CacheItemStatus        = "status"
	CacheStockLock         = "stockLock"
	Luapath                = "./etc/decut_stock.lua"
	ItemBloomKey           = "itemBloom"
	CacheQueryItemKey      = "cache#query#item"       //分页查询缓存
	CacheQueryItemTotalKey = "cache#query#item#total" //分页查询缓存
	CacheQueryItemTime     = 1800                     //分页查询缓存持续时间
)

// 分页查询默认值
const (
	SortBy    = "id"
	IsAsc     = "desc"
	Page      = 1
	PageSize  = 5
	DataQuery = 10 //数据库一次性查询条数
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
