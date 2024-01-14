package middleware

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"google.golang.org/protobuf/proto"
)

// RouterProtoUnmarshal deserializing protobuf messages
func RouterProtoUnmarshal(request ziface.IRequest) {
	var data proto.Message
	// deserialization
	err := proto.Unmarshal(request.GetData(), data)
	if err != nil {
		zlog.Ins().ErrorF("MsgId:%d Handler proto unmarshal: err:%v data:%v",
			request.GetMsgID(), err, request.GetData())
		return
	}
	request.SetResponse(data)
}
