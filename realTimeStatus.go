package epsolar

import (
	"encoding/json"
	"fmt"
)

type RealTimeStatus struct {
	OverTemperatureInsideTheDevice *bool
	Night                          *bool
	BatteryStatus                  *BatteryStatusDetails
	ChargingEquipmentStatus        *ChargingEquipmentStatusDetails
	DischargingEquipmentStatus     *DischargingEquipmentStatusDetails
}

// ----

type BatteryStatus uint16

type BatteryStatusDetails struct {
	Raw                                uint16
	VoltageStatus                      VoltageStatus
	TemperatureStatus                  TemperatureStatus
	BatteryInternalResistanceAbnormal  bool
	WrongIdentificationForRatedVoltage bool
}

func (v BatteryStatus) Details() BatteryStatusDetails {
	return BatteryStatusDetails{
		Raw:                                uint16(v),
		VoltageStatus:                      VoltageStatus(getBits(uint16(v), 0, 0x0f)),
		TemperatureStatus:                  TemperatureStatus(getBits(uint16(v), 8, 0x0f)),
		BatteryInternalResistanceAbnormal:  checkBit(uint16(v), 8),
		WrongIdentificationForRatedVoltage: checkBit(uint16(v), 15),
	}
}

// ----

type VoltageStatus uint8

const (
	VoltageStatusNormal VoltageStatus = iota
	VoltageStatusOverVoltage
	VoltageStatusUnderVoltage
	VoltageStatusOverDischarge
	VoltageStatusFault
)

func (v VoltageStatus) String() string {
	switch v {
	case VoltageStatusNormal:
		return "Normal"
	case VoltageStatusOverVoltage:
		return "Over Voltage"
	case VoltageStatusUnderVoltage:
		return "Under Voltage"
	case VoltageStatusOverDischarge:
		return "Over Discharge"
	case VoltageStatusFault:
		return "Fault"
	default:
		return fmt.Sprintf("0x%02x", uint8(v))
	}
}

