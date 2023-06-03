package handler

import (
	"MouHu/service/app/like/api/internal/logic"
	"MouHu/service/app/like/api/internal/svc"
	"MouHu/service/app/like/api/internal/types"
	"MouHu/service/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func LikeAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LikeAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLikeAddLogic(r.Context(), svcCtx)
		err := l.LikeAdd(&req)
		response.Response(w, nil, err) //②

	}
}
