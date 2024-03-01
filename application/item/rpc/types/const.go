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
	LuapathRollBack = "./etc/decut_stock_roll_back.lua"
)

const (
	KqDel    = iota //删除
	KqUpdate        //更新
)

// 商品查询忽略字段
const SelOmit = "create_time,update_time,creater,updater,status"

// 商品状态
const (
	ItemStatusNormal = iota + 1
	ItemStatusRemove
	ItemStatusDeleted
)
