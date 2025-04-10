package epsolar

import (
	"encoding/binary"
	"errors"
	"github.com/simonvetter/modbus"
	"sync"
)

type Dev struct {
	mc     *modbus.ModbusClient
	unitId uint8
	mutex  *sync.Mutex
}

func New(mc *modbus.ModbusClient, unitId uint8, mutex *sync.Mutex) *Dev {
	return &Dev{
		mc:     mc,
		unitId: unitId,
		mutex:  mutex,
	}
}

func (dev *Dev) requestSetup() error {
	err := dev.mc.SetUnitId(dev.unitId)
	if err != nil {
		return err
	}
	err = dev.mc.SetEncoding(modbus.BIG_ENDIAN, modbus.LOW_WORD_FIRST)
	if err != nil {
		return err
	}
	return nil
}

func (dev *Dev) ReadRatedData() (RatedData, error) {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	err := dev.requestSetup()
	if err != nil {
		return RatedData{}, err
	}

	var r RatedData

	r.ArrayRatedVoltage, err = dev.readInputRegisterFromUint16ToFloat64(0x3000, 100)
	if err != nil {
		return RatedData{}, err
	}
	r.ArrayRatedCurrent, err = dev.readInputRegisterFromUint16ToFloat64(0x3001, 100)
	if err != nil {
		return RatedData{}, err
	}
	r.ArrayRatedPower, err = dev.readInputRegisterFromUint32ToFloat64(0x3002, 100)
	if err != nil {
		return RatedData{}, err
	}
	r.BatteryRatedVoltage, err = dev.readInputRegisterFromUint16ToFloat64(0x3004, 100)
	if err != nil {
		return RatedData{}, err
	}
	r.BatteryRatedCurrent, err = dev.readInputRegisterFromUint16ToFloat64(0x3005, 100)
	if err != nil {
		return RatedData{}, err
	}
	r.BatteryRatedPower, err = dev.readInputRegisterFromUint32ToFloat64(0x3006, 100)
	if err != nil {
		return RatedData{}, err
	}
	r.LoadRatedVoltage, err = dev.readInputRegisterFromUint16ToFloat64(0x300d, 100)
	if err != nil {
		return RatedData{}, err
	}
	r.LoadRatedCurrent, err = dev.readInputRegisterFromUint16ToFloat64(0x300e, 100)
	if err != nil {
		return RatedData{}, err
	}
	r.LoadRatedPower, err = dev.readInputRegisterFromUint32ToFloat64(0x300f, 100)
	if err != nil {
		return RatedData{}, err
	}
	r.BatteryRealRatedVoltage, err = dev.readInputRegisterFromUint16ToFloat64(0x311d, 100)
	if err != nil {
		return RatedData{}, err
	}

	return r, nil
}

