package user

import (
	"net/http"

	"zinx-zero/apps/acommon/result"
	"zinx-zero/apps/usercenter/cmd/api/internal/logic/user"
	"zinx-zero/apps/usercenter/cmd/api/internal/svc"
	"zinx-zero/apps/usercenter/cmd/api/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func WxMiniAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WXMiniAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if err := validator.New(validator.WithRequiredStructEnabled()).
			StructCtx(r.Context(), &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewWxMiniAuthLogic(r.Context(), svcCtx)
		resp, err := l.WxMiniAuth(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
