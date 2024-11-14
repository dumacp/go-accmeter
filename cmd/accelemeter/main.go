package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/dumacp/go-accmeter"
)

var devname string
var loop bool
var reset bool

func init() {
	flag.StringVar(&devname, "devname", "i2c-2", "device name")
	flag.BoolVar(&loop, "loop", false, "read loop HPFilter registers")
	flag.BoolVar(&reset, "reset", false, "sotfware reset")
}

func main() {

	flag.Parse()

	dev := &accmeter.Device{}
	dev.Path = devname

	opts := make([]accmeter.OptFunc, 0)
	opts = append(opts, accmeter.SetOutputDataRate(accmeter.ODR_50HZ))
	opts = append(opts, accmeter.SetOWUF(accmeter.OWUF_25HZ))
	opts = append(opts, accmeter.Set2gAccelerationRange())
	opts = append(opts, accmeter.SetEnableWakeUpFunction())
	opts = append(opts, accmeter.SetWUFC(byte(100)))
	opts = append(opts, accmeter.SetATH(byte(0x0A)))
	opts = append(opts, accmeter.SetPhysicalIrq())
	opts = append(opts, accmeter.SetActiveHighIrq())
	opts = append(opts, accmeter.SetWakeUpIrq())

	if reset {
		if err := dev.Init(accmeter.StartRamReboot()); err != nil {
			log.Fatalf("error init: %v", err)
		}
		fmt.Println("reset device conf")
		if err := dev.NewConf(opts...); err != nil {
			log.Fatalf("error init: %v", err)
		}
	} else if err := dev.Init(opts...); err != nil {
		log.Fatalf("error init: %v", err)
	}

	for _, reg := range []*accmeter.Register{&accmeter.CNTL1, &accmeter.CNTL2, &accmeter.CNTL3, &accmeter.INC1, &accmeter.INC2,
		&accmeter.INC3, &accmeter.INC4, &accmeter.INT_REL, &accmeter.ODCNTL, &accmeter.WUFC, &accmeter.ATH} {
		if _, err := reg.Read(dev); err != nil {
			log.Fatalf("error read: %v", err)
		}
	}

	for _, reg := range []*accmeter.Register{&accmeter.CNTL1, &accmeter.CNTL2, &accmeter.CNTL3, &accmeter.INC1, &accmeter.INC2,
		&accmeter.INC3, &accmeter.INC4, &accmeter.INT_REL, &accmeter.ODCNTL, &accmeter.WUFC, &accmeter.ATH} {
		if v := reg.LastKnownValue(); v != nil {
			fmt.Printf("register: %s, value: %02X\n", reg.Name, v)
		}
	}

	if loop {
		tickHPFiler := time.NewTicker(300 * time.Millisecond)
		defer tickHPFiler.Stop()

		xhp := []byte{0x00}
		yhp := []byte{0x00}
		zhp := []byte{0x00}

		for range tickHPFiler.C {
			if data, err := accmeter.XHP.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else if !bytes.Equal(data, xhp) {
				xhp = data
				fmt.Printf("XHP: %02X\n", data)
			}
			if data, err := accmeter.YHP.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else if !bytes.Equal(data, yhp) {
				yhp = data
				fmt.Printf("YHP: %02X\n", data)
			}
			if data, err := accmeter.ZHP.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else if !bytes.Equal(data, zhp) {
				zhp = data
				fmt.Printf("ZHP: %02X\n", data)
			}
		}
	}
}
