package main

import (
	"github.com/ngyewch/epever-solar"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	"github.com/urfave/cli/v2"
	"os"
)

func doEpsolarRatedData(cCtx *cli.Context) error {
	dev, err := newDev(cCtx)
	if err != nil {
		return err
	}

	ratedData, err := dev.ReadRatedData()
	if err != nil {
		return err
	}

	err = dump(ratedData)
	if err != nil {
		return err
	}

	return nil
}

func doEpsolarParameters(cCtx *cli.Context) error {
	dev, err := newDev(cCtx)
	if err != nil {
		return err
	}

	parameters, err := dev.ReadParameters()
	if err != nil {
		return err
	}

	err = dump(parameters)
	if err != nil {
		return err
	}

	return nil
}

func doEpsolarRealTimeData(cCtx *cli.Context) error {
	dev, err := newDev(cCtx)
	if err != nil {
		return err
	}

	realTimeData, err := dev.ReadRealTimeData()
	if err != nil {
		return err
	}

	err = dump(realTimeData)
	if err != nil {
		return err
	}

	return nil
}

func doEpsolarRealTimeStatus(cCtx *cli.Context) error {
	dev, err := newDev(cCtx)
	if err != nil {
		return err
	}

	realTimeStatus, err := dev.ReadRealTimeStatus()
	if err != nil {
		return err
	}

	err = dump(realTimeStatus)
	if err != nil {
		return err
	}

	return nil
}

func doEpsolarStatistics(cCtx *cli.Context) error {
	dev, err := newDev(cCtx)
	if err != nil {
		return err
	}

	statistics, err := dev.ReadStatistics()
	if err != nil {
		return err
	}

	err = dump(statistics)
	if err != nil {
		return err
	}

	return nil
}

func doEpsolarPrometheus(cCtx *cli.Context) error {
	dev, err := newDev(cCtx)
	if err != nil {
		return err
	}

	reg := prometheus.NewRegistry()
	collectorHelper := epsolar.NewPrometheusCollectorHelper(nil, nil)
	c := collector{
		dev:    dev,
		helper: collectorHelper,
	}
	err = reg.Register(&c)
	if err != nil {
		return err
	}
	gatherer := prometheus.Gatherers{
		reg,
	}
	metricFamilies, err := gatherer.Gather()
	if err != nil {
		return err
	}

	fmt := expfmt.NewFormat(expfmt.TypeTextPlain)
	encoder := expfmt.NewEncoder(os.Stdout, fmt)
	for _, mf := range metricFamilies {
		err = encoder.Encode(mf)
		if err != nil {
			return err
		}
	}

	return nil
}

type collector struct {
	dev    *epsolar.Dev
	helper *epsolar.PrometheusCollectorHelper
}

func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	c.helper.Describe(ch)
}

func (c *collector) Collect(ch chan<- prometheus.Metric) {
	c.helper.Collect(c.dev, ch)
}
