package accmeter

type IRQDetectedMotion byte

const (
	IRQDetectedMotion_X_positive IRQDetectedMotion = 0x01 << 5
	IRQDetectedMotion_X_negative IRQDetectedMotion = 0x01 << 4
	IRQDetectedMotion_Y_positive IRQDetectedMotion = 0x01 << 3
	IRQDetectedMotion_Y_negative IRQDetectedMotion = 0x01 << 2
	IRQDetectedMotion_Z_positive IRQDetectedMotion = 0x01 << 1
	IRQDetectedMotion_Z_negative IRQDetectedMotion = 0x01 << 0
)

func SetDetectedMotion(irq IRQDetectedMotion) OptFunc {
	return func() {
		INC2.newValue = []byte{byte(irq)}
		INC2.update = true
	}
}

func SetDisablePhysicalIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC1.newValue) != 0 {
			offset = INC1.newValue[0]
		}

		data := offset & ^byte(0x20)
		INC1.newValue = []byte{data}
		INC1.update = true
	}
}

func SetPhysicalIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC1.newValue) != 0 {
			offset = INC1.newValue[0]
		}

		data := offset | byte(0x20)
		INC1.newValue = []byte{data}
		INC1.update = true
	}
}

func SetActiveHighIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC1.newValue) != 0 {
			offset = INC1.newValue[0]
		}
		data := offset | byte(0x10)
		INC1.newValue = []byte{data}
		INC1.update = true
	}
}

func SetActiveLowIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC1.newValue) != 0 {
			offset = INC1.newValue[0]
		}
		data := offset & ^byte(0x10)
		INC1.newValue = []byte{data}
		INC1.update = true
	}
}

func SetClearedByIntRelIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC1.newValue) != 0 {
			offset = INC1.newValue[0]
		}
		data := offset & ^byte(0x08)
		INC1.newValue = []byte{data}
		INC1.update = true
	}
}

func SetPulse50msIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC1.newValue) != 0 {
			offset = INC1.newValue[0]
		}
		data := offset | byte(0x08)
		INC1.newValue = []byte{data}
		INC1.update = true
	}
}

func SetBufferFullIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC4.newValue) != 0 {
			offset = INC4.newValue[0]
		}
		data := offset | byte(0x04)
		INC4.newValue = []byte{data}
		INC4.update = true
	}
}

func SetDisableBufferFullIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC4.newValue) != 0 {
			offset = INC4.newValue[0]
		}
		data := offset & ^byte(0x04)
		INC4.newValue = []byte{data}
		INC4.update = true
	}
}

func SetWattermarkIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC4.newValue) != 0 {
			offset = INC4.newValue[0]
		}
		data := offset | byte(0x02)
		INC4.newValue = []byte{data}
		INC4.update = true
	}
}

func SetDisableWattermarkIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC4.newValue) != 0 {
			offset = INC4.newValue[0]
		}
		data := offset & ^byte(0x02)
		INC4.newValue = []byte{data}
		INC4.update = true
	}
}

func SetDataReadyIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC4.newValue) != 0 {
			offset = INC4.newValue[0]
		}
		data := offset | byte(0x01)
		INC4.newValue = []byte{data}
		INC4.update = true
	}
}

func SetDisableDataReadyIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC4.newValue) != 0 {
			offset = INC4.newValue[0]
		}
		data := offset & ^byte(0x01)
		INC4.newValue = []byte{data}
		INC4.update = true
	}
}

func SetTapDoubleTapIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC4.newValue) != 0 {
			offset = INC4.newValue[0]
		}
		data := offset | byte(0x04)
		INC4.newValue = []byte{data}
		INC4.update = true
	}
}

func SetDisableTapDoubleTapIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC4.newValue) != 0 {
			offset = INC4.newValue[0]
		}
		data := offset & ^byte(0x04)
		INC4.newValue = []byte{data}
		INC4.update = true
	}
}

func SetWakeUpIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC4.newValue) != 0 {
			offset = INC4.newValue[0]
		}
		data := offset | byte(0x02)
		INC4.newValue = []byte{data}
		INC4.update = true
	}
}

func SetDisableWakeUpIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC4.newValue) != 0 {
			offset = INC4.newValue[0]
		}
		data := offset & ^byte(0x02)
		INC4.newValue = []byte{data}
		INC4.update = true
	}
}

func SetTiltPositionIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC4.newValue) != 0 {
			offset = INC4.newValue[0]
		}
		data := offset | byte(0x01)
		INC4.newValue = []byte{data}
		INC4.update = true
	}
}

func SetDisableTiltPositionIrq() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(INC4.newValue) != 0 {
			offset = INC4.newValue[0]
		}
		data := offset & ^byte(0x01)
		INC4.newValue = []byte{data}
		INC4.update = true
	}
}
