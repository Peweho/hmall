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
	ItemBloomKey    = "itemBloom"
	//秒杀状态，后拼接用户和商品id
	CacheFlashStatus = "item#flash"
	//秒杀重试次数
	CacheFlashReTry = 5
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

// 秒杀商品结果
const (
	FalshItemFail = iota
	FalshItemSuccess
)

// 秒杀商品状态
const (
	FalshNotStart = iota
	FalshStart
	FalshEndDecut
	FalshEndNotDecut
)
