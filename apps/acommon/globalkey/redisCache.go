package globalkey

/**
redis key except "model cache key"  in here,
but "model cache key" in model
*/

// Cache_UserTokenKey /** 用户登陆的token
const Cache_UserTokenKey = "user_token:%d"

// 用户角色id池,创角时使用 redis.SPop() 取出一个
const Cache_GenRoleId_UserIdPool = "GenRoleId:UserIdPool"
