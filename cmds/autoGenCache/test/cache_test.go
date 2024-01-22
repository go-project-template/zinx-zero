package test

import (
	"autoGenCache/cachex"
	"autoGenCache/cachex/ZS"
	"context"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestZSViewPlayerInfo
//
//	@Description: set cache
func TestZSViewPlayerInfo(t *testing.T) {
	c := GetConf()
	testCases := []struct {
		name        string
		argUserId   int64
		updateValue *int64
		updateError error
		findRes     *int64
		findError   error
	}{
		{
			name:        "set user 1 fail by type",
			argUserId:   1,
			updateValue: proto.Int64(1),
			updateError: errors.New("set user 1 fail"),
			findRes:     proto.Int64(1),
			findError:   cachex.ErrNotHandleQuery,
		},
		{
			name:        "set user 1 success",
			argUserId:   1,
			updateValue: proto.Int64(1),
			updateError: nil,
			findRes:     proto.Int64(1),
			findError:   nil,
		},
		{
			name:        "set user 1 to 2",
			argUserId:   1,
			updateValue: proto.Int64(2),
			updateError: nil,
			findRes:     proto.Int64(2),
			findError:   nil,
		},
	}
	//init cache
	viewPlayerInfoCache := ZS.NewViewPlayerInfoCache(c.Cache)

	//clear cache
	for _, tc := range testCases {
		assert.NoError(t, viewPlayerInfoCache.Delete(context.Background(), tc.argUserId))
	}
	// begin
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			// 更新缓存操作
			{
				// 更新缓存
				_, err := viewPlayerInfoCache.UpdateFromFn(ctx, tc.argUserId, func() (*int64, error) {
					return tc.updateValue, tc.updateError
				})
				// 检查更新值和返回值是否一致
				if err != nil {
					assert.EqualError(t, err, tc.updateError.Error())
				}
			}
			{
				//// 查询缓存操作
				findOne, err := viewPlayerInfoCache.FindOne(ctx, tc.argUserId)
				if err == nil {
					// 检查数据
					assert.Equal(t, *findOne, *tc.findRes)
				} else {
					assert.EqualError(t, err, tc.findError.Error())
				}
			}
		})
	}
}
