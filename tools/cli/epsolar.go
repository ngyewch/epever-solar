package main

import (
	"context"
	"fmt"
	"github.com/ngyewch/epever-solar"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	"github.com/urfave/cli/v3"
	"os"
	"time"
)

func doEpsolarRatedData(ctx context.Context, cmd *cli.Command) error {
	dev, err := newDev(cmd)
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

func doEpsolarParameters(ctx context.Context, cmd *cli.Command) error {
	dev, err := newDev(cmd)
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

func doEpsolarRealTimeData(ctx context.Context, cmd *cli.Command) error {
	dev, err := newDev(cmd)
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

func doEpsolarRealTimeStatus(ctx context.Context, cmd *cli.Command) error {
	dev, err := newDev(cmd)
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

func doEpsolarStatistics(ctx context.Context, cmd *cli.Command) error {
	dev, err := newDev(cmd)
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

func doEpsolarPrometheus(ctx context.Context, cmd *cli.Command) error {
	dev, err := newDev(cmd)
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

func doEpsolarRTCGet(ctx context.Context, cmd *cli.Command) error {
	dev, err := newDev(cmd)
	if err != nil {
		return err
	}

	rtc, err := dev.ReadRealTimeClock()
	if err != nil {
		return err
	}

	err = dump(rtc)
	if err != nil {
		return err
	}

	return nil
}

func doEpsolarRTCSet(ctx context.Context, cmd *cli.Command) error {
	t, err := func() (time.Time, error) {
		switch cmd.NArg() {
		case 0:
			return time.Now(), nil
		case 1:
			return time.Parse(time.DateTime, cmd.Args().Get(0))
		default:
			return time.Time{}, fmt.Errorf("too many arguments")
		}
	}()
	if err != nil {
		return err
	}

	rtcData := epsolar.RTCData{
		Year:   uint8(t.Year() % 100),
		Month:  uint8(t.Month()),
		Day:    uint8(t.Day()),
		Hour:   uint8(t.Hour()),
		Minute: uint8(t.Minute()),
		Second: uint8(t.Second()),
	}

	dev, err := newDev(cmd)
	if err != nil {
		return err
	}

	err = dev.SetRealTimeClock(rtcData)
	if err != nil {
		return err
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
