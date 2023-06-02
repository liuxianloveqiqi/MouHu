package handler

import (
	"MouHu/service/app/qa/api/internal/logic/answer"
	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/common/response"
	"net/http"
)

func GetAnswerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := answer.NewGetAnswerLogic(r.Context(), svcCtx)
		resp, err := l.GetAnswer()
		response.Response(w, resp, err) //②

	}
}
