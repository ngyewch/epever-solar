package main

import (
	"sync"

	epsolar "github.com/ngyewch/epever-solar"
	"github.com/urfave/cli/v3"
)

func newDev(cmd *cli.Command) (*epsolar.Dev, error) {
	client, err := newModbusClient(cmd, nil)
	if err != nil {
		return nil, err
	}

	modbusUnitId := cmd.Uint(modbusUnitIdFlag.Name)

	err = client.Open()
	if err != nil {
		return nil, err
	}

	var mutex sync.Mutex

	dev := epsolar.New(client, uint8(modbusUnitId), &mutex)

	return dev, nil
}
