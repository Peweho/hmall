package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/zeromicro/go-zero/core/logx"
	"hmall/application/search/api/internal/svc"
	"hmall/application/search/api/internal/types"
	"log"
	"strings"
)

type SearchListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchListLogic {
	return &SearchListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchListLogic) SearchList(req *types.SearchListReq) (resp *types.SearchListResp, err error) {
	//Todo: 根据关键字，带有限定条件的聚合
	//Todo: 实现分页和排序
	/*分类：
	1、条件齐全；
	2、缺少关键字；
	3、缺少品牌；如果有key聚合品牌 should，
	4、缺少分类
	5、缺少价格
	*/
	//假设条件齐全
	query := map[string]any{
		"query": map[string]any{
			"bool": map[string]any{
				"must": []map[string]any{
					{"match": map[string]any{"name": req.Key}},
					{"match": map[string]any{"brand": req.Brand}},
				},
				"should": []map[string]any{},
				"filter": []map[string]any{
					{"term": map[string]any{"category": map[string]any{"value": req.Category}}},
					{"range": map[string]any{"price": map[string]any{"gte": req.MinPrice, "lte": req.MaxPrice}}},
				},
			},
		},
		"highlight": map[string]any{
			"fields": map[string]any{"name": map[string]any{"pre_tags": "<em>", "post_tags": "</em>"}},
		},
	}
	marshal, err := json.Marshal(query)
	if err != nil {
		logx.Errorf("json.Marshal: %v, error: %v", query, err)
		return nil, err
	}
	response, err := esapi.SearchRequest{
		Index: []string{types.EsItemsIndex},
		Body:  bytes.NewReader(marshal),
	}.Do(context.Background(), l.svcCtx.Es)
	if err != nil {
		logx.Errorf("esapi.SearchRequest: %v, error: %v", string(marshal), err)
		return nil, err
	}

	start := strings.Index(response.String(), "]")
	respStr := response.String()[start+1:]

	var res Response
	err = json.Unmarshal([]byte(respStr), &res)
	if err != nil {
		logx.Errorf("json.Unmarshal: %v, error: %v", respStr, err)
		return nil, err
	}
	log.Println(res)

	//构造返回对象
	//1、商品
	items := make([]types.SearchItemDTO, 0, len(res.Hits.HitDetails))
	for i, _ := range res.Hits.HitDetails {
		items = append(items, res.Hits.HitDetails[i].Source)
	}
	return &types.SearchListResp{
		Items: items,
	}, nil

	//item, err := l.svcCtx.ItemModel.InserItem(l.ctx)
	//for _, val := range *item {
	//	it := types.SearchItemDTO{
	//		Id:       val.Id,
	//		Category: val.Category,
	//		Status:   val.Status,
	//		Stock:    val.Stock,
	//		Spec:     val.Spec,
	//		Sold:     val.Sold,
	//		Name:     val.Name,
	//		Brand:    val.Brand,
	//	}
	//	marshal, err := json.Marshal(it)
	//	if err != nil {
	//		logx.Errorf("json.Marshal: %v, error: %v", it, err)
	//		return nil, err
	//	}
	//	_, err = esapi.IndexRequest{
	//		Index:      "items",
	//		DocumentID: strconv.Itoa(int(val.Id)),
	//		Body:       bytes.NewReader(marshal),
	//		Refresh:    "true",
	//	}.Do(context.Background(), l.svcCtx.Es)
	//	if err != nil {
	//		logx.Errorf("esapi.IndexRequest: %v, error: %v", it, err)
	//		return nil, err
	//	}
	//}
	//
	//return nil, nil
}

type Response struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Hits struct {
	Total      Total       `json:"total"`
	MaxScore   float64     `json:"max_score"`
	HitDetails []HitDetail `json:"hits"`
}

type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type HitDetail struct {
	Index     string              `json:"_index"`
	ID        string              `json:"_id"`
	Score     float64             `json:"_score"`
	Source    types.SearchItemDTO `json:"_source"`
	Highlight Highlight           `json:"highlight"`
}

type Highlight struct {
	Name []string `json:"name"`
}
