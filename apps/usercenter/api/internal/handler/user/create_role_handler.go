package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	"zinx-zero/apps/acommon/result"
	"zinx-zero/apps/usercenter/api/internal/logic/user"
	"zinx-zero/apps/usercenter/api/internal/svc"
	"zinx-zero/apps/usercenter/api/internal/types"
)

func CreateRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := validator.New(validator.WithRequiredStructEnabled()).
			StructCtx(r.Context(), &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewCreateRoleLogic(r.Context(), svcCtx)
		resp, err := l.CreateRole(&req)
		result.HttpResult(r, w, resp, err)
	}
}