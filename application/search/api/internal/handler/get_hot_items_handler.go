package handler

import (
	"hmall/application/search/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hmall/application/search/api/internal/logic"
	"hmall/application/search/api/internal/svc"
)

func GetHotItemsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetHotItemsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetHotItemsLogic(r.Context(), svcCtx)
		resp, err := l.GetHotItems(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
