package epsolar

import (
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
	regs3005, err := dev.mc.ReadRegisters(0x3005, 1, modbus.INPUT_REGISTER)
	if err != nil {
		return RatedData{}, err
	}
	regs300e, err := dev.mc.ReadRegisters(0x300e, 1, modbus.INPUT_REGISTER)
	if err != nil {
		return RatedData{}, err
	}
	regs311d, err := dev.mc.ReadRegisters(0x311d, 1, modbus.INPUT_REGISTER)
	if err != nil {
		return RatedData{}, err
	}
	return RatedData{
		RatedChargingCurrent:    convert16BitRegister(regs3005[0x00], 100), // 0x3005 Array
		RatedLoadCurrent:        convert16BitRegister(regs300e[0x00], 100), // 0x300e DC Load
		BatteryRealRatedVoltage: convert16BitRegister(regs311d[0x00], 100), // 0x311d Battery
	}, nil
}

func (dev *Dev) ReadParameters() (Parameters, error) {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	err := dev.requestSetup()
	if err != nil {
		return Parameters{}, err
	}
	regs9000, err := dev.mc.ReadRegisters(0x9000, 15, modbus.HOLDING_REGISTER)
	if err != nil {
		return Parameters{}, err
	}
	regs9067, err := dev.mc.ReadRegisters(0x9067, 1, modbus.HOLDING_REGISTER)
	if err != nil {
		return Parameters{}, err
	}
	regs906a, err := dev.mc.ReadRegisters(0x906a, 5, modbus.HOLDING_REGISTER)
	if err != nil {
		return Parameters{}, err
	}
	regs9070, err := dev.mc.ReadRegisters(0x9070, 1, modbus.HOLDING_REGISTER)
	if err != nil {
		return Parameters{}, err
	}
	regs9107, err := dev.mc.ReadRegisters(0x9107, 1, modbus.HOLDING_REGISTER)
	if err != nil {
		return Parameters{}, err
	}
	return Parameters{
		BatteryType:                                    BatteryType(regs9000[0x00]),                                    // 0x9000
		BatteryCapacity:                                convert16BitRegister(regs9000[0x01], 1),                        // 0x9001
		TemperatureCompensationCoefficient:             convert16BitRegister(regs9000[0x02], 100),                      // 0x9002
		OverVoltageDisconnectVoltage:                   convert16BitRegister(regs9000[0x03], 100),                      // 0x9003
		ChargingLimitVoltage:                           convert16BitRegister(regs9000[0x04], 100),                      // 0x9004
		OverVoltageReconnectVoltage:                    convert16BitRegister(regs9000[0x05], 100),                      // 0x9005
		EqualizeChargingVoltage:                        convert16BitRegister(regs9000[0x06], 100),                      // 0x9006
		BoostChargingVoltage:                           convert16BitRegister(regs9000[0x07], 100),                      // 0x9007
		FloatChargingVoltage:                           convert16BitRegister(regs9000[0x08], 100),                      // 0x9008
		BoostReconnectChargingVoltage:                  convert16BitRegister(regs9000[0x09], 100),                      // 0x9009
		LowVoltageReconnectVoltage:                     convert16BitRegister(regs9000[0x0A], 100),                      // 0x900a
		UnderVoltageWarningRecoverVoltage:              convert16BitRegister(regs9000[0x0B], 100),                      // 0x900b
		UnderVoltageWarningVoltage:                     convert16BitRegister(regs9000[0x0C], 100),                      // 0x900c
		LowVoltageDisconnectVoltage:                    convert16BitRegister(regs9000[0x0D], 100),                      // 0x900d
		DischargingLimitVoltage:                        convert16BitRegister(regs9000[0x0E], 100),                      // 0x900e
		BatteryRatedVoltageLevel:                       BatteryRatedVoltageLevel(regs9067[0x00]),                       // 0x9067
		DefaultLoadOnOffInManualMode:                   regs906a[0x00],                                                 // 0x906a
		EqualizeDuration:                               regs906a[0x01],                                                 // 0x906b
		BoostDuration:                                  regs906a[0x02],                                                 // 0x906c
		BatteryDischarge:                               convert16BitRegister(regs906a[0x03], 100),                      // 0x906d
		BatteryCharge:                                  convert16BitRegister(regs906a[0x04], 100),                      // 0x906e
		ChargingMode:                                   ChargingMode(regs9070[0x00]),                                   // 0x9070
		LiBatteryProtectionAndOverTemperatureDropPower: LiBatteryProtectionAndOverTemperatureDropPower(regs9107[0x00]), // 0x9107
	}, nil
}

