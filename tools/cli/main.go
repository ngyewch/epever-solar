package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"runtime/debug"
)

var (
	serialPortFlag = &cli.StringFlag{
		Name:     "serial-port",
		Usage:    "serial port",
		Required: true,
		EnvVars:  []string{"SERIAL_PORT"},
		Category: "Serial",
	}
	baudRateFlag = &cli.UintFlag{
		Name:     "baud-rate",
		Usage:    "baud rate",
		Value:    115200,
		EnvVars:  []string{"BAUD_RATE"},
		Category: "Serial",
	}
	dataBitsFlag = &cli.UintFlag{
		Name:     "data-bits",
		Usage:    "data bits",
		Value:    8,
		EnvVars:  []string{"DATA_BITS"},
		Category: "Serial",
	}
	parityFlag = &cli.StringFlag{
		Name:    "parity",
		Usage:   "parity",
		Value:   "N",
		EnvVars: []string{"PARITY"},
		Action: func(cCtx *cli.Context, s string) error {
			_, err := parseParity(s)
			return err
		},
		Category: "Serial",
	}
	stopBitsFlag = &cli.UintFlag{
		Name:     "stop-bits",
		Usage:    "stop bits",
		Value:    1,
		EnvVars:  []string{"STOP_BITS"},
		Category: "Serial",
	}
	modbusUnitIdFlag = &cli.UintFlag{
		Name:    "modbus-unit-id",
		Usage:   "ModBus unit ID",
		Value:   1,
		EnvVars: []string{"MODBUS_UNIT_ID"},
		Action: func(cCtx *cli.Context, v uint) error {
			if (v < 1) || (v > 247) {
				return fmt.Errorf("invalid modbus-unit-id: %d", v)
			}
			return nil
		},
		Category: "Modbus",
	}

	app = &cli.App{
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
				Subcommands: []*cli.Command{
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
	buildInfo, _ := debug.ReadBuildInfo()
	if buildInfo != nil {
		app.Version = buildInfo.Main.Version
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
