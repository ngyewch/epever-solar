package epsolar

type RealTimeData struct {
	ChargingEquipmentInputVoltage     float64 // V
	ChargingEquipmentInputCurrent     float64 // A
	ChargingEquipmentInputPower       float64 // W
	ChargingEquipmentOutputVoltage    float64 // V
	ChargingEquipmentOutputCurrent    float64 // A
	ChargingEquipmentOutputPower      float64 // W
	DischargingEquipmentOutputVoltage float64 // V
	DischargingEquipmentOutputCurrent float64 // A
	DischargingEquipmentOutputPower   float64 // W
	BatteryTemperature                float64 // C
	TemperatureInsideEquipment        float64 // C
	PowerComponentsTemperature        float64 // C
	BatterySOC                        float64 // %
	RemoteBatteryTemperature          float64 // C
	BatteryRealRatedPower             float64 // V
}
