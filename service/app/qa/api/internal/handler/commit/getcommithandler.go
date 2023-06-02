package handler

import (
	"MouHu/service/app/qa/api/internal/logic/commit"
	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/common/response"
	"net/http"
)

func GetCommitHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := commit.NewGetCommitLogic(r.Context(), svcCtx)
		resp, err := l.GetCommit()
		response.Response(w, resp, err) //②

	}
}
