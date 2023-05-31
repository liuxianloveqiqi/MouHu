package handler

import (
	"MouHu/service/app/qa/api/internal/logic/question"
	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"
	"MouHu/service/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func AltQuestionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AltQuestionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := question.NewAltQuestionLogic(r.Context(), svcCtx)
		err := l.AltQuestion(&req)
		response.Response(w, nil, err) //②

	}
}
