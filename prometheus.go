package epsolar

import (
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
)

type PrometheusCollectorHelper struct {
	// rated data
	chargingEquipmentRatedInputVoltage  *prometheus.Desc
	chargingEquipmentRatedInputCurrent  *prometheus.Desc
	chargingEquipmentRatedInputPower    *prometheus.Desc
	chargingEquipmentRatedOutputVoltage *prometheus.Desc
	chargingEquipmentRatedOutputCurrent *prometheus.Desc
	chargingEquipmentRatedOutputPower   *prometheus.Desc
	chargingMode                        *prometheus.Desc
	ratedOutputCurrentOfLoad            *prometheus.Desc

	// real-time data
	chargingEquipmentInputVoltage     *prometheus.Desc
	chargingEquipmentInputCurrent     *prometheus.Desc
	chargingEquipmentInputPower       *prometheus.Desc
	chargingEquipmentOutputVoltage    *prometheus.Desc
	chargingEquipmentOutputCurrent    *prometheus.Desc
	chargingEquipmentOutputPower      *prometheus.Desc
	dischargingEquipmentOutputVoltage *prometheus.Desc
	dischargingEquipmentOutputCurrent *prometheus.Desc
	dischargingEquipmentOutputPower   *prometheus.Desc
	batteryTemperature                *prometheus.Desc
	temperatureInsideEquipment        *prometheus.Desc
	powerComponentsTemperature        *prometheus.Desc
	batterySOC                        *prometheus.Desc
	currentSystemRatedVoltage         *prometheus.Desc

	// real-time status
	batteryStatus           *prometheus.Desc
	chargingEquipmentStatus *prometheus.Desc

	// statistics
	maxInputVoltageToday     *prometheus.Desc
	minInputVoltageToday     *prometheus.Desc
	maxBatteryVoltageToday   *prometheus.Desc
	minBatteryVoltageToday   *prometheus.Desc
	consumedEnergyToday      *prometheus.Desc
	consumedEnergyThisMonth  *prometheus.Desc
	consumedEnergyThisYear   *prometheus.Desc
	totalConsumedEnergy      *prometheus.Desc
	generatedEnergyToday     *prometheus.Desc
	generatedEnergyThisMonth *prometheus.Desc
	generatedEnergyThisYear  *prometheus.Desc
	totalGeneratedEnergy     *prometheus.Desc
	carbonDioxideReduction   *prometheus.Desc
}

