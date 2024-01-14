package user

import (
	"net/http"

	"zinx-zero/apps/acommon/result"
	"zinx-zero/apps/usercenter/api/internal/logic/user"
	"zinx-zero/apps/usercenter/api/internal/svc"
	"zinx-zero/apps/usercenter/api/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := validator.New(validator.WithRequiredStructEnabled()).
			StructCtx(r.Context(), &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		result.HttpResult(r, w, resp, err)
	}
}
