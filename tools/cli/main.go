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
	}

	app = &cli.App{
		Name:  "epsolar",
		Usage: "EPsolar CLI",
		Flags: []cli.Flag{
			serialPortFlag,
			modbusUnitIdFlag,
		},
		Commands: []*cli.Command{
			{
				Name:   "rated-data",
				Usage:  "rated data",
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
