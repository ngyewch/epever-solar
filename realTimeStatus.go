package epsolar

import (
	"encoding/json"
	"fmt"
)

type RealTimeStatus struct {
	BatteryStatus           BatteryStatus
	ChargingEquipmentStatus ChargingEquipmentStatus
}

// ----

type BatteryStatus uint16

type BatteryStatusDetails struct {
	VoltageStatus                      VoltageStatus
	TemperatureStatus                  TemperatureStatus
	BatteryInternalResistanceAbnormal  bool
	WrongIdentificationForRatedVoltage bool
}

func (bs BatteryStatus) Details() BatteryStatusDetails {
	return BatteryStatusDetails{
		VoltageStatus:                      VoltageStatus(getBits(uint16(bs), 0, 0x0f)),
		TemperatureStatus:                  TemperatureStatus(getBits(uint16(bs), 8, 0x0f)),
		BatteryInternalResistanceAbnormal:  checkBit(uint16(bs), 8),
		WrongIdentificationForRatedVoltage: checkBit(uint16(bs), 15),
	}
}

// ----

type VoltageStatus uint8

const (
	VoltageStatusNormal VoltageStatus = iota
	VoltageStatusOverVolt
	VoltageStatusUnderVolt
	VoltageStatusLowVoltDisconnect
	VoltageStatusFault
)

func (vs VoltageStatus) String() string {
	switch vs {
	case VoltageStatusNormal:
		return "Normal"
	case VoltageStatusOverVolt:
		return "Over Volt"
	case VoltageStatusUnderVolt:
		return "Under Volt"
	case VoltageStatusLowVoltDisconnect:
		return "Low Volt Disconnect"
	case VoltageStatusFault:
		return "Fault"
	default:
		return fmt.Sprintf("%02x", uint8(vs))
	}
}

func (vs VoltageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(vs.String())
}

// ----

type TemperatureStatus uint8

const (
	TemperatureStatusNormal TemperatureStatus = iota
	TemperatureStatusOverTemp
	TemperatureStatusLowTemp
)

func (ts TemperatureStatus) String() string {
	switch ts {
	case TemperatureStatusNormal:
		return "Normal"
	case TemperatureStatusOverTemp:
		return "Over Temp"
	case TemperatureStatusLowTemp:
		return "Low Temp"
	default:
		return fmt.Sprintf("%02x", uint8(ts))
	}
}

func (ts TemperatureStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(ts.String())
}

// ----

type ChargingEquipmentStatus uint16

type ChargingEquipmentStatusDetails struct {
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

func (ces ChargingEquipmentStatus) Details() ChargingEquipmentStatusDetails {
	return ChargingEquipmentStatusDetails{
		Running:                            checkBit(uint16(ces), 0),
		Fault:                              checkBit(uint16(ces), 1),
		ChargingStatus:                     ChargingStatus(getBits(uint16(ces), 2, 0x03)),
		PVInputIsShort:                     checkBit(uint16(ces), 4),
		LoadMOSFETIsShort:                  checkBit(uint16(ces), 7),
		LoadIsShort:                        checkBit(uint16(ces), 8),
		LoadIsOverCurrent:                  checkBit(uint16(ces), 9),
		InputIsOverCurrent:                 checkBit(uint16(ces), 10),
		AntiReverseMOSFETIsShort:           checkBit(uint16(ces), 11),
		ChargingOrAntiReverseMOSFETIsShort: checkBit(uint16(ces), 12),
		ChargingMOSFETIsShort:              checkBit(uint16(ces), 13),
		InputVoltageStatus:                 InputVoltageStatus(getBits(uint16(ces), 14, 0x03)),
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

func (cs ChargingStatus) String() string {
	switch cs {
	case ChargingStatusNoCharging:
		return "No Charging"
	case ChargingStatusFloat:
		return "Float"
	case ChargingStatusBoost:
		return "Boost"
	case ChargingStatusEqualization:
		return "Equalization"
	default:
		return fmt.Sprintf("%02x", uint8(cs))
	}
}

func (cs ChargingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(cs.String())
}

// ----

type InputVoltageStatus uint8

const (
	InputVoltageStatusNormal InputVoltageStatus = iota
	InputVoltageStatusNoPowerConnected
	InputVoltageStatusHigherVoltInput
	InputVoltageStatusInputVoltError
)

func (ivs InputVoltageStatus) String() string {
	switch ivs {
	case InputVoltageStatusNormal:
		return "Normal"
	case InputVoltageStatusNoPowerConnected:
		return "No Power Connected"
	case InputVoltageStatusHigherVoltInput:
		return "Higher Volt Input"
	case InputVoltageStatusInputVoltError:
		return "Input Volt Error"
	default:
		return fmt.Sprintf("%02x", uint8(ivs))
	}
}

func (ivs InputVoltageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(ivs.String())
}
