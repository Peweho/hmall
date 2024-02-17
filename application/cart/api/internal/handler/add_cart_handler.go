package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hmall/application/cart/api/internal/logic"
	"hmall/application/cart/api/internal/svc"
	"hmall/application/cart/api/internal/types"
)

func AddCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCartReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddCartLogic(r.Context(), svcCtx)
		err := l.AddCart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
