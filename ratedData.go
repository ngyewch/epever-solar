package epsolar

type RatedData struct {
	ArrayRatedVoltage       *float64 // V
	ArrayRatedCurrent       *float64 // A
	ArrayRatedPower         *float64 // W
	BatteryRatedVoltage     *float64 // V
	BatteryRatedCurrent     *float64 // A
	BatteryRatedPower       *float64 // W
	LoadRatedVoltage        *float64 // V
	LoadRatedCurrent        *float64 // A
	LoadRatedPower          *float64 // W
	BatteryRealRatedVoltage *float64 // V
}
