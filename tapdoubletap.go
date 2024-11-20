package accmeter

func SetXNegativeTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset | byte(0x20)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}

func SetDisableXNegativeTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset & ^byte(0x20)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}

func SetXPositiveTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset | byte(0x10)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}

func SetDisableXPositiveTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset & ^byte(0x10)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}

func SetYNegativeTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset | byte(0x08)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}

func SetDisableYNegativeTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset & ^byte(0x08)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}

func SetYPositiveTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset | byte(0x04)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}

func SetDisableYPositiveTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset & ^byte(0x04)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}

func SetZNegativeTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset | byte(0x02)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}

func SetDisableZNegativeTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset & ^byte(0x02)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}

func SetZPositiveTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset | byte(0x01)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}

func SetDisableZPositiveTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset & ^byte(0x01)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}

func SetDisableAllTapIrq() OptFunc {
	return func() {
		offset := byte(0x3F)
		if len(INC3.newValue) != 0 {
			offset = INC3.newValue[0]
		}

		data := offset & ^byte(0x3F)
		INC3.newValue = []byte{data}
		INC3.update = true
	}
}
