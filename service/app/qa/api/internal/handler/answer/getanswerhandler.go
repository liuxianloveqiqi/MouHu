package handler

import (
	"MouHu/service/app/qa/api/internal/logic/answer"
	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"
	"MouHu/service/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetAnswerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAnswerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := answer.NewGetAnswerLogic(r.Context(), svcCtx)
		resp, err := l.GetAnswer(&req)
		response.Response(w, resp, err) //②

	}
}
