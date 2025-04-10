package main

import (
	epsolar "github.com/ngyewch/epever-solar"
	"github.com/urfave/cli/v2"
	"sync"
)

func newDev(cCtx *cli.Context) (*epsolar.Dev, error) {
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

	dev := epsolar.New(client, uint8(modbusUnitId), &mutex)

	return dev, nil
}
