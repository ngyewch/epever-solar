package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/urfave/cli/v3"
)

var (
	version string

	serialPortFlag = &cli.StringFlag{
		Name:     "serial-port",
		Usage:    "serial port",
		Required: true,
		Sources:  cli.EnvVars("SERIAL_PORT"),
		Category: "Serial",
	}
	baudRateFlag = &cli.UintFlag{
		Name:     "baud-rate",
		Usage:    "baud rate",
		Value:    115200,
		Sources:  cli.EnvVars("BAUD_RATE"),
		Category: "Serial",
	}
	dataBitsFlag = &cli.UintFlag{
		Name:     "data-bits",
		Usage:    "data bits",
		Value:    8,
		Sources:  cli.EnvVars("DATA_BITS"),
		Category: "Serial",
	}
	parityFlag = &cli.StringFlag{
		Name:    "parity",
		Usage:   "parity",
		Value:   "N",
		Sources: cli.EnvVars("PARITY"),
		Action: func(ctx context.Context, cmd *cli.Command, s string) error {
			_, err := parseParity(s)
			return err
		},
		Category: "Serial",
	}
	stopBitsFlag = &cli.UintFlag{
		Name:     "stop-bits",
		Usage:    "stop bits",
		Value:    1,
		Sources:  cli.EnvVars("STOP_BITS"),
		Category: "Serial",
	}
	modbusUnitIdFlag = &cli.UintFlag{
		Name:    "modbus-unit-id",
		Usage:   "ModBus unit ID",
		Value:   1,
		Sources: cli.EnvVars("MODBUS_UNIT_ID"),
		Action: func(ctx context.Context, cmd *cli.Command, v uint) error {
			if (v < 1) || (v > 247) {
				return fmt.Errorf("invalid modbus-unit-id: %d", v)
			}
			return nil
		},
		Category: "Modbus",
	}

	app = &cli.Command{
		Name:  "epsolar",
		Usage: "EPsolar CLI",
		Flags: []cli.Flag{
			serialPortFlag,
			baudRateFlag,
			dataBitsFlag,
			parityFlag,
			stopBitsFlag,
			modbusUnitIdFlag,
		},
		Commands: []*cli.Command{
			{
				Name:   "rated-data",
				Usage:  "rated-data",
				Action: doEpsolarRatedData,
			},
			{
				Name:   "real-time-data",
				Usage:  "real-time data",
				Action: doEpsolarRealTimeData,
			},
			{
				Name:   "real-time-status",
				Usage:  "real-time status",
				Action: doEpsolarRealTimeStatus,
			},
			{
				Name:   "statistics",
				Usage:  "statistics",
				Action: doEpsolarStatistics,
			},
			{
				Name:   "parameters",
				Usage:  "parameters",
				Action: doEpsolarParameters,
			},
			{
				Name:  "rtc",
				Usage: "rtc",
				Commands: []*cli.Command{
					{
						Name:   "get",
						Usage:  "get",
						Action: doEpsolarRTCGet,
					},
					{
						Name:      "set",
						Usage:     "set",
						ArgsUsage: "[(date time)]",
						Action:    doEpsolarRTCSet,
					},
				},
			},
			{
				Name:   "prometheus",
				Usage:  "prometheus",
				Action: doEpsolarPrometheus,
			},
		},
	}
)

func main() {
	if version == "" {
		buildInfo, _ := debug.ReadBuildInfo()
		if buildInfo != nil {
			version = buildInfo.Main.Version
		}
	}
	app.Version = version

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
