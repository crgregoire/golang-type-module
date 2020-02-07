package types

import "time"

//
// LambdaMessage models the message
// sent from the lambda function
//
type LambdaMessage struct {
	Command     string    `json:"command"`
	Environment string    `json:"environment"`
	Timestamp   time.Time `json:"timestamp"`
	Payload     Payload   `json:"payload"`
}

//
// Payload is used by all lambda functions
//
type Payload struct {
	Customer  PayloadCustomer  `json:"customer"`
	Dispenser PayloadDispenser `json:"dispenser"`
	Pod       PayloadPod       `json:"pod"`
}

//
// PayloadCustomer is used in payload
//
type PayloadCustomer struct {
	ID string `json:"id"`
}

//
// PayloadDispenser is used in payload
//
type PayloadDispenser struct {
	Name                      string `json:"name"`
	Serial                    string `json:"serial"`
	ControllerFirmwareVersion string `json:"controllerFirmwareVersion"`
	WifiFirmwareVersion       string `json:"wifiFirmwareVersion"`
	PcbFirmwareVersion        string `json:"pcbFirmwareVersion"`
}

//
// PayloadPod is used in payload
//
type PayloadPod struct {
	Barcode           string `json:"barcode"`
	Flags             int    `json:"flags"`
	Inserted          *bool  `json:"inserted"`
	ServingsRemaining int    `json:"servingsRemaining"`
}
