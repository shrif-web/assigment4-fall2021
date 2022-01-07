package go_server

import "time"

type VehicleLogReceivedEvent struct {
	OccurredAt time.Time
	Log *VehicleLog
}