func NewPrometheusCollectorHelper(variableLabels []string, constLabels prometheus.Labels) *PrometheusCollectorHelper {
	return &PrometheusCollectorHelper{
		chargingEquipmentRatedInputVoltage: prometheus.NewDesc(
			"epsolar_charging_equipment_rated_input_voltage",
			"Charging equipment rated input voltage (V)",
			variableLabels, constLabels),
		chargingEquipmentRatedInputCurrent: prometheus.NewDesc(
			"epsolar_charging_equipment_rated_input_current",
			"Charging equipment rated input current (A)",
			variableLabels, constLabels),
		chargingEquipmentRatedInputPower: prometheus.NewDesc(
			"epsolar_charging_equipment_rated_input_power",
			"Charging equipment rated input power (W)",
			variableLabels, constLabels),
		chargingEquipmentRatedOutputVoltage: prometheus.NewDesc(
			"epsolar_charging_equipment_rated_output_voltage",
			"Charging equipment rated output voltage (V)",
			variableLabels, constLabels),
		chargingEquipmentRatedOutputCurrent: prometheus.NewDesc(
			"epsolar_charging_equipment_rated_output_current",
			"Charging equipment rated output current (A)",
			variableLabels, constLabels),
		chargingEquipmentRatedOutputPower: prometheus.NewDesc(
			"epsolar_charging_equipment_rated_output_power",
			"Charging equipment rated output power (W)",
			variableLabels, constLabels),
		chargingMode: prometheus.NewDesc(
			"epsolar_charging_mode",
			"Charging mode",
			variableLabels, constLabels),
		ratedOutputCurrentOfLoad: prometheus.NewDesc(
			"epsolar_rated_output_current_of_load",
			"Rated output current of load (A)",
			variableLabels, constLabels),

		chargingEquipmentInputVoltage: prometheus.NewDesc(
			"epsolar_charging_equipment_input_voltage",
			"Charging equipment input voltage (V)",
			variableLabels, constLabels),
		chargingEquipmentInputCurrent: prometheus.NewDesc(
			"epsolar_charging_equipment_input_current",
			"Charging equipment input current (A)",
			variableLabels, constLabels),
		chargingEquipmentInputPower: prometheus.NewDesc(
			"epsolar_charging_equipment_input_power",
			"Charging equipment input power (W)",
			variableLabels, constLabels),
		chargingEquipmentOutputVoltage: prometheus.NewDesc(
			"epsolar_charging_equipment_output_voltage",
			"Charging equipment output voltage (V)",
			variableLabels, constLabels),
		chargingEquipmentOutputCurrent: prometheus.NewDesc(
			"epsolar_charging_equipment_output_current",
			"Charging equipment output current (A)",
			variableLabels, constLabels),
		chargingEquipmentOutputPower: prometheus.NewDesc(
			"epsolar_charging_equipment_output_power",
			"Charging equipment output power (W)",
			variableLabels, constLabels),
		dischargingEquipmentOutputVoltage: prometheus.NewDesc(
			"epsolar_discharging_equipment_output_voltage",
			"Discharging equipment output voltage (V)",
			variableLabels, constLabels),
		dischargingEquipmentOutputCurrent: prometheus.NewDesc(
			"epsolar_discharging_equipment_output_current",
			"Discharging equipment output current (A)",
			variableLabels, constLabels),
		dischargingEquipmentOutputPower: prometheus.NewDesc(
			"epsolar_discharging_equipment_output_power",
			"Discharging equipment output power (W)",
			variableLabels, constLabels),
		batteryTemperature: prometheus.NewDesc(
			"epsolar_battery_temperature",
			"Battery temperature (°C)",
			variableLabels, constLabels),
		temperatureInsideEquipment: prometheus.NewDesc(
			"epsolar_temperature_inside_equipment",
			"Battery temperature (°C)",
			variableLabels, constLabels),
		powerComponentsTemperature: prometheus.NewDesc(
			"epsolar_power_components_temperature",
			"Power components temperature (°C)",
			variableLabels, constLabels),
		batterySOC: prometheus.NewDesc(
			"epsolar_battery_remaining_capacity",
			"Battery remaining capacity (%)",
			variableLabels, constLabels),
		currentSystemRatedVoltage: prometheus.NewDesc(
			"epsolar_current_system_rated_voltage",
			"Current system rated voltage (V)",
			variableLabels, constLabels),

		batteryStatus: prometheus.NewDesc(
			"epsolar_battery_status",
			"Battery status",
			variableLabels, constLabels),
		chargingEquipmentStatus: prometheus.NewDesc(
			"epsolar_charging_equipment_status",
			"Charging equipment status",
			variableLabels, constLabels),

		maxInputVoltageToday: prometheus.NewDesc(
			"epsolar_max_input_voltage_today",
			"Max input voltage today (V)",
			variableLabels, constLabels),
		minInputVoltageToday: prometheus.NewDesc(
			"epsolar_min_input_voltage_today",
			"Min input voltage today (V)",
			variableLabels, constLabels),
		maxBatteryVoltageToday: prometheus.NewDesc(
			"epsolar_max_battery_voltage_today",
			"Max battery voltage today (V)",
			variableLabels, constLabels),
		minBatteryVoltageToday: prometheus.NewDesc(
			"epsolar_min_battery_voltage_today",
			"Min battery voltage today (V)",
			variableLabels, constLabels),
		consumedEnergyToday: prometheus.NewDesc(
			"epsolar_consumed_energy_today",
			"Consumed energy today (kWh)",
			variableLabels, constLabels),
		consumedEnergyThisMonth: prometheus.NewDesc(
			"epsolar_consumed_energy_this_month",
			"Consumed energy this month (kWh)",
			variableLabels, constLabels),
		consumedEnergyThisYear: prometheus.NewDesc(
			"epsolar_consumed_energy_this_year",
			"Consumed energy this year (kWh)",
			variableLabels, constLabels),
		totalConsumedEnergy: prometheus.NewDesc(
			"epsolar_total_consumed_energy",
			"Total consumed energy (kWh)",
			variableLabels, constLabels),
		generatedEnergyToday: prometheus.NewDesc(
			"epsolar_generated_energy_today",
			"Generated energy today (kWh)",
			variableLabels, constLabels),
		generatedEnergyThisMonth: prometheus.NewDesc(
			"epsolar_generated_energy_this_month",
			"Generated energy this month (kWh)",
			variableLabels, constLabels),
		generatedEnergyThisYear: prometheus.NewDesc(
			"epsolar_generated_energy_this_year",
			"Generated energy this year (kWh)",
			variableLabels, constLabels),
		totalGeneratedEnergy: prometheus.NewDesc(
			"epsolar_total_generated_energy",
			"Total generated energy (kWh)",
			variableLabels, constLabels),
		carbonDioxideReduction: prometheus.NewDesc(
			"epsolar_carbon_dioxide_reduction",
			"Carbon dioxide reduction (ton)",
			variableLabels, constLabels),
	}
}

