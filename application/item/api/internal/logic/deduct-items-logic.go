package logic

import (
	"bufio"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/api/internal/model"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"
	utils "hmall/application/item/api/internal/util"
	"hmall/pkg/util"
	"io"
	"os"
	"strconv"
	"sync"
)

type DeductItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeductItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductItemsLogic {
	return &DeductItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Order struct {
	ItemId int
	Num    int
}

func (l *DeductItemsLogic) DeductItems(req *types.DeductItemsReq) error {
	// 1、校验参数
	if len(req.Order) == 0 {
		return nil
	}

	// 2、用于同步es
	pusherSearch := utils.NewPusherSearchLogic(l.ctx, l.svcCtx)

	_, err := mr.MapReduce[Order, map[int]*model.ItemDTO, int](func(source chan<- Order) {
		//解析参数
		for _, val := range req.Order {
			source <- Order{
				ItemId: val.ItemId,
				Num:    val.Num,
			}
		}
	}, func(order Order, writer mr.Writer[map[int]*model.ItemDTO], cancel func(error)) {
		//2、扣减库存
		newItem, err := l.svcCtx.ItemModel.DecutStock(l.ctx, order.ItemId, order.Num)
		if err != nil {
			logx.Errorf("ItemModel.DecutStock: id=%v,num=%v, error: %v", order.ItemId, order.Num, err)
			cancel(err)
		}
		writer.Write(map[int]*model.ItemDTO{order.Num: newItem})
	}, func(pipe <-chan map[int]*model.ItemDTO, writer mr.Writer[int], cancel func(error)) {
		//3、同步缓存
		wg := &sync.WaitGroup{}
		for stck_newItem := range pipe {
			var stock int
			var newItem *model.ItemDTO
			for k, v := range stck_newItem {
				stock = k
				newItem = v
			}
			//更新缓存
			wg.Add(2)
			threading.GoSafe(func() {
				defer wg.Done()
				itemId := strconv.FormatInt(newItem.Id, 10)
				lockKey := util.CacheKey(types.CacheStockLock, itemId)
				key := util.CacheKey(types.CacheItemKey, itemId)
				exists, _ := l.svcCtx.BizRedis.Exists(key)
				//缓存存在，更新缓存
				if exists {
					//读取lua脚本
					luaScript, err := ReadLua()
					if err != nil {
						panic(err)
					}
					// 执行Lua脚本
					_, err = l.svcCtx.BizRedis.Eval(luaScript, []string{
						lockKey,
						key,
						types.CacheItemStock,
					}, itemId, strconv.Itoa(stock), types.CacheItemTime)
					if err != nil && err.Error() != "redis: nil" {
						logx.Errorf("BizRedis.Eval, error: %v", err)
						cancel(err)
					}
				}
			})

			//同步es
			threading.GoSafe(func() {
				defer wg.Done()
				err := pusherSearch.PusherSearch(types.KqUpdate, newItem)
				if err != nil {
					logx.Errorf("pusherSearch.PusherSearch: %v, error: %v", *newItem, err)
					cancel(err)
				}
			})
		}
		wg.Wait()
		//没有结果，随便输出
		writer.Write(-1)
	})
	if err != nil {
		return err
	}
	//4、返回
	return nil
}

func ReadLua() (string, error) {
	//打开脚本
	file, err := os.Open(types.Luapath)
	if err != nil {
		logx.Errorf("os.Open: %v,error: %v", types.Luapath, err)
		return "", err
	}
	reader := bufio.NewReader(file)
	var luaScript string
	//逐行读取
	for {
		line, err := reader.ReadString('\n')
		luaScript += line
		if err == io.EOF {
			break
		}
	}

	return luaScript, nil
}
