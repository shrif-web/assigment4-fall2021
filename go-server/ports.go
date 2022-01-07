package go_server

import (
	"context"
	"time"
)

type IVehicleLogWriteRepository interface {
	Save(ctx context.Context,logs []*VehicleLog) error
	SaveAsync(ctx context.Context,logs []VehicleLog, notifier func())
	IMigrator
}

type IVehicleLogReadRepository interface {
	Select(ctx context.Context,timeRange TimeRange,pager interface{}) ([]*VehicleLog,error)
}

type IFraudRulesGenericRepository interface {
	Save(ctx context.Context,rules []interface{}) error
	FindAll(ctx context.Context,pointerSlice interface{}) error
	IMigrator
}

type IFraudWriteRepository interface {
	Save(ctx context.Context,logs []*Fraud) error
	SaveAsync(ctx context.Context,logs []Fraud, errorHandler func())
	IMigrator
}

type IFraudReadRepository interface {
	FindByTimeRange(ctx context.Context,timeRange TimeRange,pager interface{}) ([]*Fraud,error)
}

type IMigrator interface {
	Setup(ctx context.Context)
}
// IMessageReceiver ,implement by VehicleLogUnmarshaler
type IMessageReceiver interface {
	Receive(ctx context.Context, message Message, ackFu AckFunc) error
}

type ISubscribable = func(ctx context.Context, events []interface{}) error

type IEventDispatcher interface {
	Dispatch(ctx context.Context, events []interface{}, ack AckFunc) error
	RegisterAtLeastOnce(name string, handler ISubscribable)
	RegisterAtMostOnce(name string, handler ISubscribable)
}

type Message struct {
	Payload []byte
	Topic   string
}

// https://www.rabbitmq.com/confirms.html
type AckFunc = func() error

type TimeRange struct {
	start time.Time
	end time.Time
}

func (r TimeRange) isValid() bool {
	return r.start.Unix() < r.end.Unix()
}