func (dev *Dev) ReadParameters() (Parameters, error) {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	err := dev.requestSetup()
	if err != nil {
		return Parameters{}, err
	}

	var r Parameters

	{
		v, err := dev.mc.ReadRegister(0x9000, modbus.HOLDING_REGISTER)
		if err != nil {
			if !errors.Is(err, modbus.ErrIllegalDataAddress) {
				return Parameters{}, err
			}
		} else {
			v2 := BatteryType(v)
			r.BatteryType = &v2
		}
	}
	r.BatteryCapacity, err = dev.readHoldingRegisterFromUint16ToFloat64(0x9001, 1)
	if err != nil {
		return Parameters{}, err
	}
	r.TemperatureCompensationCoefficient, err = dev.readHoldingRegisterFromUint16ToFloat64(0x9002, 100)
	if err != nil {
		return Parameters{}, err
	}
	r.OverVoltageDisconnectVoltage, err = dev.readHoldingRegisterFromUint16ToFloat64(0x9003, 100)
	if err != nil {
		return Parameters{}, err
	}
	r.ChargingLimitVoltage, err = dev.readHoldingRegisterFromUint16ToFloat64(0x9004, 100)
	if err != nil {
		return Parameters{}, err
	}
	r.OverVoltageReconnectVoltage, err = dev.readHoldingRegisterFromUint16ToFloat64(0x9005, 100)
	if err != nil {
		return Parameters{}, err
	}
	r.EqualizeChargingVoltage, err = dev.readHoldingRegisterFromUint16ToFloat64(0x9006, 100)
	if err != nil {
		return Parameters{}, err
	}
	r.BoostChargingVoltage, err = dev.readHoldingRegisterFromUint16ToFloat64(0x9007, 100)
	if err != nil {
		return Parameters{}, err
	}
	r.FloatChargingVoltage, err = dev.readHoldingRegisterFromUint16ToFloat64(0x9008, 100)
	if err != nil {
		return Parameters{}, err
	}
	r.BoostReconnectChargingVoltage, err = dev.readHoldingRegisterFromUint16ToFloat64(0x9009, 100)
	if err != nil {
		return Parameters{}, err
	}
	r.LowVoltageReconnectVoltage, err = dev.readHoldingRegisterFromUint16ToFloat64(0x900a, 100)
	if err != nil {
		return Parameters{}, err
	}
	r.UnderVoltageWarningRecoverVoltage, err = dev.readHoldingRegisterFromUint16ToFloat64(0x900b, 100)
	if err != nil {
		return Parameters{}, err
	}
	r.UnderVoltageWarningVoltage, err = dev.readHoldingRegisterFromUint16ToFloat64(0x900c, 100)
	if err != nil {
		return Parameters{}, err
	}
	r.LowVoltageDisconnectVoltage, err = dev.readHoldingRegisterFromUint16ToFloat64(0x900d, 100)
	if err != nil {
		return Parameters{}, err
	}
	r.DischargingLimitVoltage, err = dev.readHoldingRegisterFromUint16ToFloat64(0x900e, 100)
	if err != nil {
		return Parameters{}, err
	}
	{
		v, err := dev.mc.ReadRegister(0x9067, modbus.HOLDING_REGISTER)
		if err != nil {
			if !errors.Is(err, modbus.ErrIllegalDataAddress) {
				return Parameters{}, err
			}
		} else {
			v2 := BatteryRatedVoltageLevel(v)
			r.BatteryRatedVoltageLevel = &v2
		}
	}
	r.DefaultLoadOnOffInManualMode, err = dev.readHoldingRegister(0x906a)
	if err != nil {
		return Parameters{}, err
	}
	r.EqualizeDuration, err = dev.readHoldingRegister(0x906b)
	if err != nil {
		return Parameters{}, err
	}
	r.BoostDuration, err = dev.readHoldingRegister(0x906c)
	if err != nil {
		return Parameters{}, err
	}
	r.BatteryDischarge, err = dev.readHoldingRegisterFromUint16ToFloat64(0x906d, 100) // NOTE: possibly incorrect documentation (divisor)
	if err != nil {
		return Parameters{}, err
	}
	r.BatteryCharge, err = dev.readHoldingRegisterFromUint16ToFloat64(0x906e, 100) // NOTE: possibly incorrect documentation (divisor)
	if err != nil {
		return Parameters{}, err
	}
	{
		v, err := dev.mc.ReadRegister(0x9070, modbus.HOLDING_REGISTER)
		if err != nil {
			if !errors.Is(err, modbus.ErrIllegalDataAddress) {
				return Parameters{}, err
			}
		} else {
			v2 := ChargingMode(v)
			r.ChargingMode = &v2
		}
	}
	{
		v, err := dev.mc.ReadRegister(0x9107, modbus.HOLDING_REGISTER)
		if err != nil {
			if !errors.Is(err, modbus.ErrIllegalDataAddress) {
				return Parameters{}, err
			}
		} else {
			v2 := LiBatteryProtectionAndOverTemperatureDropPower(v)
			details := v2.Details()
			r.LiBatteryProtectionAndOverTemperatureDropPower = &details
		}
	}

	return r, nil
}

