package epsolar

type Statistics struct {
	MaximumInputVoltageToday   float64 // V
	MinimumInputVoltageToday   float64 // V
	MaximumBatteryVoltageToday float64 // V
	MinimumBatteryVoltageToday float64 // V
	ConsumedEnergyToday        float64 // kWh
	ConsumedEnergyThisMonth    float64 // kWh
	ConsumedEnergyThisYear     float64 // kWh
	TotalConsumedEnergy        float64 // kWh
	GeneratedEnergyToday       float64 // kWh
	GeneratedEnergyThisMonth   float64 // kWh
	GeneratedEnergyThisYear    float64 // kWh
	TotalGeneratedEnergy       float64 // kWh
	CarbonDioxideReduction     float64 // Ton
	BatteryCurrent             float64 // A
	BatteryTemperature         float64 // C
	AmbientTemperature         float64 // C
}
