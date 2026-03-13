package rabbitmqx

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/ve-weiyi/pkg/kit/mqx"
)

// producer RabbitMQ生产者实现
type producer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	config  *Config
}

// newProducer 创建生产者
func newProducer(conn *amqp.Connection, config *Config) (*producer, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	// 声明交换机
	err = channel.ExchangeDeclare(
		config.ExchangeName, // 交换机名称
		config.ExchangeType, // 交换机类型
		config.Durable,      // 是否持久化
		config.AutoDelete,   // 是否自动删除
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		channel.Close()
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	return &producer{
		conn:    conn,
		channel: channel,
		config:  config,
	}, nil
}

// Send 发送单条消息
func (p *producer) Send(ctx context.Context, message *mqx.Message) error {
	if message == nil {
		return fmt.Errorf("message cannot be nil")
	}

	// 构建AMQP消息
	publishing := amqp.Publishing{
		ContentType:  "application/json",
		Body:         message.Body,
		Timestamp:    message.Timestamp,
		MessageId:    message.ID,
		DeliveryMode: amqp.Persistent, // 持久化消息
	}

	// 转换Headers
	if message.Headers != nil {
		publishing.Headers = amqp.Table(message.Headers)
	}

	// 发布消息
	err := p.channel.PublishWithContext(
		ctx,
		p.config.ExchangeName, // 交换机
		message.Key,           // 路由键
		false,                 // mandatory
		false,                 // immediate
		publishing,
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

// SendBatch 批量发送消息
func (p *producer) SendBatch(ctx context.Context, messages []*mqx.Message) error {
	for _, message := range messages {
		if err := p.Send(ctx, message); err != nil {
			return err
		}
	}
	return nil
}

// Close 关闭生产者
func (p *producer) Close() error {
	if p.channel != nil {
		return p.channel.Close()
	}
	return nil
}

// 确保实现了接口
var _ mqx.Producer = (*producer)(nil)
