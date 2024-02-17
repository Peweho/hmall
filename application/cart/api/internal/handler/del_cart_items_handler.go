package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hmall/application/cart/api/internal/logic"
	"hmall/application/cart/api/internal/svc"
	"hmall/application/cart/api/internal/types"
)

func DelCartItemsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelCartItemsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDelCartItemsLogic(r.Context(), svcCtx)
		err := l.DelCartItems(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
