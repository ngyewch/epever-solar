package main

import (
	"github.com/ngyewch/epever-solar"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	"github.com/urfave/cli/v2"
	"os"
	"sync"
)

func newEpsolar(cCtx *cli.Context) (*epsolar.Dev, error) {
	client, err := newModbusClient(cCtx, nil)
	if err != nil {
		return nil, err
	}

	modbusUnitId := modbusUnitIdFlag.Get(cCtx)

	err = client.Open()
	if err != nil {
		return nil, err
	}

	var mutex sync.Mutex

	e := epsolar.New(client, uint8(modbusUnitId), &mutex)

	return e, nil
}

func doEpsolarRatedData(cCtx *cli.Context) error {
	e, err := newEpsolar(cCtx)
	if err != nil {
		return err
	}

	ratedData, err := e.ReadRatedData()
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
	e, err := newEpsolar(cCtx)
	if err != nil {
		return err
	}

	parameters, err := e.ReadParameters()
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
	e, err := newEpsolar(cCtx)
	if err != nil {
		return err
	}

	realTimeData, err := e.ReadRealTimeData()
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
	e, err := newEpsolar(cCtx)
	if err != nil {
		return err
	}

	realTimeStatus, err := e.ReadRealTimeStatus()
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
	e, err := newEpsolar(cCtx)
	if err != nil {
		return err
	}

	statistics, err := e.ReadStatistics()
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
	e, err := newEpsolar(cCtx)
	if err != nil {
		return err
	}

	reg := prometheus.NewRegistry()
	collectorHelper := epsolar.NewPrometheusCollectorHelper(nil, nil)
	c := collector{
		dev:    e,
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
