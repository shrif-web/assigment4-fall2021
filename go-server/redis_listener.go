package go_server

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type RedisListener struct {
	client   *redis.Client
	receiver IMessageReceiver
	topic    string
}

func (l *RedisListener) Start(ctx context.Context) error {
	sub := l.client.Subscribe(ctx, l.topic)
	for ctx.Err() == nil {
		message, err := sub.ReceiveMessage(ctx)
		if err != nil {
			return err
		}
		err = l.receiver.Receive(ctx, Message{
			Payload: []byte(message.Payload),
			Topic:   l.topic,
		}, func() error {
			return nil
		})
		if err != nil {
			return err
		}
	}
	return sub.Close()
}

func (l *RedisListener) OnReceive(ctx context.Context, callBack func(message Message) error) error {
	sub := l.client.Subscribe(ctx, l.topic)
	for ctx.Err() == nil {
		message, err := sub.ReceiveMessage(ctx)
		if err != nil {
			return err
		}
		err = callBack(Message{
			Payload: []byte(message.Payload),
			Topic:   message.Channel,
		})
		if err != nil {
			return err
		}
	}
	return sub.Close()
}

func (l *RedisListener) ByChannel(ctx context.Context) (chan<- Message, error) {
	result := make(chan Message, 1)
	sub := l.client.Subscribe(ctx, l.topic)
	var err error
	go func() {
		defer func() {
			close(result)
			if err := sub.Close(); err != nil {
				logrus.Errorln(err)
			}
		}()
		for ctx.Err() == nil {
			message, err1 := sub.ReceiveMessage(ctx)
			if err1 != nil {
				err = err1
				return
			}
			result <- Message{
				Payload: []byte(message.Payload),
				Topic:   message.Channel,
			}
		}
	}()
	return result, err
}
