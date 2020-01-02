package model

import (
	"time"
)

// IrrigationSystemConfiguration is the configuration provided by the user for the application.
// Configuration should be provided in JSON format and be unmarshalled by the application.
// There are optional fields in the configuration. Even though all the program fields are optional,
// at least one program must be specified. This is validated when the configuration is applied.
type IrrigationSystemConfiguration struct {
	// Location is the name of the city to be used for weather data query.
	// Data is optional, not needed when program is defined with explicit start times.
	Location string `json:"location,omitempty"`

	// Circles is the list of the irrigation circles.
	// Field is mandatory.
	Circles []IrrigationCircle `json:"circles"`

	// Irrigation groups in the configuration. Every member of an irrigation circle group are irrigated
	// parallel.
	// Field is optional.
	CircleGroups []IrrigationCircleGroup `json:"circlegroups,omitempty"`

	// StartTimeProgram contains the list of irrigation programs where all the configuration details are
	// specified manually.
	// Field is optional.
	StartTimeProgram []IrrigationProgramWithStartTime `json:"starttimeprogram,omitempty"`

	// StartCountProgram contains the list of irrigation programs where it is specified how many times the
	// irrigation program should be started, but the start time is calculated based on weather data.
	// Field is optional.
	StartCountProgram []IrrigationProgramWithStartCount `json:"startcountprogram,omitempty"`

	// DailyWaterProgram contains the list of daily water needs of different circles.
	// Based on that the application calculates the number of irrigations and the schedule.
	// Field is optional.
	DailyWaterProgram []IrrigationProgramWaterNeed `json:"dailywaterprogram,omitempty"`
}

// IrrigationProgramBase contains the common fields of an irrigation program.
type IrrigationProgramBase struct {
	// Name of the program.
	// Mandatory field.
	Name string `json:"name"`

	// ProgramRepeat defines on which days the program should be repeated.
	// Optional field, in case of absence the program is repeated every day.
	ProgramRepeat RepeatPattern `json:"repeat,omitempty"`

	// AdjustToWeather decides if the application should adjust the water amount for the circle
	// based on the season and weather forcast. Practically provide more water in hot summer days.
	// This requires the programs to be defined for the spring season.
	// Field is optional and defaults to false.
	AdjustToWeather bool `json:"adjusttoweather,omitempty"`
}

// IrrigationProgramWithStartTime contains an irrigation program completely specified by the user.
// By this program the schedule is not affected by weather forcast.
type IrrigationProgramWithStartTime struct {
	IrrigationProgramBase

	// Duration of the involved circles.
	// Mandatory field.
	CircleDurations []CircleDuration `json:"circledurations"`

	// StartTimes is the schedule of the irrigation program.
	// Mandatory field.
	StartTimes []time.Time `json:"starttimes"`
}

// IrrigationProgramWithStartCount contains an irrigation program where the user only specifies how many
// times the program should start. The start times are scheduled then based on the specified count, sun rise,
// sun set and the forcasted maximum temperature.
type IrrigationProgramWithStartCount struct {
	IrrigationProgramBase

	// Duration of the involved circles.
	// Mandatory field.
	CircleDurations []CircleDuration `json:"circledurations"`

	// DailyStartCount specifies how many times this program should run. Based on this number and weather forcast
	// the application counts the start time(s).
	// Mandatory field.
	DailyStartCount uint8 `json:"dailystartcount"`
}

// IrrigationProgramWaterNeed contains an irrigation program where the user only specifies the daily amount
// of the water needed for the specific circles (amount is defined in circle / relay opening time).
// The number of program starts and the schedule is calculated based on the water need and weather forcast.
type IrrigationProgramWaterNeed struct {
	IrrigationProgramBase

	// DailyWaterNeed defines the daily water need for watering circles. The amount is defined not in liter
	// but in seconds how long the circle (relay) should be open per day. Every parameter for scheduling programs
	// will be calculated base on this input, season and weather forcast.
	// Mandatory field.
	DailyWaterNeed []CircleDuration `json:"dailywaterneed"`
}

// RepeatPattern defines when and how the irrigation should be repeated.
// If this is not defined then irrigation is done on every day.
// All the fields in this structure are optional, but at least one of them is required to be defined and only
// one. This is validated.
type RepeatPattern struct {
	// WeekdayPattern defines a repeat pattern where it is defined on which days the irrigation needs to be
	// done.
	WeekdayPattern WeekDays `json:"weekdays,omitempty"`

	// MonthPattern defines a repeat pattern where it is defined on which days of the month the irrigation
	// needs to be done
	MonthPattern [31]bool `json:"monthdays,omitempty"`

	// NDayPattern defines a pattern the irrigation will be done on every n'th day.
	NDayPattern uint8 `json:"nday,omitempty"`
}

// WeekDays defines flags for each day on the week. That determines if irrigation is needed on that day or not.
type WeekDays struct {
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
	Sunday    bool `json:"sunday"`
}
