package middleware

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/spf13/cast"
	"math"
	"time"
)

const (
	// reason holds the error reason.
	reason            string = "UNAUTHORIZED"
	headerKeySign     string = "X-Sign"
	headerKeyNonce    string = "X-Nonce"
	headerKeyTime     string = "X-Timestamp"
	headerKeyDeviceId string = "X-DeviceId"
	SignatureFormat   string = "X-Sign=%s&X-Nonce=%s&X-Timestamp=%s&X-DeviceId=%s&%s"
)

var (
	ErrForbidden = errors.Forbidden(reason, "header is invalid")
	ErrUnknown   = errors.Unauthorized(reason, "server broken")
)

// VerifySign is a server middleware that recovers from any panics.
func VerifySign() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				sign := header.RequestHeader().Get(headerKeySign)
				nonce := header.RequestHeader().Get(headerKeyNonce)
				timestamp := header.RequestHeader().Get(headerKeyTime)
				deviceId := header.RequestHeader().Get(headerKeyDeviceId)

				if sign == "" || timestamp == "" || nonce == "" || deviceId == "" {
					return nil, ErrForbidden
				}
				// 请求时间与服务器时间误差不允许超过5分钟
				timeDf := cast.ToFloat64(time.Now().Unix() - cast.ToInt64(timestamp))
				if math.Abs(timeDf) > 300 {
					return nil, ErrForbidden
				}

				signBefore := fmt.Sprintf(SignatureFormat, sign, nonce, timestamp, deviceId, "")
				h := md5.New()
				h.Write([]byte(signBefore))
				checkSign := hex.EncodeToString(h.Sum(nil))
				if checkSign != sign {
					return nil, ErrForbidden
				}

				return handler(ctx, req)
			}

			return nil, ErrUnknown
		}
	}
}
