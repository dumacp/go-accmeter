package accmeter

import "encoding/binary"

// Sensivity is the sensitivity of the accelerometer in g/LSB. The value is
type Sensivity float64

const (
	Sensivity4g Sensivity = 8192.0
	Sensivity2g Sensivity = 16384.0
	Sensivity8g Sensivity = 4096.0
)

func ReadAcceleration(reg []byte, sensitivity Sensivity) float64 {
	// Leer los registros L y H

	buf := make([]byte, 0)
	buf = append(buf, reg...)

	if len(buf) < 16 {
		for range make([]int, 16-len(buf)) {
			buf = append(buf, 0x00)
		}
	}

	// Combinar los registros L y H en un entero de 16 bits
	raw := int16(binary.LittleEndian.Uint16(reg))

	// Convertir a aceleración en g
	return float64(raw) / float64(sensitivity)
}
