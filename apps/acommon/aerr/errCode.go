package aerr

import "net/http"

// 成功返回
const OK uint32 = http.StatusOK
const Unauthorized uint32 = http.StatusUnauthorized

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const (
	SERVER_COMMON_ERROR uint32 = iota + 100001
	REUQEST_PARAM_ERROR
	TOKEN_EXPIRE_ERROR
	TOKEN_GENERATE_ERROR
	DB_ERROR
	DB_UPDATE_AFFECTED_ZERO_ERROR
	REDIS_ERROR
)

//用户模块
