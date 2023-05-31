package handler

import (
	"MouHu/service/app/qa/api/internal/logic/AnsOrCom"
	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/app/qa/api/internal/types"
	"MouHu/service/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func AltAnswerOrCommitHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AltAnswerOrCommitReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := AnsOrCom.NewAltAnswerOrCommitLogic(r.Context(), svcCtx)
		err := l.AltAnswerOrCommit(&req)
		response.Response(w, nil, err) //②

	}
}