func (v VoltageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

// ----

type TemperatureStatus uint8

const (
	TemperatureStatusNormal TemperatureStatus = iota
	TemperatureStatusOverTemp
	TemperatureStatusLowTemp
)

func (v TemperatureStatus) String() string {
	switch v {
	case TemperatureStatusNormal:
		return "Normal"
	case TemperatureStatusOverTemp:
		return "Over Temp"
	case TemperatureStatusLowTemp:
		return "Low Temp"
	default:
		return fmt.Sprintf("0x%02x", uint8(v))
	}
}

func (v TemperatureStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

// ----

type ChargingEquipmentStatus uint16

type ChargingEquipmentStatusDetails struct {
	Raw                                uint16
	Running                            bool
	Fault                              bool
	ChargingStatus                     ChargingStatus
	PVInputIsShort                     bool
	LoadMOSFETIsShort                  bool
	LoadIsShort                        bool
	LoadIsOverCurrent                  bool
	InputIsOverCurrent                 bool
	AntiReverseMOSFETIsShort           bool
	ChargingOrAntiReverseMOSFETIsShort bool
	ChargingMOSFETIsShort              bool
	InputVoltageStatus                 InputVoltageStatus
}

func (v ChargingEquipmentStatus) Details() ChargingEquipmentStatusDetails {
	return ChargingEquipmentStatusDetails{
		Raw:                                uint16(v),
		Running:                            checkBit(uint16(v), 0),
		Fault:                              checkBit(uint16(v), 1),
		ChargingStatus:                     ChargingStatus(getBits(uint16(v), 2, 0x03)),
		PVInputIsShort:                     checkBit(uint16(v), 4),
		LoadMOSFETIsShort:                  checkBit(uint16(v), 7),
		LoadIsShort:                        checkBit(uint16(v), 8),
		LoadIsOverCurrent:                  checkBit(uint16(v), 9),
		InputIsOverCurrent:                 checkBit(uint16(v), 10),
		AntiReverseMOSFETIsShort:           checkBit(uint16(v), 11),
		ChargingOrAntiReverseMOSFETIsShort: checkBit(uint16(v), 12),
		ChargingMOSFETIsShort:              checkBit(uint16(v), 13),
		InputVoltageStatus:                 InputVoltageStatus(getBits(uint16(v), 14, 0x03)),
	}
}

// ----

type ChargingStatus uint8

const (
	ChargingStatusNoCharging ChargingStatus = iota
	ChargingStatusFloat
	ChargingStatusBoost
	ChargingStatusEqualization
)

func (v ChargingStatus) String() string {
	switch v {
	case ChargingStatusNoCharging:
		return "No Charging"
	case ChargingStatusFloat:
		return "Float"
	case ChargingStatusBoost:
		return "Boost"
	case ChargingStatusEqualization:
		return "Equalization"
	default:
		return fmt.Sprintf("0x%02x", uint8(v))
	}
}

func (v ChargingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

// ----

type InputVoltageStatus uint8

const (
	InputVoltageStatusNormal InputVoltageStatus = iota
	InputVoltageStatusNoInputPowerConnected
	InputVoltageStatusHigherInputVoltage
	InputVoltageStatusInputVoltageError
)

func (v InputVoltageStatus) String() string {
	switch v {
	case InputVoltageStatusNormal:
		return "Normal"
	case InputVoltageStatusNoInputPowerConnected:
		return "No Input Power Connected"
	case InputVoltageStatusHigherInputVoltage:
		return "Higher Input Voltage"
	case InputVoltageStatusInputVoltageError:
		return "Input Voltage Error"
	default:
		return fmt.Sprintf("0x%02x", uint8(v))
	}
}

func (v InputVoltageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

// ----

type DischargingEquipmentStatus uint16

type DischargingEquipmentStatusDetails struct {
	Raw                           uint16
	Running                       bool
	Fault                         bool
	OutputOverVoltage             bool
	BoostOverVoltage              bool
	ShortCircuitInHighVoltageSide bool
	InputOverVoltage              bool
	OutputVoltageAbnormal         bool
	UnableToStopDischarging       bool
	UnableToDischarge             bool
	ShortCircuit                  bool
	OutputPowerStatus             OutputPowerStatus
	InputVoltageStatus            DischargingEquipmentInputVoltageStatus
}

func (v DischargingEquipmentStatus) Details() DischargingEquipmentStatusDetails {
	return DischargingEquipmentStatusDetails{
		Raw:                           uint16(v),
		Running:                       checkBit(uint16(v), 0),
		Fault:                         checkBit(uint16(v), 1),
		OutputOverVoltage:             checkBit(uint16(v), 4),
		BoostOverVoltage:              checkBit(uint16(v), 5),
		ShortCircuitInHighVoltageSide: checkBit(uint16(v), 6),
		InputOverVoltage:              checkBit(uint16(v), 7),
		OutputVoltageAbnormal:         checkBit(uint16(v), 8),
		UnableToStopDischarging:       checkBit(uint16(v), 9),
		UnableToDischarge:             checkBit(uint16(v), 10),
		ShortCircuit:                  checkBit(uint16(v), 11),
		OutputPowerStatus:             OutputPowerStatus(getBits(uint16(v), 12, 0x03)),
		InputVoltageStatus:            DischargingEquipmentInputVoltageStatus(getBits(uint16(v), 14, 0x03)),
	}
}

// ----

type OutputPowerStatus uint8

const (
	OutputPowerStatusLightLoad OutputPowerStatus = iota
	OutputPowerStatusModerateLoad
	OutputPowerStatusRatedLoad
	OutputPowerStatusOverload
)

func (v OutputPowerStatus) String() string {
	switch v {
	case OutputPowerStatusLightLoad:
		return "Light Load"
	case OutputPowerStatusModerateLoad:
		return "Moderate Load"
	case OutputPowerStatusRatedLoad:
		return "Rated Load"
	case OutputPowerStatusOverload:
		return "Overload"
	default:
		return fmt.Sprintf("0x%02x", uint8(v))
	}
}

func (v OutputPowerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

// ----

type DischargingEquipmentInputVoltageStatus uint8

const (
	DischargingEquipmentInputVoltageStatusNormal DischargingEquipmentInputVoltageStatus = iota
	DischargingEquipmentInputVoltageStatusInputVoltageLow
	DischargingEquipmentInputVoltageStatusInputVoltageHigh
	DischargingEquipmentInputVoltageStatusNoAccess
)

func (v DischargingEquipmentInputVoltageStatus) String() string {
	switch v {
	case DischargingEquipmentInputVoltageStatusNormal:
		return "Normal"
	case DischargingEquipmentInputVoltageStatusInputVoltageLow:
		return "Input Voltage Low"
	case DischargingEquipmentInputVoltageStatusInputVoltageHigh:
		return "Input Voltage High"
	case DischargingEquipmentInputVoltageStatusNoAccess:
		return "No Access"
	default:
		return fmt.Sprintf("0x%02x", uint8(v))
	}
}

func (v DischargingEquipmentInputVoltageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}
