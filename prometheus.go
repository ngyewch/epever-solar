package epsolar

import (
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusCollectorHelper struct {
	// real-time data
	pvArrayInputVoltage *prometheus.Desc
	pvArrayInputCurrent *prometheus.Desc
	pvArrayInputPower   *prometheus.Desc
	loadVoltage         *prometheus.Desc
	loadCurrent         *prometheus.Desc
	loadPower           *prometheus.Desc
	batteryTemperature  *prometheus.Desc
	deviceTemperature   *prometheus.Desc
	batterySOC          *prometheus.Desc
	batteryVoltage      *prometheus.Desc
	batteryCurrent      *prometheus.Desc

	// real-time status
	batteryStatus              *prometheus.Desc
	chargingEquipmentStatus    *prometheus.Desc
	dischargingEquipmentStatus *prometheus.Desc

	// statistics
	maxArrayVoltageToday     *prometheus.Desc
	minArrayVoltageToday     *prometheus.Desc
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
}

func NewPrometheusCollectorHelper(variableLabels []string, constLabels prometheus.Labels) *PrometheusCollectorHelper {
	return &PrometheusCollectorHelper{
		pvArrayInputVoltage: prometheus.NewDesc(
			"epever_solar_pv_array_input_voltage",
			"PV array input voltage (V)",
			variableLabels, constLabels),
		pvArrayInputCurrent: prometheus.NewDesc(
			"epever_solar_pv_array_input_current",
			"PV array input current (A)",
			variableLabels, constLabels),
		pvArrayInputPower: prometheus.NewDesc(
			"epever_solar_pv_array_input_power",
			"PV array input power (W)",
			variableLabels, constLabels),
		loadVoltage: prometheus.NewDesc(
			"epever_solar_load_voltage",
			"Load voltage (V)",
			variableLabels, constLabels),
		loadCurrent: prometheus.NewDesc(
			"epever_solar_load_current",
			"Load current (A)",
			variableLabels, constLabels),
		loadPower: prometheus.NewDesc(
			"epever_solar_load_power",
			"Load power (W)",
			variableLabels, constLabels),
		batteryTemperature: prometheus.NewDesc(
			"epever_solar_battery_temperature",
			"Battery temperature (°C)",
			variableLabels, constLabels),
		deviceTemperature: prometheus.NewDesc(
			"epever_solar_device_temperature",
			"Device temperature (°C)",
			variableLabels, constLabels),
		batterySOC: prometheus.NewDesc(
			"epever_solar_battery_remaining_capacity",
			"Battery remaining capacity (%)",
			variableLabels, constLabels),
		batteryVoltage: prometheus.NewDesc(
			"epever_solar_battery_voltage",
			"Battery voltage (V)",
			variableLabels, constLabels),
		batteryCurrent: prometheus.NewDesc(
			"epever_solar_battery_current",
			"Battery current (A)",
			variableLabels, constLabels),

		batteryStatus: prometheus.NewDesc(
			"epever_solar_battery_status",
			"Battery status",
			variableLabels, constLabels),
		chargingEquipmentStatus: prometheus.NewDesc(
			"epever_solar_charging_equipment_status",
			"Charging equipment status",
			variableLabels, constLabels),
		dischargingEquipmentStatus: prometheus.NewDesc(
			"epever_solar_discharging_equipment_status",
			"Discharging equipment status",
			variableLabels, constLabels),

		maxArrayVoltageToday: prometheus.NewDesc(
			"epever_solar_max_array_voltage_today",
			"Max array voltage today (V)",
			variableLabels, constLabels),
		minArrayVoltageToday: prometheus.NewDesc(
			"epever_solar_min_array_voltage_today",
			"Min array voltage today (V)",
			variableLabels, constLabels),
		maxBatteryVoltageToday: prometheus.NewDesc(
			"epever_solar_max_battery_voltage_today",
			"Max battery voltage today (V)",
			variableLabels, constLabels),
		minBatteryVoltageToday: prometheus.NewDesc(
			"epever_solar_min_battery_voltage_today",
			"Min battery voltage today (V)",
			variableLabels, constLabels),
		consumedEnergyToday: prometheus.NewDesc(
			"epever_solar_consumed_energy_today",
			"Consumed energy today (kWh)",
			variableLabels, constLabels),
		consumedEnergyThisMonth: prometheus.NewDesc(
			"epever_solar_consumed_energy_this_month",
			"Consumed energy this month (kWh)",
			variableLabels, constLabels),
		consumedEnergyThisYear: prometheus.NewDesc(
			"epever_solar_consumed_energy_this_year",
			"Consumed energy this year (kWh)",
			variableLabels, constLabels),
		totalConsumedEnergy: prometheus.NewDesc(
			"epever_solar_total_consumed_energy",
			"Total consumed energy (kWh)",
			variableLabels, constLabels),
		generatedEnergyToday: prometheus.NewDesc(
			"epever_solar_generated_energy_today",
			"Generated energy today (kWh)",
			variableLabels, constLabels),
		generatedEnergyThisMonth: prometheus.NewDesc(
			"epever_solar_generated_energy_this_month",
			"Generated energy this month (kWh)",
			variableLabels, constLabels),
		generatedEnergyThisYear: prometheus.NewDesc(
			"epever_solar_generated_energy_this_year",
			"Generated energy this year (kWh)",
			variableLabels, constLabels),
		totalGeneratedEnergy: prometheus.NewDesc(
			"epever_solar_total_generated_energy",
			"Total generated energy (kWh)",
			variableLabels, constLabels),
	}
}

func (c *PrometheusCollectorHelper) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.pvArrayInputVoltage
	ch <- c.pvArrayInputCurrent
	ch <- c.pvArrayInputPower
	ch <- c.loadVoltage
	ch <- c.loadCurrent
	ch <- c.loadPower
	ch <- c.batteryTemperature
	ch <- c.deviceTemperature
	ch <- c.batterySOC
	ch <- c.batteryVoltage
	ch <- c.batteryCurrent

	ch <- c.batteryStatus
	ch <- c.chargingEquipmentStatus
	ch <- c.dischargingEquipmentStatus

	ch <- c.maxArrayVoltageToday
	ch <- c.minArrayVoltageToday
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
}

