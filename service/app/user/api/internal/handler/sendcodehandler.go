package handler

import (
	"MouHu/service/app/user/api/internal/logic"
	"MouHu/service/app/user/api/internal/svc"
	"MouHu/service/app/user/api/internal/types"
	"MouHu/service/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func SendcodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterByPhoneRep
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSendcodeLogic(r.Context(), svcCtx)
		resp, err := l.Sendcode(&req)
		response.Response(w, resp, err) //②

	}
}
