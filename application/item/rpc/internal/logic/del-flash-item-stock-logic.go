package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"hmall/application/item/rpc/internal/svc"
	"hmall/application/item/rpc/pb"
	"hmall/application/item/rpc/types"
	"hmall/pkg/util"
	"log"
	"strconv"
	"time"
)

type DelFlashItemStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelFlashItemStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFlashItemStockLogic {
	return &DelFlashItemStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var luaFlashItemV1 = `
		-- KEYS 分别是 1:键值，2:库存字段，3:锁名，4:秒杀状态键 
		-- ARGV 分别是 1:扣减库存数量,2:秒杀未开始，3:秒杀开始，4:秒杀失败，5:秒杀成功，6:秒杀持续时间 ,7:重试次数
		local maxRetries = tonumber(ARGV[7])
		local acquired = 0
		local res = 0
		local slot = 0
		local retries = 0
		repeat
			redis.call("SET", KEYS[4] ,ARGV[2] ) --设置秒杀状态未开始
			acquired = redis.call("SETNX", KEYS[3], "1") -- 尝试加锁
			if acquired == 1 then
				redis.call("SET", KEYS[4],ARGV[3]) --设置秒杀状态已经开始
				redis.call("EXPIRE", KEYS[3], 3)  -- 设置锁的过期时间
				local stock = redis.call("HGET",KEYS[1],KEYS[2]) -- 获取库存
				if tonumber(stock) < tonumber(ARGV[1]) then --检查库存是否足够
					redis.call("SET", KEYS[4], ARGV[4]) --设置秒杀未扣减库存
				else
					res = 1 
					redis.call("HSET",KEYS[1],KEYS[2],stock - ARGV[1]) -- 扣减库存
					redis.call("SETEX", KEYS[4],ARGV[6], ARGV[5]) --设置秒杀成功库存
				end

				redis.call("DEL", KEYS[3]) --释放锁

			else -- 未获得锁
				retries = retries + 1
			end
		until acquired == 1 or retries > maxRetries

		return res
	`

var luaFlashItemV2 = `
	-- KEYS 分别是 1:键值，2:库存字段，3:锁名，4:秒杀状态键 
	-- ARGV 分别是 1:扣减库存数量,2:秒杀未开始，3:秒杀开始，4:秒杀失败，5:秒杀成功，6:秒杀持续时间 ,7:重试次数

	local res = 0
	redis.call("SET", KEYS[4],ARGV[3]) --设置秒杀状态已经开始
	redis.call("EXPIRE", KEYS[3], 3)  -- 设置锁的过期时间
	local stock = redis.call("HGET",KEYS[1],KEYS[2]) -- 获取库存

	if tonumber(stock) < tonumber(ARGV[1]) then --检查库存是否足够
		redis.call("SET", KEYS[4], ARGV[4]) --设置秒杀未扣减库存
	else
		res = 1 
		redis.call("HSET",KEYS[1],KEYS[2],stock - ARGV[1]) -- 扣减库存
		redis.call("SETEX", KEYS[4],ARGV[6], ARGV[5]) --设置秒杀成功库存
	end
	
	redis.call("DEL", KEYS[3]) --释放锁
	return res
`

// 秒杀商品服务
func (l *DelFlashItemStockLogic) DelFlashItemStock(in *pb.DelFlashItemStockReq) (*pb.DelFlashItemStockResp, error) {

	key := util.CacheKey(types.CacheItemKey, in.ItemId)
	lock := fmt.Sprintf("%s#%s", types.CacheStockLock, in.ItemId)
	//秒杀状态键
	falshStateKey := fmt.Sprintf("%s#%d#%s", types.CacheFlashStatus, in.Uid, in.ItemId)

	//休眠时间
	stime := l.svcCtx.Config.Timeout / types.CacheFlashReTry
	//记录重试次数
	retries := 0
	//设置秒杀状态
	err := l.svcCtx.BizRedis.Setex(falshStateKey, strconv.Itoa(types.FalshNotStart), int(in.Duration))
	if err != nil {
		logx.Errorf("BizRedis.Setex: %v, error: %v", falshStateKey, err)
		return nil, err
	}

	for {
		if retries >= types.CacheFlashReTry {
			return nil, status.Error(200, "达到最大重试次数，秒杀失败")
		}
		setnx, _ := l.svcCtx.BizRedis.Setnx(lock, "1")
		if setnx {
			//调用脚本
			scriptRes, err := l.svcCtx.BizRedis.Eval(luaFlashItemV2, []string{
				key,
				types.CacheItemStock,
				lock,
				falshStateKey,
			}, in.Num, types.FalshNotStart, types.FalshStart, types.FalshEndNotDecut, types.FalshEndDecut, in.Duration, types.CacheFlashReTry)
			if err != nil {
				log.Println("执行lua脚本失败")
				logx.Errorf("BizRedis.Eval: %v, error: %v", key, err)
				return nil, err
			}
			//结果处理
			switch scriptRes.(int64) {
			case types.FalshItemFail:
				log.Println("扣减库存失败")
				return nil, status.Error(200, "扣减库存失败")
			case types.FalshItemSuccess:
				log.Println("成功返回")
				return &pb.DelFlashItemStockResp{}, nil
			default:
				panic("unhandled default case")
			}
		} else {
			//未获得锁进行休眠然后重试
			time.Sleep(time.Duration(stime) * time.Millisecond)
			retries++
		}
	}
}
