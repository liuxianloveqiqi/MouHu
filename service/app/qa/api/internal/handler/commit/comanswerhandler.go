package handler

import (
	"MouHu/service/app/qa/api/internal/logic/commit"
	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"
	"MouHu/service/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ComAnswerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ComAnswerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := commit.NewComAnswerLogic(r.Context(), svcCtx)
		err := l.ComAnswer(&req)
		response.Response(w, nil, err) //②

	}
}
