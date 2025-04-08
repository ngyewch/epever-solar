package epsolar

import "fmt"

type Parameters struct {
	BatteryType                                    *BatteryType
	BatteryCapacity                                *float64 // AH
	TemperatureCompensationCoefficient             *float64 // mV / C / 2V
	OverVoltageDisconnectVoltage                   *float64 // V
	ChargingLimitVoltage                           *float64 // V
	OverVoltageReconnectVoltage                    *float64 // V
	EqualizeChargingVoltage                        *float64 // V
	BoostChargingVoltage                           *float64 // V
	FloatChargingVoltage                           *float64 // V
	BoostReconnectChargingVoltage                  *float64 // V
	LowVoltageReconnectVoltage                     *float64 // V
	UnderVoltageWarningRecoverVoltage              *float64 // V
	UnderVoltageWarningVoltage                     *float64 // V
	LowVoltageDisconnectVoltage                    *float64 // V
	DischargingLimitVoltage                        *float64 // V
	BatteryRatedVoltageLevel                       *BatteryRatedVoltageLevel
	DefaultLoadOnOffInManualMode                   *uint16
	EqualizeDuration                               *uint16  // minutes
	BoostDuration                                  *uint16  // minutes
	BatteryDischarge                               *float64 // %
	BatteryCharge                                  *float64 // %
	ChargingMode                                   *ChargingMode
	LiBatteryProtectionAndOverTemperatureDropPower *LiBatteryProtectionAndOverTemperatureDropPowerDetails
}

// ---

type BatteryType uint16

const (
	BatteryTypeUserDefined BatteryType = iota
	BatteryTypeSealed
	BatteryTypeGEL
	BatteryTypeFlooded
	BatteryTypeLiFePO4_4s
	BatteryTypeLiFePO4_8s
	BatteryTypeLiFePO4_15s
	BatteryTypeLiFePO4_16s
	BatteryTypeLiNiCoMnO2_3s
	BatteryTypeLiNiCoMnO2_6s
	BatteryTypeLiNiCoMnO2_7s
	BatteryTypeLiNiCoMnO2_13s
	BatteryTypeLiNiCoMnO2_14s
)

func (v BatteryType) String() string {
	switch v {
	case BatteryTypeUserDefined:
		return "User Defined"
	case BatteryTypeSealed:
		return "Sealed"
	case BatteryTypeGEL:
		return "GEL"
	case BatteryTypeFlooded:
		return "Flooded"
	case BatteryTypeLiFePO4_4s:
		return "LiFePO4 (4s)"
	case BatteryTypeLiFePO4_8s:
		return "LiFePO4 (8s)"
	case BatteryTypeLiFePO4_15s:
		return "LiFePO4 (15s)"
	case BatteryTypeLiFePO4_16s:
		return "LiFePO4 (16s)"
	case BatteryTypeLiNiCoMnO2_3s:
		return "Li(NiCoMn)O2 (3s)"
	case BatteryTypeLiNiCoMnO2_6s:
		return "Li(NiCoMn)O2 (6s)"
	case BatteryTypeLiNiCoMnO2_7s:
		return "Li(NiCoMn)O2 (7s)"
	case BatteryTypeLiNiCoMnO2_13s:
		return "Li(NiCoMn)O2 (13s)"
	case BatteryTypeLiNiCoMnO2_14s:
		return "Li(NiCoMn)O2 (14s)"
	default:
		return fmt.Sprintf("0x%02x", uint8(v))
	}
}

// ---

type BatteryRatedVoltageLevel uint16

const (
	BatteryRatedVoltageLevelAutoRecognize BatteryRatedVoltageLevel = iota
	BatteryRatedVoltageLevel12V
	BatteryRatedVoltageLevel24V
	BatteryRatedVoltageLevel36V
	BatteryRatedVoltageLevel48V
	BatteryRatedVoltageLevel60V
	BatteryRatedVoltageLevel110V
	BatteryRatedVoltageLevel120V
	BatteryRatedVoltageLevel220V
	BatteryRatedVoltageLevel240V
)

func (v BatteryRatedVoltageLevel) String() string {
	switch v {
	case BatteryRatedVoltageLevelAutoRecognize:
		return "Auto Recognize"
	case BatteryRatedVoltageLevel12V:
		return "12V"
	case BatteryRatedVoltageLevel24V:
		return "24V"
	case BatteryRatedVoltageLevel36V:
		return "36V"
	case BatteryRatedVoltageLevel48V:
		return "48V"
	case BatteryRatedVoltageLevel60V:
		return "60V"
	case BatteryRatedVoltageLevel110V:
		return "110V"
	case BatteryRatedVoltageLevel120V:
		return "120V"
	case BatteryRatedVoltageLevel220V:
		return "220V"
	case BatteryRatedVoltageLevel240V:
		return "240V"
	default:
		return fmt.Sprintf("0x%02x", uint8(v))
	}
}

// ---

type ChargingMode uint16

const (
	ChargingModeVoltageCompensation ChargingMode = iota
	ChargingModeSOC
)

func (v ChargingMode) String() string {
	switch v {
	case ChargingModeVoltageCompensation:
		return "Voltage Compensation"
	case ChargingModeSOC:
		return "SOC"
	default:
		return fmt.Sprintf("0x%02x", uint8(v))
	}
}

// ---

type LiBatteryProtectionAndOverTemperatureDropPower uint16

func (v LiBatteryProtectionAndOverTemperatureDropPower) Details() LiBatteryProtectionAndOverTemperatureDropPowerDetails {
	return LiBatteryProtectionAndOverTemperatureDropPowerDetails{
		Raw:                                    uint16(v),
		LowTemperatureProtectionForCharging:    checkBit(uint16(v), 8),
		LowTemperatureProtectionForDischarging: checkBit(uint16(v), 9),
		OverTemperatureDropPower:               checkBit(uint16(v), 11),
	}
}

type LiBatteryProtectionAndOverTemperatureDropPowerDetails struct {
	Raw                                    uint16
	LowTemperatureProtectionForCharging    bool
	LowTemperatureProtectionForDischarging bool
	OverTemperatureDropPower               bool
}
