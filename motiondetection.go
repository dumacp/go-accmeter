package accmeter

import "fmt"

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

func SetWUFC(count byte) OptFunc {
	return func() {
		fmt.Println("SetWUFC")
		WUFC.newValue = []byte{count}
		fmt.Printf("WUFC.newValue: %v\n", WUFC.newValue)
		WUFC.update = true
	}
}

func SetATH(ath byte) OptFunc {
	return func() {
		ATH.newValue = []byte{ath}
		ATH.update = true
	}
}
