package epsolar

type RealTimeData struct {
	PVArrayInputVoltage float64 // V
	PVArrayInputCurrent float64 // A
	PVArrayInputPower   float64 // W
	LoadVoltage         float64 // V
	LoadCurrent         float64 // A
	LoadPower           float64 // W
	BatteryTemperature  float64 // C
	DeviceTemperature   float64 // C
	BatterySOC          float64 // %
	BatteryVoltage      float64 // V
	BatteryCurrent      float64 // A
}
