package accmeter

import (
	"fmt"
	"sync"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

type Device struct {
	Path string `json:"path"`
	bus  i2c.BusCloser
	dev  *i2c.Dev
	mux  sync.Mutex
}

func (d *Device) Close() error {
	return d.bus.Close()
}

func (d *Device) String() string {
	return d.Path
}

func (d *Device) Init(opts ...OptFunc) error {
	if _, err := host.Init(); err != nil {
		return err
	}
	if bus, err := i2creg.Open(d.Path); err != nil {
		return err
	} else {
		d.bus = bus
		dev := &i2c.Dev{Addr: 0x1E, Bus: bus}
		d.dev = dev
	}

	fmt.Println("init device")
	for _, reg := range []*Register{&CNTL1, &CNTL2, &CNTL3, &INC1, &INC2, &INC3, &INC4, &INT_REL, &ODCNTL, &WUFC, &ATH} {
		if data, err := reg.Read(d); err != nil {
			reg.newValue = data
			return err
		}
	}
	opt1 := SetStandByMode()
	opt1()
	if len(opts) == 0 {
		return nil
	}
	for _, opt := range opts {
		opt()
	}
	for _, reg := range []*Register{&CNTL1, &CNTL2, &CNTL3, &INC1, &INC2, &INC3, &INC4, &INT_REL, &ODCNTL, &WUFC, &ATH} {
		if err := reg.updatereg(d); err != nil {
			return err
		}
	}

	opt2 := SetOperatingMode()
	opt2()
	for _, opt := range opts {
		opt()
	}
	for _, reg := range []*Register{&CNTL1} {
		if err := reg.updatereg(d); err != nil {
			return err
		}
	}

	return nil
}

func (d *Device) NewConf(opts ...OptFunc) error {
	if d.dev == nil {
		return d.Init(opts...)
	}

	if len(opts) == 0 {
		return nil
	}

	opt1 := SetStandByMode()
	opt1()
	for _, opt := range opts {
		opt()
	}
	for _, reg := range []*Register{&CNTL1, &CNTL2, &CNTL3, &INC1, &INC2, &INC3, &INC4, &INT_REL, &ODCNTL, &WUFC, &ATH} {
		if err := reg.updatereg(d); err != nil {
			return err
		}
	}
	opt2 := SetOperatingMode()
	opt2()
	for _, reg := range []*Register{&CNTL1} {
		if err := reg.updatereg(d); err != nil {
			return err
		}
	}

	return nil
}
