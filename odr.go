package accmeter

type ODR byte

const (
	ODR_12_5HZ  ODR = 0x00
	ODR_25HZ    ODR = 0x01
	ODR_50HZ    ODR = 0x02
	ODR_100HZ   ODR = 0x03
	ODR_200HZ   ODR = 0x04
	ODR_400HZ   ODR = 0x05
	ODR_800HZ   ODR = 0x06
	ODR_1600HZ  ODR = 0x07
	ODR_0_781HZ ODR = 0x08
	ODR_1_563HZ ODR = 0x09
	ODR_3_125HZ ODR = 0x0A
	ODR_6_25HZ  ODR = 0x0B
)

func SetOutputDataRate(f ODR) OptFunc {
	return func() {
		offset := byte(0x02)
		if len(ODCNTL.newValue) != 0 {
			offset = ODCNTL.newValue[0]
		}
		data := (offset & byte(0xF0)) | byte(f)
		ODCNTL.newValue = []byte{data}
		ODCNTL.update = true
	}
}
