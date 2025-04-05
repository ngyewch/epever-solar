package epsolar

import (
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
)

type PrometheusCollectorHelper struct {
	// rated data
	ratedChargingCurrent    *prometheus.Desc
	ratedLoadCurrent        *prometheus.Desc
	batteryRealRatedVoltage *prometheus.Desc

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
}

func NewPrometheusCollectorHelper(variableLabels []string, constLabels prometheus.Labels) *PrometheusCollectorHelper {
	return &PrometheusCollectorHelper{
		ratedChargingCurrent: prometheus.NewDesc(
			"epever_solar_rated_charging_current",
			"Rated chargine current (A)",
			variableLabels, constLabels),
		ratedLoadCurrent: prometheus.NewDesc(
			"epever_solar_rated_load_current",
			"Rated load current (A)",
			variableLabels, constLabels),
		batteryRealRatedVoltage: prometheus.NewDesc(
			"epever_solar_battery_real_rated_voltage",
			"Battery real rated voltage (V)",
			variableLabels, constLabels),

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

		maxInputVoltageToday: prometheus.NewDesc(
			"epever_solar_max_input_voltage_today",
			"Max input voltage today (V)",
			variableLabels, constLabels),
		minInputVoltageToday: prometheus.NewDesc(
			"epever_solar_min_input_voltage_today",
			"Min input voltage today (V)",
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
	ch <- c.ratedChargingCurrent
	ch <- c.ratedLoadCurrent
	ch <- c.batteryRealRatedVoltage

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
			ch <- prometheus.MustNewConstMetric(c.ratedChargingCurrent, prometheus.GaugeValue, ratedData.RatedChargingCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.ratedLoadCurrent, prometheus.GaugeValue, ratedData.RatedLoadCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.batteryRealRatedVoltage, prometheus.GaugeValue, ratedData.BatteryRealRatedVoltage, labelValues...)
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
			ch <- prometheus.MustNewConstMetric(c.pvArrayInputVoltage, prometheus.GaugeValue, realTimeData.PVArrayInputVoltage, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.pvArrayInputCurrent, prometheus.GaugeValue, realTimeData.PVArrayInputCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.pvArrayInputPower, prometheus.GaugeValue, realTimeData.PVArrayInputCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.loadVoltage, prometheus.GaugeValue, realTimeData.LoadVoltage, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.loadCurrent, prometheus.GaugeValue, realTimeData.LoadCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.loadPower, prometheus.GaugeValue, realTimeData.LoadCurrent, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.batteryTemperature, prometheus.GaugeValue, realTimeData.BatteryTemperature, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.deviceTemperature, prometheus.GaugeValue, realTimeData.DeviceTemperature, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.batterySOC, prometheus.GaugeValue, realTimeData.BatterySOC, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.batteryVoltage, prometheus.GaugeValue, realTimeData.BatteryVoltage, labelValues...)
			ch <- prometheus.MustNewConstMetric(c.batteryCurrent, prometheus.GaugeValue, realTimeData.BatteryCurrent, labelValues...)
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
		}()
	}
}
