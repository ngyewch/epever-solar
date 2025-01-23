package epsolar

type RatedData struct {
	ChargingEquipmentRatedInputVoltage  float64 // V
	ChargingEquipmentRatedInputCurrent  float64 // A
	ChargingEquipmentRatedInputPower    float64 // W
	ChargingEquipmentRatedOutputVoltage float64 // V
	ChargingEquipmentRatedOutputCurrent float64 // A
	ChargingEquipmentRatedOutputPower   float64 // W
	ChargingMode                        uint16
	RatedOutputCurrentOfLoad            float64 // A
}
