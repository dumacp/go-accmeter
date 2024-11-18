package accmeter

import (
	"fmt"
	"time"
)

type OWUF byte

const (
	OWUF_0_781HZ OWUF = 0x00
	OWUF_1_563HZ OWUF = 0x01
	OWUF_3_125HZ OWUF = 0x02
	OWUF_6_25HZ  OWUF = 0x03
	OWUF_12_5HZ  OWUF = 0x04
	OWUF_25HZ    OWUF = 0x05
	OWUF_50HZ    OWUF = 0x06
	OWUF_100HZ   OWUF = 0x07
)

func (t OWUF) Hz() float64 {
	switch t {
	case OWUF_0_781HZ:
		return 0.781
	case OWUF_1_563HZ:
		return 1.563
	case OWUF_3_125HZ:
		return 3.125
	case OWUF_6_25HZ:
		return 6.25
	case OWUF_12_5HZ:
		return 12.5
	case OWUF_25HZ:
		return 25
	case OWUF_50HZ:
		return 50
	case OWUF_100HZ:
		return 100
	default:
		return 0.781
	}
}

func SetOWUF(f OWUF) OptFunc {
	return func() {
		offset := byte(0x98)
		if len(CNTL3.newValue) != 0 {
			offset = CNTL3.newValue[0]
		}
		data := (offset & byte(0xF8)) | byte(f)
		CNTL3.newValue = []byte{data}
		CNTL3.update = true
	}
}

func SetWUFC(delaysec time.Duration, frec OWUF) OptFunc {
	return func() {
		fmt.Println("SetWUFC")
		count := byte(delaysec.Seconds() * float64(frec.Hz()))
		WUFC.newValue = []byte{count}
		fmt.Printf("WUFC.newValue: %v\n", WUFC.newValue)
		WUFC.update = true
	}
}

func SetATH(threshold float64) OptFunc {
	return func() {
		ath := byte(threshold * 16)
		ATH.newValue = []byte{ath}
		ATH.update = true
	}
}
