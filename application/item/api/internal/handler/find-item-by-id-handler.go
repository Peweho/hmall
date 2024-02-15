package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hmall/application/item/api/internal/logic"
	"hmall/application/item/api/internal/svc"
	"hmall/application/item/api/internal/types"
)

func FindItemByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FindItemByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFindItemByIdLogic(r.Context(), svcCtx)
		resp, err := l.FindItemById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
