package model

import (
	"time"
)

// CircleDuration specifies how long a circle (relay at the end) should be open / operate
// either in a program or over a day, depending on the chosen configuration.
type CircleDuration struct {
	// Circle is the number of the relay involved in the particular program.
	Circle uint8 `json:"circle"`
	// Duration is the duration the relay needs to be opened for in seconds.
	Duration uint `json:"duration"`
}

// IrrigationSchedule contains strictly those information which are necessary for the application to schedule
// an irrigation. This information is not provided directly by the user but it is calculated from the configuration
// provided by the user.
type IrrigationSchedule struct {
	// StartTime is the time when the irrigation should start
	StartTime time.Time
	// CircleDurations contains the time amount needed for particular irrigation circles.
	CircleDurations []CircleDuration
}
