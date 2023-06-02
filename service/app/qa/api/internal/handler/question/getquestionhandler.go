package handler

import (
	"MouHu/service/app/qa/api/internal/logic/question"
	"MouHu/service/app/qa/api/internal/svc"
	"MouHu/service/common/response"
	"net/http"
)

func GetQuestionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := question.NewGetQuestionLogic(r.Context(), svcCtx)
		resp, err := l.GetQuestion()
		response.Response(w, resp, err) //②

	}
}