func (c *PrometheusCollectorHelper) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.chargingEquipmentRatedInputVoltage
	ch <- c.chargingEquipmentRatedInputCurrent
	ch <- c.chargingEquipmentRatedInputPower
	ch <- c.chargingEquipmentOutputVoltage
	ch <- c.chargingEquipmentOutputCurrent
	ch <- c.chargingEquipmentOutputPower
	ch <- c.chargingMode
	ch <- c.ratedOutputCurrentOfLoad

	ch <- c.chargingEquipmentInputVoltage
	ch <- c.chargingEquipmentInputCurrent
	ch <- c.chargingEquipmentInputPower
	ch <- c.chargingEquipmentOutputVoltage
	ch <- c.chargingEquipmentOutputCurrent
	ch <- c.chargingEquipmentOutputPower
	ch <- c.dischargingEquipmentOutputVoltage
	ch <- c.dischargingEquipmentOutputCurrent
	ch <- c.dischargingEquipmentOutputPower
	ch <- c.batteryTemperature
	ch <- c.temperatureInsideEquipment
	ch <- c.powerComponentsTemperature
	ch <- c.batterySOC
	ch <- c.currentSystemRatedVoltage

	ch <- c.batteryStatus
	ch <- c.chargingEquipmentStatus

	ch <- c.maxInputVoltageToday
	ch <- c.minInputVoltageToday
	ch <- c.maxBatteryVoltageToday
	ch <- c.minBatteryVoltageToday
	ch <- c.consumedEnergyToday
	ch <- c.consumedEnergyThisMonth
	ch <- c.consumedEnergyThisYear
	ch <- c.totalConsumedEnergy
	ch <- c.generatedEnergyToday
	ch <- c.generatedEnergyThisMonth
	ch <- c.generatedEnergyThisYear
	ch <- c.totalGeneratedEnergy
	ch <- c.carbonDioxideReduction
}

