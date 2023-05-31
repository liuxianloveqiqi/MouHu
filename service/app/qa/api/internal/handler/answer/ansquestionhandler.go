package handler

import (
	"MouHu/service/app/qa/api/internal/logic/answer"
	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"
	"MouHu/service/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func AnsQuestionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AnsQuestionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := answer.NewAnsQuestionLogic(r.Context(), svcCtx)
		err := l.AnsQuestion(&req)
		response.Response(w, nil, err) //②

	}
}
