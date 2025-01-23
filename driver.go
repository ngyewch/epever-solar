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
	regs, err := dev.mc.ReadRegisters(0x3000, 9, modbus.INPUT_REGISTER)
	if err != nil {
		return RatedData{}, err
	}
	regs2, err := dev.mc.ReadRegisters(0x300e, 1, modbus.INPUT_REGISTER)
	if err != nil {
		return RatedData{}, err
	}
	return RatedData{
		ChargingEquipmentRatedInputVoltage:  convert16BitRegister(regs[0x00], 100),
		ChargingEquipmentRatedInputCurrent:  convert16BitRegister(regs[0x01], 100),
		ChargingEquipmentRatedInputPower:    convert32BitRegister(regs[0x03], regs[0x02], 100),
		ChargingEquipmentRatedOutputVoltage: convert16BitRegister(regs[0x04], 100),
		ChargingEquipmentRatedOutputCurrent: convert16BitRegister(regs[0x05], 100),
		ChargingEquipmentRatedOutputPower:   convert32BitRegister(regs[0x07], regs[0x06], 100),
		ChargingMode:                        regs[0x08],
		RatedOutputCurrentOfLoad:            convert16BitRegister(regs2[0x00], 100),
	}, nil
}

func (dev *Dev) ReadRealTimeData() (RealTimeData, error) {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	err := dev.requestSetup()
	if err != nil {
		return RealTimeData{}, err
	}
	regs, err := dev.mc.ReadRegisters(0x3100, 8, modbus.INPUT_REGISTER)
	if err != nil {
		return RealTimeData{}, err
	}
	regs2, err := dev.mc.ReadRegisters(0x310c, 7, modbus.INPUT_REGISTER)
	if err != nil {
		return RealTimeData{}, err
	}
	regs3, err := dev.mc.ReadRegisters(0x311a, 2, modbus.INPUT_REGISTER)
	if err != nil {
		return RealTimeData{}, err
	}
	regs4, err := dev.mc.ReadRegisters(0x311d, 1, modbus.INPUT_REGISTER)
	if err != nil {
		return RealTimeData{}, err
	}
	return RealTimeData{
		ChargingEquipmentInputVoltage:     convert16BitRegister(regs[0x00], 100),
		ChargingEquipmentInputCurrent:     convert16BitRegister(regs[0x01], 100),
		ChargingEquipmentInputPower:       convert32BitRegister(regs[0x03], regs[0x02], 100),
		ChargingEquipmentOutputVoltage:    convert16BitRegister(regs[0x04], 100),
		ChargingEquipmentOutputCurrent:    convert16BitRegister(regs[0x05], 100),
		ChargingEquipmentOutputPower:      convert32BitRegister(regs[0x07], regs[0x06], 100),
		DischargingEquipmentOutputVoltage: convert16BitRegister(regs2[0x00], 100),
		DischargingEquipmentOutputCurrent: convert16BitRegister(regs2[0x01], 100),
		DischargingEquipmentOutputPower:   convert32BitRegister(regs2[0x03], regs2[0x02], 100),
		BatteryTemperature:                convert16BitRegister(regs2[0x04], 100),
		TemperatureInsideEquipment:        convert16BitRegister(regs2[0x05], 100),
		PowerComponentsTemperature:        convert16BitRegister(regs2[0x06], 100),
		BatterySOC:                        convert16BitRegister(regs3[0x00], 100),
		RemoteBatteryTemperature:          convert16BitRegister(regs3[0x01], 100),
		BatteryRealRatedPower:             convert16BitRegister(regs4[0x00], 100),
	}, nil
}

func (dev *Dev) ReadRealTimeStatus() (RealTimeStatus, error) {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	err := dev.requestSetup()
	if err != nil {
		return RealTimeStatus{}, err
	}
	regs, err := dev.mc.ReadRegisters(0x3200, 2, modbus.INPUT_REGISTER)
	if err != nil {
		return RealTimeStatus{}, err
	}
	return RealTimeStatus{
		BatteryStatus:           BatteryStatus(regs[0x00]),
		ChargingEquipmentStatus: ChargingEquipmentStatus(regs[0x01]),
	}, nil
}

func (dev *Dev) ReadStatistics() (Statistics, error) {
	dev.mutex.Lock()
	defer dev.mutex.Unlock()

	err := dev.requestSetup()
	if err != nil {
		return Statistics{}, err
	}
	regs, err := dev.mc.ReadRegisters(0x3300, 22, modbus.INPUT_REGISTER)
	if err != nil {
		return Statistics{}, err
	}
	regs2, err := dev.mc.ReadRegisters(0x331b, 4, modbus.INPUT_REGISTER)
	if err != nil {
		return Statistics{}, err
	}
	return Statistics{
		MaximumInputVoltageToday:   convert16BitRegister(regs[0x00], 100),
		MinimumInputVoltageToday:   convert16BitRegister(regs[0x01], 100),
		MaximumBatteryVoltageToday: convert16BitRegister(regs[0x02], 100),
		MinimumBatteryVoltageToday: convert16BitRegister(regs[0x03], 100),
		ConsumedEnergyToday:        convert32BitRegister(regs[0x05], regs[0x04], 100),
		ConsumedEnergyThisMonth:    convert32BitRegister(regs[0x07], regs[0x06], 100),
		ConsumedEnergyThisYear:     convert32BitRegister(regs[0x09], regs[0x08], 100),
		TotalConsumedEnergy:        convert32BitRegister(regs[0x0b], regs[0x0a], 100),
		GeneratedEnergyToday:       convert32BitRegister(regs[0x0d], regs[0x0c], 100),
		GeneratedEnergyThisMonth:   convert32BitRegister(regs[0x0f], regs[0x0e], 100),
		GeneratedEnergyThisYear:    convert32BitRegister(regs[0x11], regs[0x10], 100),
		TotalGeneratedEnergy:       convert32BitRegister(regs[0x13], regs[0x12], 100),
		CarbonDioxideReduction:     convert32BitRegister(regs[0x15], regs[0x14], 100),
		BatteryCurrent:             convert32BitRegister(regs2[0x01], regs2[0x00], 100),
		BatteryTemperature:         convert16BitRegister(regs2[0x02], 100),
		AmbientTemperature:         convert16BitRegister(regs2[0x03], 100),
	}, nil
}
