package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hmall/application/cart/api/internal/logic"
	"hmall/application/cart/api/internal/svc"
)

func QueryCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewQueryCartLogic(r.Context(), svcCtx)
		resp, err := l.QueryCart()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