func (c *PrometheusCollectorHelper) Collect(dev *Dev, ch chan<- prometheus.Metric, labelValues ...string) {
	ratedData, err := dev.ReadRatedData()
	if err != nil {
		slog.Warn("failed to read rated data",
			slog.Any("error", err),
		)
	} else {
		func() {
			defer func() {
				if err := recover(); err != nil {
					slog.Error("failed to create metric",
						slog.Any("error", err),
					)
				}
			}()
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentRatedInputVoltage, prometheus.GaugeValue, ratedData.ChargingEquipmentRatedInputVoltage, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentRatedInputCurrent, prometheus.GaugeValue, ratedData.ChargingEquipmentRatedInputCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentRatedInputPower, prometheus.GaugeValue, ratedData.ChargingEquipmentRatedInputPower, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentRatedOutputVoltage, prometheus.GaugeValue, ratedData.ChargingEquipmentRatedOutputVoltage, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentRatedOutputCurrent, prometheus.GaugeValue, ratedData.ChargingEquipmentRatedOutputCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentRatedOutputPower, prometheus.GaugeValue, ratedData.ChargingEquipmentRatedOutputPower, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.chargingMode, prometheus.GaugeValue, float64(ratedData.ChargingMode), labelValues...)
			ch <- prometheus.MustNewConstMetric(c.ratedOutputCurrentOfLoad, prometheus.GaugeValue, ratedData.RatedOutputCurrentOfLoad, labelValues...)
		}()
	}

	realTimeData, err := dev.ReadRealTimeData()
	if err != nil {
		slog.Warn("failed to read real-time data",
			slog.Any("error", err),
		)
	} else {
		func() {
			defer func() {
				if err := recover(); err != nil {
					slog.Error("failed to create metric",
						slog.Any("error", err),
					)
				}
			}()
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentInputVoltage, prometheus.GaugeValue, realTimeData.ChargingEquipmentInputVoltage, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentInputCurrent, prometheus.GaugeValue, realTimeData.ChargingEquipmentInputCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentInputPower, prometheus.GaugeValue, realTimeData.ChargingEquipmentInputCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentOutputVoltage, prometheus.GaugeValue, realTimeData.ChargingEquipmentOutputVoltage, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentOutputCurrent, prometheus.GaugeValue, realTimeData.ChargingEquipmentOutputCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentOutputPower, prometheus.GaugeValue, realTimeData.ChargingEquipmentOutputCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.dischargingEquipmentOutputVoltage, prometheus.GaugeValue, realTimeData.DischargingEquipmentOutputVoltage, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.dischargingEquipmentOutputCurrent, prometheus.GaugeValue, realTimeData.DischargingEquipmentOutputCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.dischargingEquipmentOutputPower, prometheus.GaugeValue, realTimeData.DischargingEquipmentOutputCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.batteryTemperature, prometheus.GaugeValue, realTimeData.BatteryTemperature, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.temperatureInsideEquipment, prometheus.GaugeValue, realTimeData.TemperatureInsideEquipment, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.powerComponentsTemperature, prometheus.GaugeValue, realTimeData.PowerComponentsTemperature, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.batterySOC, prometheus.GaugeValue, realTimeData.BatterySOC, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.currentSystemRatedVoltage, prometheus.GaugeValue, realTimeData.BatteryRealRatedPower, labelValues...)
		}()
	}

	realTimeStatus, err := dev.ReadRealTimeStatus()
	if err != nil {
		slog.Warn("failed to read real-time status",
			slog.Any("error", err),
		)
	} else {
		func() {
			defer func() {
				if err := recover(); err != nil {
					slog.Error("failed to create metric",
						slog.Any("error", err),
					)
				}
			}()
			ch <- prometheus.MustNewConstMetric(c.batteryStatus, prometheus.GaugeValue, float64(realTimeStatus.BatteryStatus), labelValues...)
			ch <- prometheus.MustNewConstMetric(c.chargingEquipmentStatus, prometheus.GaugeValue, float64(realTimeStatus.ChargingEquipmentStatus), labelValues...)
		}()
	}

	statistics, err := dev.ReadStatistics()
	if err != nil {
		slog.Warn("failed to read statistics",
			slog.Any("error", err),
		)
	} else {
		func() {
			defer func() {
				if err := recover(); err != nil {
					slog.Error("failed to create metric",
						slog.Any("error", err),
					)
				}
			}()
			ch <- prometheus.MustNewConstMetric(c.maxInputVoltageToday, prometheus.GaugeValue, statistics.MaximumInputVoltageToday, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.minInputVoltageToday, prometheus.GaugeValue, statistics.MinimumInputVoltageToday, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.maxBatteryVoltageToday, prometheus.GaugeValue, statistics.MaximumBatteryVoltageToday, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.minBatteryVoltageToday, prometheus.GaugeValue, statistics.MinimumBatteryVoltageToday, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.consumedEnergyToday, prometheus.GaugeValue, statistics.ConsumedEnergyToday, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.consumedEnergyThisMonth, prometheus.GaugeValue, statistics.ConsumedEnergyThisMonth, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.consumedEnergyThisYear, prometheus.GaugeValue, statistics.ConsumedEnergyThisYear, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.totalConsumedEnergy, prometheus.GaugeValue, statistics.TotalConsumedEnergy, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.generatedEnergyToday, prometheus.GaugeValue, statistics.GeneratedEnergyToday, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.generatedEnergyThisMonth, prometheus.GaugeValue, statistics.GeneratedEnergyThisMonth, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.generatedEnergyThisYear, prometheus.GaugeValue, statistics.GeneratedEnergyThisYear, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.totalGeneratedEnergy, prometheus.GaugeValue, statistics.TotalGeneratedEnergy, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.carbonDioxideReduction, prometheus.GaugeValue, statistics.CarbonDioxideReduction, labelValues...)
		}()
	}
}
