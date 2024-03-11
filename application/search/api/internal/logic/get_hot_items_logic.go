package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"hmall/application/item/rpc/item"
	"strings"

	"hmall/application/search/api/internal/svc"
	"hmall/application/search/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHotItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotItemsLogic {
	return &GetHotItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type response struct {
	Took         int          `json:"took"`
	TimedOut     bool         `json:"timed_out"`
	Shards       shards       `json:"_shards"`
	Hits         hits         `json:"hits"`
	Aggregations aggregations `json:"aggregations"`
}

type shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type hits struct {
	Total    total         `json:"total"`
	MaxScore interface{}   `json:"max_score"`
	Hits     []interface{} `json:"hits"`
}

type total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type aggregations struct {
	ItemCount itemCount `json:"item_count"`
}

type itemCount struct {
	DocCountErrorUpperBound int      `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int      `json:"sum_other_doc_count"`
	Buckets                 []bucket `json:"buckets"`
}

type bucket struct {
	Key      string `json:"key"`
	DocCount int    `json:"doc_count"`
}

func (l *GetHotItemsLogic) GetHotItems(req *types.GetHotItemsReq) (resp *types.GetHotItemsResp, err error) {
	query := map[string]any{
		"aggs": map[string]any{
			"item_count": map[string]any{
				"terms": map[string]any{
					"field": "item_id",
					"size":  req.Num,
				},
			},
		},
	}

	marshal, err := json.Marshal(query)
	if err != nil {
		logx.Errorf("json.Marshal: %v, error: %v", query, err)
		return nil, err
	}
	//向es发起请求
	esResp, err := esapi.SearchRequest{
		Index: []string{types.EsUserSearchInfoIndex},
		Body:  bytes.NewReader(marshal),
	}.Do(context.Background(), l.svcCtx.Es)
	if err != nil {
		logx.Errorf("esapi.SearchRequest: %v, error: %v", string(marshal), err)
		return nil, err
	}

	//处理数据
	start := strings.Index(esResp.String(), "]")
	respStr := esResp.String()[start+1:]

	var res response
	err = json.Unmarshal([]byte(respStr), &res)
	if err != nil {
		logx.Errorf("json.Unmarshal: %v, error: %v", respStr, err)
		return nil, err
	}

	//构造商品id，查询商品
	ids := make([]string, 0, len(res.Aggregations.ItemCount.Buckets))
	for _, v := range res.Aggregations.ItemCount.Buckets {
		ids = append(ids, v.Key)
	}

	itemResp, err := l.svcCtx.ItemRPC.FindItemByIds(l.ctx, &item.FindItemByIdsReq{Ids: ids})
	if err != nil {
		logx.Errorf("jItemRPC.FindItemByIds: %v, error: %v", ids, err)
		return nil, err
	}

	items := make([]types.SearchItemDTO, 0, len(res.Aggregations.ItemCount.Buckets))
	for _, v := range itemResp.Data {
		items = append(items, types.SearchItemDTO{
			Brand:        v.Brand,
			Category:     v.Category,
			CommentCount: v.CommentCount,
			Id:           v.Id,
			Image:        v.Image,
			IsAD:         v.IsAD,
			Name:         v.Name,
			Price:        v.Price,
			Sold:         v.Sold,
			Spec:         v.Spec,
			Status:       v.Status,
			Stock:        v.Stock,
		})
	}

	return &types.GetHotItemsResp{
		Items: items,
	}, nil
}
