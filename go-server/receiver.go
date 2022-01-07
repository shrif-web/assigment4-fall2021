package go_server

import (
	"context"
	"time"
)

type VehicleLogUnmarshaler struct {
	bus IEventDispatcher
}

func (v *VehicleLogUnmarshaler) Receive(ctx context.Context, message Message, ackFu AckFunc) error {
	event := VehicleLogReceivedEvent{
		OccurredAt: time.Now(),
	}
	vehicleLog := &VehicleLog{}
	if err := vehicleLog.UnmarshalBinary(message.Payload) ; err != nil {
		return err
	}
	event.Log = vehicleLog
	return  v.bus.Dispatch(ctx,[]interface{}{vehicleLog},ackFu)
}

