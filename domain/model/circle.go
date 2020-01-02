package model

// IrrigationCircle defines an irrigation circle. Every irrigation circle is associated with a relay,
// meaning that the number of the circle identifies the relay to be turned on / off.
type IrrigationCircle struct {
	// Number is the relay associated with this water pipe instance.
	Number uint8 `json:"number"`
	// Name is a name or short description of the particular irrigation circle.
	Name string `json:"name"`
}

// IrrigationCircleGroup defines a named group of irrigation circles.
// Circles in a group can be opened parallel, making irrigation less time consuming.
// Circle groups and individual circles are opened sequentially while an irrigation program is running.
// It is important to note that one circle can be member of only one group. This is validated when the
// configuration is applied.
type IrrigationCircleGroup struct {
	// Name is the name of the irrigation group.
	Name string `json:"name"`
	// Members specifies the irrigation group. Elements of the list are the Number fields of the IrrigationCircle structure.
	Members []uint8 `json:"members"`
}
