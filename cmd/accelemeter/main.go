package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/dumacp/go-accmeter"
	ui "github.com/gizak/termui/v3"
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
	opts = append(opts, accmeter.SetOWUF(accmeter.OWUF_100HZ))
	opts = append(opts, accmeter.Set2gAccelerationRange())
	opts = append(opts, accmeter.SetEnableWakeUpFunction())
	opts = append(opts, accmeter.SetWUFC(3*time.Second, accmeter.OWUF_25HZ))
	opts = append(opts, accmeter.SetATH(0.5))
	opts = append(opts, accmeter.SetPhysicalIrq())
	opts = append(opts, accmeter.SetActiveHighIrq())
	opts = append(opts, accmeter.SetWakeUpIrq())
	opts = append(opts, accmeter.SetDisableAllTapIrq())
	// opts = append(opts, accmeter.SetDisableAllMotionDetectionIrq())
	// opts = append(opts, accmeter.SetYPositiveMotionDetectionIrq(), accmeter.SetYNegativeMotionDetectionIrq())
	// opts = append(opts, accmeter.SetZPositiveMotionDetectionIrq(), accmeter.SetZNegativeMotionDetectionIrq())

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
		tickHPFiler := time.NewTicker(100 * time.Millisecond)
		defer tickHPFiler.Stop()

		xhp := []byte{0x00}
		yhp := []byte{0x00}
		zhp := []byte{0x00}

		xout := []byte{0x00}
		yout := []byte{0x00}
		zout := []byte{0x00}

		readx := 300.0
		ready := 300.0
		readz := 300.

		zoutbuf := accmeter.NewCircularBuffer(40)
		youtbuf := accmeter.NewCircularBuffer(40)
		xoutbuf := accmeter.NewCircularBuffer(40)

		sens := accmeter.Sensivity2g

		for range tickHPFiler.C {
			tx := time.Now().UnixMilli()
			if data, err := accmeter.XHP.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else if !bytes.Equal(data, xhp) {
				read := accmeter.ReadAcceleration(data, sens)
				xhp = data
				if math.Abs(read) > 0.1 {
					log.Printf("XHP: %02X (%.2f)\n", data, read)
				}
			}
			if data, err := accmeter.XOUT.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else {
				read := accmeter.ReadAcceleration(data, sens)
				xoutbuf.Add(float64(tx), read)
				if !bytes.Equal(data, xout) {
					xout = data
					if math.Abs(read-readx) > 0.5 {
						readx = read
						log.Printf("XOUT: %02X (%.2f)\n", data, read)
					}
				}
			}

			if data, err := accmeter.YHP.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else if !bytes.Equal(data, yhp) {
				yhp = data
				read := accmeter.ReadAcceleration(data, sens)
				if math.Abs(read) > 0.1 {
					log.Printf("YHP: %02X (%.2f)\n", data, read)
				}
			}
			if data, err := accmeter.YOUT.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else {
				read := accmeter.ReadAcceleration(data, sens)
				youtbuf.Add(float64(tx), read)
				if !bytes.Equal(data, yout) {
					yout = data

					if math.Abs(read-ready) > 0.5 {
						ready = read
						log.Printf("YOUT: %02X (%.2f)\n", data, read)
					}
				}
			}

			if data, err := accmeter.ZHP.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else if !bytes.Equal(data, zhp) {
				zhp = data
				read := accmeter.ReadAcceleration(data, sens)
				if math.Abs(read) > 0.1 {
					log.Printf("ZHP: %02X (%.2f)\n", data, read)
				}
			}
			if data, err := accmeter.ZOUT.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else {
				read := accmeter.ReadAcceleration(data, sens)
				zoutbuf.Add(float64(tx), read)
				if !bytes.Equal(data, zout) {
					zout = data
					if math.Abs(read-readz) > 0.5 {
						readz = read
						log.Printf("ZOUT: %02X (%.2f)\n", data, read)
					}
				}
			}

			if data, err := accmeter.INS1.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else if data[0] != 0x00 {
				log.Printf("INS1: %02X\n", data)
			}
			if data, err := accmeter.INS2.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else if data[0] != 0x00 {
				log.Printf("INS2: %02X\n", data)
			}
			if data, err := accmeter.INS3.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else if data[0] != 0x00 {
				log.Printf("INS3: %02X\n", data)
			}
			if data, err := accmeter.STATUS_REG.Read(dev); err != nil {
				log.Fatalf("error read: %v", err)
			} else if data[0] != 0x00 {
				log.Printf("STATUS_REG: %02X\n", data)
				break
			}

		}

		fmt.Println("////////////////////////////// FINAL DATA //////////////////////////////")

		xdata := xoutbuf.Get()
		// fmt.Printf("XOUT: %v\n", xdata)
		tx0 := xdata[0][0]
		for i, v := range xdata {
			fmt.Printf("%d - XOUT[%.03f]: %v\n", i, (v[0]-tx0)/1000, v[1])
		}
		fmt.Println("//////////////////////////////")
		ydata := youtbuf.Get()
		ty0 := ydata[0][0]
		for i, v := range ydata {
			fmt.Printf("%d - YOUT[%.03f]: %v\n", i, (v[0]-ty0)/1000, v[1])
		}
		fmt.Println("//////////////////////////////")
		zdata := zoutbuf.Get()
		tz0 := zdata[0][0]
		for i, v := range zdata {
			fmt.Printf("%d - ZOUT[%.03f]: %v\n", i, (v[0]-tz0)/1000, v[1])
		}
		fmt.Println("//////////////////////////////")

		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()

		xPlot, _ := accmeter.Graph("X", []int{0, 0, 30, 10}, xdata)
		yPlot, _ := accmeter.Graph("Y", []int{0, 10, 30, 20}, ydata)
		zPlot, _ := accmeter.Graph("Z", []int{30, 0, 60, 20}, zdata)

		ui.Render(xPlot, yPlot, zPlot)

		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		}

	}

}