func (dev *Dev) ReadRealTimeData() (RealTimeData, error) {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	err := dev.requestSetup()
	if err != nil {
		return RealTimeData{}, err
	}
	regs3100, err := dev.mc.ReadRegisters(0x3100, 4, modbus.INPUT_REGISTER)
	if err != nil {
		return RealTimeData{}, err
	}
	regs310c, err := dev.mc.ReadRegisters(0x310c, 6, modbus.INPUT_REGISTER)
	if err != nil {
		return RealTimeData{}, err
	}
	regs311a, err := dev.mc.ReadRegisters(0x311a, 1, modbus.INPUT_REGISTER)
	if err != nil {
		return RealTimeData{}, err
	}
	regs331a, err := dev.mc.ReadRegisters(0x331a, 3, modbus.INPUT_REGISTER)
	if err != nil {
		return RealTimeData{}, err
	}
	return RealTimeData{
		PVArrayInputVoltage: convert16BitRegister(regs3100[0x00], 100),                 // 0x3100        Array
		PVArrayInputCurrent: convert16BitRegister(regs3100[0x01], 100),                 // 0x3101        Array
		PVArrayInputPower:   convert32BitRegister(regs3100[0x03], regs3100[0x02], 100), // 0x3102-0x3103 Array
		LoadVoltage:         convert16BitRegister(regs310c[0x00], 100),                 // 0x310c        DC Load
		LoadCurrent:         convert16BitRegister(regs310c[0x01], 100),                 // 0x310d        DC Load
		LoadPower:           convert32BitRegister(regs310c[0x03], regs310c[0x02], 100), // 0x310e-0x310f DC Load
		BatteryTemperature:  convert16BitRegister(regs310c[0x04], 100),                 // 0x3110        Battery
		DeviceTemperature:   convert16BitRegister(regs310c[0x05], 100),                 // 0x3111        Device
		BatterySOC:          convert16BitRegister(regs311a[0x00], 1),                   // 0x311a        Battery
		BatteryVoltage:      convert16BitRegister(regs331a[0x00], 100),                 // 0x331a        Battery
		BatteryCurrent:      convert32BitRegister(regs331a[0x02], regs331a[0x01], 100), // 0x331b-0x331c Battery
	}, nil
}

func (dev *Dev) ReadRealTimeStatus() (RealTimeStatus, error) {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	err := dev.requestSetup()
	if err != nil {
		return RealTimeStatus{}, err
	}
	regs3200, err := dev.mc.ReadRegisters(0x3200, 3, modbus.INPUT_REGISTER)
	if err != nil {
		return RealTimeStatus{}, err
	}
	return RealTimeStatus{
		BatteryStatus:              BatteryStatus(regs3200[0x00]),              // 0x3200 Battery
		ChargingEquipmentStatus:    ChargingEquipmentStatus(regs3200[0x01]),    // 0x3201 Array
		DischargingEquipmentStatus: DischargingEquipmentStatus(regs3200[0x02]), // 0x3202 Load
	}, nil
}

func (dev *Dev) ReadStatistics() (Statistics, error) {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	err := dev.requestSetup()
	if err != nil {
		return Statistics{}, err
	}
	regs3300, err := dev.mc.ReadRegisters(0x3300, 20, modbus.INPUT_REGISTER)
	if err != nil {
		return Statistics{}, err
	}
	return Statistics{
		MaximumInputVoltageToday:   convert16BitRegister(regs3300[0x00], 100),                 // 0x3300        Array (undocumented)
		MinimumInputVoltageToday:   convert16BitRegister(regs3300[0x01], 100),                 // 0x3301        Array (undocumented)
		MaximumBatteryVoltageToday: convert16BitRegister(regs3300[0x02], 100),                 // 0x3302        Battery
		MinimumBatteryVoltageToday: convert16BitRegister(regs3300[0x03], 100),                 // 0x3303        Battery
		ConsumedEnergyToday:        convert32BitRegister(regs3300[0x05], regs3300[0x04], 100), // 0x3304-0x3305 Consumed
		ConsumedEnergyThisMonth:    convert32BitRegister(regs3300[0x07], regs3300[0x06], 100), // 0x3306-0x3307 Consumed
		ConsumedEnergyThisYear:     convert32BitRegister(regs3300[0x09], regs3300[0x08], 100), // 0x3308-0x3309 Consumed
		TotalConsumedEnergy:        convert32BitRegister(regs3300[0x0b], regs3300[0x0a], 100), // 0x330a-0x330b Consumed
		GeneratedEnergyToday:       convert32BitRegister(regs3300[0x0d], regs3300[0x0c], 100), // 0x330c-0x330d Generated
		GeneratedEnergyThisMonth:   convert32BitRegister(regs3300[0x0f], regs3300[0x0e], 100), // 0x330e-0x330f Generated
		GeneratedEnergyThisYear:    convert32BitRegister(regs3300[0x11], regs3300[0x10], 100), // 0x3310-0x3311 Generated
		TotalGeneratedEnergy:       convert32BitRegister(regs3300[0x13], regs3300[0x12], 100), // 0x3312-0x3313 Generated
	}, nil
}