func (dev *Dev) ReadRealTimeData() (RealTimeData, error) {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	err := dev.requestSetup()
	if err != nil {
		return RealTimeData{}, err
	}

	var r RealTimeData

	r.PVArrayInputVoltage, err = dev.readInputRegisterFromUint16ToFloat64(0x3100, 100)
	if err != nil {
		return RealTimeData{}, err
	}
	r.PVArrayInputCurrent, err = dev.readInputRegisterFromUint16ToFloat64(0x3101, 100)
	if err != nil {
		return RealTimeData{}, err
	}
	r.PVArrayInputPower, err = dev.readInputRegisterFromUint32ToFloat64(0x3102, 100)
	if err != nil {
		return RealTimeData{}, err
	}
	r.LoadVoltage, err = dev.readInputRegisterFromUint16ToFloat64(0x310c, 100)
	if err != nil {
		return RealTimeData{}, err
	}
	r.LoadCurrent, err = dev.readInputRegisterFromUint16ToFloat64(0x310d, 100)
	if err != nil {
		return RealTimeData{}, err
	}
	r.LoadPower, err = dev.readInputRegisterFromUint32ToFloat64(0x310e, 100)
	if err != nil {
		return RealTimeData{}, err
	}
	r.BatteryTemperature, err = dev.readInputRegisterFromUint16ToFloat64(0x3110, 100)
	if err != nil {
		return RealTimeData{}, err
	}
	r.DeviceTemperature, err = dev.readInputRegisterFromUint16ToFloat64(0x3111, 100)
	if err != nil {
		return RealTimeData{}, err
	}
	r.BatterySOC, err = dev.readInputRegisterFromUint16ToFloat64(0x311a, 1)
	if err != nil {
		return RealTimeData{}, err
	}
	r.BatteryVoltage, err = dev.readInputRegisterFromUint16ToFloat64(0x331a, 100)
	if err != nil {
		return RealTimeData{}, err
	}
	r.BatteryCurrent, err = dev.readInputRegisterFromUint32ToFloat64(0x331b, 100)
	if err != nil {
		return RealTimeData{}, err
	}

	return r, nil
}

func (dev *Dev) ReadRealTimeStatus() (RealTimeStatus, error) {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	err := dev.requestSetup()
	if err != nil {
		return RealTimeStatus{}, err
	}

	var r RealTimeStatus

	r.OverTemperatureInsideTheDevice, err = dev.readDiscreteInput(0x2000)
	if err != nil {
		return RealTimeStatus{}, err
	}
	r.Night, err = dev.readDiscreteInput(0x200c)
	if err != nil {
		return RealTimeStatus{}, err
	}
	{
		v, err := dev.mc.ReadRegister(0x3200, modbus.INPUT_REGISTER)
		if err != nil {
			if !errors.Is(err, modbus.ErrIllegalDataAddress) {
				return RealTimeStatus{}, err
			}
		} else {
			v2 := BatteryStatus(v)
			details := v2.Details()
			r.BatteryStatus = &details
		}
	}
	{
		v, err := dev.mc.ReadRegister(0x3201, modbus.INPUT_REGISTER)
		if err != nil {
			if !errors.Is(err, modbus.ErrIllegalDataAddress) {
				return RealTimeStatus{}, err
			}
		} else {
			v2 := ChargingEquipmentStatus(v)
			details := v2.Details()
			r.ChargingEquipmentStatus = &details
		}
	}
	{
		v, err := dev.mc.ReadRegister(0x3202, modbus.INPUT_REGISTER)
		if err != nil {
			if !errors.Is(err, modbus.ErrIllegalDataAddress) {
				return RealTimeStatus{}, err
			}
		} else {
			v2 := DischargingEquipmentStatus(v)
			details := v2.Details()
			r.DischargingEquipmentStatus = &details
		}
	}

	return r, nil
}

