package ctxdata

import (
	"context"
	"encoding/json"
	"zinx-zero/apps/acommon/aerr"

	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtAccountId get uid from ctx
var CtxKeyJwtAccountId = "jwtAccountId"

func GetAccountIdFromCtx(ctx context.Context) (int64, error) {
	var accountId int64
	if jsonAccountId, ok := ctx.Value(CtxKeyJwtAccountId).(json.Number); ok {
		if int64AccountId, err := jsonAccountId.Int64(); err == nil {
			accountId = int64AccountId
		} else {
			logx.WithContext(ctx).Errorf("GetAccountIdFromCtx err : %+v", err)
		}
	}
	var err error
	if accountId == 0 {
		err = aerr.NewErrCode(aerr.Unauthorized)
	}
	return accountId, err
}