func (c *PrometheusCollectorHelper) Collect(dev *Dev, ch chan<- prometheus.Metric, labelValues ...string) {
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
			if realTimeData.PVArrayInputVoltage != nil {
				ch <- prometheus.MustNewConstMetric(c.pvArrayInputVoltage, prometheus.GaugeValue, *realTimeData.PVArrayInputVoltage, labelValues...)
			}
			if realTimeData.PVArrayInputCurrent != nil {
				ch <- prometheus.MustNewConstMetric(c.pvArrayInputCurrent, prometheus.GaugeValue, *realTimeData.PVArrayInputCurrent, labelValues...)
			}
			if realTimeData.PVArrayInputPower != nil {
				ch <- prometheus.MustNewConstMetric(c.pvArrayInputPower, prometheus.GaugeValue, *realTimeData.PVArrayInputPower, labelValues...)
			}
			if realTimeData.LoadVoltage != nil {
				ch <- prometheus.MustNewConstMetric(c.loadVoltage, prometheus.GaugeValue, *realTimeData.LoadVoltage, labelValues...)
			}
			if realTimeData.LoadCurrent != nil {
				ch <- prometheus.MustNewConstMetric(c.loadCurrent, prometheus.GaugeValue, *realTimeData.LoadCurrent, labelValues...)
			}
			if realTimeData.LoadPower != nil {
				ch <- prometheus.MustNewConstMetric(c.loadPower, prometheus.GaugeValue, *realTimeData.LoadPower, labelValues...)
			}
			if realTimeData.BatteryTemperature != nil {
				ch <- prometheus.MustNewConstMetric(c.batteryTemperature, prometheus.GaugeValue, *realTimeData.BatteryTemperature, labelValues...)
			}
			if realTimeData.DeviceTemperature != nil {
				ch <- prometheus.MustNewConstMetric(c.deviceTemperature, prometheus.GaugeValue, *realTimeData.DeviceTemperature, labelValues...)
			}
			if realTimeData.BatterySOC != nil {
				ch <- prometheus.MustNewConstMetric(c.batterySOC, prometheus.GaugeValue, *realTimeData.BatterySOC, labelValues...)
			}
			if realTimeData.BatteryVoltage != nil {
				ch <- prometheus.MustNewConstMetric(c.batteryVoltage, prometheus.GaugeValue, *realTimeData.BatteryVoltage, labelValues...)
			}
			if realTimeData.BatteryCurrent != nil {
				ch <- prometheus.MustNewConstMetric(c.batteryCurrent, prometheus.GaugeValue, *realTimeData.BatteryCurrent, labelValues...)
			}
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
			if realTimeStatus.BatteryStatus != nil {
				ch <- prometheus.MustNewConstMetric(c.batteryStatus, prometheus.GaugeValue, float64((*realTimeStatus.BatteryStatus).Raw), labelValues...)
			}
			if realTimeStatus.ChargingEquipmentStatus != nil {
				ch <- prometheus.MustNewConstMetric(c.chargingEquipmentStatus, prometheus.GaugeValue, float64((realTimeStatus.ChargingEquipmentStatus).Raw), labelValues...)
			}
			if realTimeStatus.DischargingEquipmentStatus != nil {
				ch <- prometheus.MustNewConstMetric(c.dischargingEquipmentStatus, prometheus.GaugeValue, float64((realTimeStatus.DischargingEquipmentStatus).Raw), labelValues...)
			}
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
			if statistics.MaximumArrayVoltageToday != nil {
				ch <- prometheus.MustNewConstMetric(c.maxArrayVoltageToday, prometheus.GaugeValue, *statistics.MaximumArrayVoltageToday, labelValues...)
			}
			if statistics.MinimumArrayVoltageToday != nil {
				ch <- prometheus.MustNewConstMetric(c.minArrayVoltageToday, prometheus.GaugeValue, *statistics.MinimumArrayVoltageToday, labelValues...)
			}
			if statistics.MaximumBatteryVoltageToday != nil {
				ch <- prometheus.MustNewConstMetric(c.maxBatteryVoltageToday, prometheus.GaugeValue, *statistics.MaximumBatteryVoltageToday, labelValues...)
			}
			if statistics.MinimumBatteryVoltageToday != nil {
				ch <- prometheus.MustNewConstMetric(c.minBatteryVoltageToday, prometheus.GaugeValue, *statistics.MinimumBatteryVoltageToday, labelValues...)
			}
			if statistics.ConsumedEnergyToday != nil {
				ch <- prometheus.MustNewConstMetric(c.consumedEnergyToday, prometheus.GaugeValue, *statistics.ConsumedEnergyToday, labelValues...)
			}
			if statistics.ConsumedEnergyThisMonth != nil {
				ch <- prometheus.MustNewConstMetric(c.consumedEnergyThisMonth, prometheus.GaugeValue, *statistics.ConsumedEnergyThisMonth, labelValues...)
			}
			if statistics.ConsumedEnergyThisYear != nil {
				ch <- prometheus.MustNewConstMetric(c.consumedEnergyThisYear, prometheus.GaugeValue, *statistics.ConsumedEnergyThisYear, labelValues...)
			}
			if statistics.TotalConsumedEnergy != nil {
				ch <- prometheus.MustNewConstMetric(c.totalConsumedEnergy, prometheus.GaugeValue, *statistics.TotalConsumedEnergy, labelValues...)
			}
			if statistics.GeneratedEnergyToday != nil {
				ch <- prometheus.MustNewConstMetric(c.generatedEnergyToday, prometheus.GaugeValue, *statistics.GeneratedEnergyToday, labelValues...)
			}
			if statistics.GeneratedEnergyThisMonth != nil {
				ch <- prometheus.MustNewConstMetric(c.generatedEnergyThisMonth, prometheus.GaugeValue, *statistics.GeneratedEnergyThisMonth, labelValues...)
			}
			if statistics.GeneratedEnergyThisYear != nil {
				ch <- prometheus.MustNewConstMetric(c.generatedEnergyThisYear, prometheus.GaugeValue, *statistics.GeneratedEnergyThisYear, labelValues...)
			}
			if statistics.TotalGeneratedEnergy != nil {
				ch <- prometheus.MustNewConstMetric(c.totalGeneratedEnergy, prometheus.GaugeValue, *statistics.TotalGeneratedEnergy, labelValues...)
			}
		}()
	}
}