func (dev *Dev) ReadStatistics() (Statistics, error) {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	err := dev.requestSetup()
	if err != nil {
		return Statistics{}, err
	}

	var r Statistics

	r.MaximumArrayVoltageToday, err = dev.readInputRegisterFromUint16ToFloat64(0x3300, 100) // undocumented
	if err != nil {
		return Statistics{}, err
	}
	r.MinimumArrayVoltageToday, err = dev.readInputRegisterFromUint16ToFloat64(0x3301, 100) // undocumented
	if err != nil {
		return Statistics{}, err
	}
	r.MaximumBatteryVoltageToday, err = dev.readInputRegisterFromUint16ToFloat64(0x3302, 100)
	if err != nil {
		return Statistics{}, err
	}
	r.MinimumBatteryVoltageToday, err = dev.readInputRegisterFromUint16ToFloat64(0x3303, 100)
	if err != nil {
		return Statistics{}, err
	}
	r.ConsumedEnergyToday, err = dev.readInputRegisterFromUint32ToFloat64(0x3304, 100)
	if err != nil {
		return Statistics{}, err
	}
	r.ConsumedEnergyThisMonth, err = dev.readInputRegisterFromUint32ToFloat64(0x3306, 100)
	if err != nil {
		return Statistics{}, err
	}
	r.ConsumedEnergyThisYear, err = dev.readInputRegisterFromUint32ToFloat64(0x3308, 100)
	if err != nil {
		return Statistics{}, err
	}
	r.TotalConsumedEnergy, err = dev.readInputRegisterFromUint32ToFloat64(0x330a, 100)
	if err != nil {
		return Statistics{}, err
	}
	r.GeneratedEnergyToday, err = dev.readInputRegisterFromUint32ToFloat64(0x330c, 100)
	if err != nil {
		return Statistics{}, err
	}
	r.GeneratedEnergyThisMonth, err = dev.readInputRegisterFromUint32ToFloat64(0x330e, 100)
	if err != nil {
		return Statistics{}, err
	}
	r.GeneratedEnergyThisYear, err = dev.readInputRegisterFromUint32ToFloat64(0x3310, 100)
	if err != nil {
		return Statistics{}, err
	}
	r.TotalGeneratedEnergy, err = dev.readInputRegisterFromUint32ToFloat64(0x3312, 100)
	if err != nil {
		return Statistics{}, err
	}

	return r, nil
}

func (dev *Dev) readInputRegisterFromUint16ToFloat64(addr uint16, divisor float64) (*float64, error) {
	v, err := dev.mc.ReadRegister(addr, modbus.INPUT_REGISTER)
	if err != nil {
		if errors.Is(err, modbus.ErrIllegalDataAddress) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	f64 := float64(v) / divisor
	return &f64, nil
}

func (dev *Dev) readInputRegisterFromUint32ToFloat64(addr uint16, divisor float64) (*float64, error) {
	v, err := dev.mc.ReadRegisters(addr, 2, modbus.INPUT_REGISTER)
	if err != nil {
		if errors.Is(err, modbus.ErrIllegalDataAddress) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	var b []byte
	b = binary.BigEndian.AppendUint16(b, v[1])
	b = binary.BigEndian.AppendUint16(b, v[0])
	f64 := float64(int32(binary.BigEndian.Uint32(b))) / divisor
	return &f64, nil
}

func (dev *Dev) readHoldingRegister(addr uint16) (*uint16, error) {
	v, err := dev.mc.ReadRegister(addr, modbus.HOLDING_REGISTER)
	if err != nil {
		if errors.Is(err, modbus.ErrIllegalDataAddress) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &v, nil
}

func (dev *Dev) readHoldingRegisterFromUint16ToFloat64(addr uint16, divisor float64) (*float64, error) {
	v, err := dev.mc.ReadRegister(addr, modbus.HOLDING_REGISTER)
	if err != nil {
		if errors.Is(err, modbus.ErrIllegalDataAddress) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	f64 := float64(v) / divisor
	return &f64, nil
}

func (dev *Dev) readDiscreteInput(addr uint16) (*bool, error) {
	v, err := dev.mc.ReadDiscreteInput(addr)
	if err != nil {
		if errors.Is(err, modbus.ErrIllegalDataAddress) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &v, nil
}
