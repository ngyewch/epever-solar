package epsolar

type Parameters struct {
	BatteryType                                    BatteryType
	BatteryCapacity                                float64 // AH
	TemperatureCompensationCoefficient             float64 // mV / C / 2V
	OverVoltageDisconnectVoltage                   float64 // V
	ChargingLimitVoltage                           float64 // V
	OverVoltageReconnectVoltage                    float64 // V
	EqualizeChargingVoltage                        float64 // V
	BoostChargingVoltage                           float64 // V
	FloatChargingVoltage                           float64 // V
	BoostReconnectChargingVoltage                  float64 // V
	LowVoltageReconnectVoltage                     float64 // V
	UnderVoltageWarningRecoverVoltage              float64 // V
	UnderVoltageWarningVoltage                     float64 // V
	LowVoltageDisconnectVoltage                    float64 // V
	DischargingLimitVoltage                        float64 // V
	BatteryRatedVoltageLevel                       BatteryRatedVoltageLevel
	DefaultLoadOnOffInManualMode                   uint16
	EqualizeDuration                               uint16  // minutes
	BoostDuration                                  uint16  // minutes
	BatteryDischarge                               float64 // %
	BatteryCharge                                  float64 // %
	ChargingMode                                   ChargingMode
	LiBatteryProtectionAndOverTemperatureDropPower LiBatteryProtectionAndOverTemperatureDropPower
}

// ---

type BatteryType uint16

// ---

type BatteryRatedVoltageLevel uint16

// ---

type ChargingMode uint16

// ---

type LiBatteryProtectionAndOverTemperatureDropPower uint16
