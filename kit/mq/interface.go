package mq

import (
	"context"
)

type Producer interface {
	// 发布消息
	Publish(ctx context.Context, msg []byte) error
}

type Consumer interface {
	// 订阅消息
	Subscribe(handler func(ctx context.Context, msg []byte) error)
}
