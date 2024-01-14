package ctxdata

import (
	"context"
	"encoding/json"
	"zinx-zero/apps/acommon/aerr"

	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) (int64, error) {
	var uid int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	var err error
	if uid == 0 {
		err = aerr.NewErrCode(aerr.Unauthorized)
	}
	return uid, err
}
