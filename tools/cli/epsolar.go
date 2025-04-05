package main

import (
	"encoding/json"
	"fmt"
	"github.com/ngyewch/epsolar"
	"github.com/simonvetter/modbus"
	"github.com/urfave/cli/v2"
	"os"
	"sync"
	"time"
)

func newEpsolar(cCtx *cli.Context) (*epsolar.Dev, error) {
	serialPort := serialPortFlag.Get(cCtx)
	modbusUnitId := modbusUnitIdFlag.Get(cCtx)

	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      "rtu://" + serialPort,
		Speed:    115200,
		DataBits: 8,
		Parity:   modbus.PARITY_NONE,
		StopBits: 1,
		Timeout:  1 * time.Second,
	})
	if err != nil {
		return nil, err
	}

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

	err = printJSON(ratedData)
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

	err = printJSON(realTimeData)
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

	err = printJSON(realTimeStatus)
	if err != nil {
		return err
	}
	fmt.Println()

	fmt.Println("# Battery status details")
	err = printJSON(realTimeStatus.BatteryStatus.Details())
	if err != nil {
		return err
	}
	fmt.Println()

	fmt.Println("# Charging equipment status details")
	err = printJSON(realTimeStatus.ChargingEquipmentStatus.Details())
	if err != nil {
		return err
	}
	fmt.Println()

	fmt.Println("# Discharging equipment status details")
	err = printJSON(realTimeStatus.DischargingEquipmentStatus.Details())
	if err != nil {
		return err
	}
	fmt.Println()

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

	err = printJSON(statistics)
	if err != nil {
		return err
	}

	return nil
}

func printJSON(v any) error {
	jsonEncoder := json.NewEncoder(os.Stdout)
	jsonEncoder.SetIndent("", "  ")
	jsonEncoder.SetEscapeHTML(false)
	return jsonEncoder.Encode(v)
}
