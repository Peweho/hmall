package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"hmall/application/item/rpc/item"
	"hmall/application/order/api/internal/svc"
	"hmall/application/order/api/internal/types"
	"hmall/pkg/util"
	"log"
	"strconv"
	"time"
)

type CreateFlashOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFlashOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFlashOrderLogic {
	return &CreateFlashOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFlashOrderLogic) CreateFlashOrder(req *types.CreateFlashOrdeReq) error {
	uid, _ := util.GetUsr(l.ctx, types.JwtKey)
	//构建扣减库存请求
	itemDelFlashReq := item.DelFlashItemStockReq{
		Uid:      int64(uid),
		ItemId:   req.ItemId,
		Num:      int64(req.Num),
		Duration: int64(req.Duration),
	}

	chFlash := make(chan error, 1)
	threading.GoSafe(func() {
		log.Println("调用rpc")
		_, err := l.svcCtx.ItemRPC.DelFlashItemStock(l.ctx, &itemDelFlashReq)
		log.Println("调用rpc结束")
		chFlash <- err
	})

isExit: //是否退出
	for {
		select {
		case status := <-chFlash:
			if status != nil {
				//退出事务
				log.Println("退出事务1")
				break isExit
			} else {
				//创建订单
				cretaOrder()
				break isExit
			}
		case <-time.After(types.FalshTimeOut * time.Second):
			log.Println("超时检查")
			//超时检查
			key := fmt.Sprintf("%s#%d#%s", types.CacheFlashStatus, uid, req.ItemId)
			log.Println("key:", key)
			get, err := l.svcCtx.BizRedis.Get(key)
			log.Println("get:", get)
			if err != nil {
				logx.Errorf("BizRedis.Get: %v, error: %v", key, err)
				//TODo:处理宕机错误
				return err
			}
			status, _ := strconv.Atoi(get)

			//根据状态，进行下一步处理
			switch status {
			//未开始，退出事务
			case types.FalshNotStart:
				log.Println("退出事务2")
				break isExit
			//开始处理，继续等待
			case types.FalshStart:
				log.Println("循环等待")
			//扣减失败，退出事务
			case types.FalshEndNotDecut:
				log.Println("退出事务3")
				break isExit
			//已经扣减，创建订单
			case types.FalshEndDecut:
				cretaOrder()
				break isExit
			default:
				panic("unhandled default case")
			}
		}
	}
	return nil
}

func cretaOrder() {
	log.Println("开始创建订单")
}
