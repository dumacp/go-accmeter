package accmeter

func SetOperatingMode() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(CNTL1.newValue) != 0 {
			offset = CNTL1.newValue[0]
		}
		data := offset | byte(0x80)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func SetStandByMode() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(CNTL1.newValue) != 0 {

			offset = CNTL1.newValue[0]
		}
		data := offset & ^byte(0x80)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func SetLowResolution() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(CNTL1.newValue) != 0 {
			offset = CNTL1.newValue[0]
		}
		data := offset & ^byte(0x40)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func SetHighResolution() OptFunc {
	return func() {
		offset := byte(0x00)
		if len(CNTL1.newValue) != 0 {
			offset = CNTL1.newValue[0]
		}
		data := offset | byte(0x40)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func SetEnableNewAccleratinInterrupt() OptFunc {
	return func() {
		offset := byte(0x98)
		if len(CNTL3.newValue) != 0 {
			offset = CNTL1.newValue[0]
		}
		data := offset | byte(0x20)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func SetDisableNewAccleratinInterrupt() OptFunc {
	return func() {
		offset := byte(0x98)
		if len(CNTL3.newValue) != 0 {
			offset = CNTL1.newValue[0]
		}
		data := offset & ^byte(0x20)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func Set2gAccelerationRange() OptFunc {
	return func() {
		offset := byte(0x98)
		if len(CNTL3.newValue) != 0 {
			offset = CNTL1.newValue[0]
		}
		data := offset & ^byte(0x18)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func Set4gAccelerationRange() OptFunc {
	return func() {
		offset := byte(0x98)
		if len(CNTL3.newValue) != 0 {
			offset = CNTL1.newValue[0]
		}
		data := (offset & ^byte(0x18)) | byte(0x08)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func Set8gAccelerationRange() OptFunc {
	return func() {
		offset := byte(0x98)
		if len(CNTL3.newValue) != 0 {
			offset = CNTL3.newValue[0]
		}
		data := (offset & ^byte(0x18)) | byte(0x10)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func SetEnableDirectionalTap() OptFunc {
	return func() {
		offset := byte(0x98)
		if len(CNTL3.newValue) != 0 {
			offset = CNTL1.newValue[0]
		}
		data := offset | byte(0x04)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func SetDisableDirectionalTap() OptFunc {
	return func() {
		offset := byte(0x98)
		if len(CNTL3.newValue) != 0 {
			offset = CNTL1.newValue[0]
		}
		data := offset & ^byte(0x04)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func SetEnableWakeUpFunction() OptFunc {
	return func() {
		offset := byte(0x98)
		if len(CNTL3.newValue) != 0 {
			offset = CNTL1.newValue[0]
		}
		data := offset | byte(0x02)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func SetDisableWakeUpFunction() OptFunc {
	return func() {
		offset := byte(0x98)
		if len(CNTL3.newValue) != 0 {
			offset = CNTL1.newValue[0]
		}
		data := offset & ^byte(0x02)
		CNTL1.newValue = []byte{data}
		CNTL1.update = true
	}
}

func SetEnableTiltPosition() OptFunc {
	return func() {
		offset := byte(0x98)
		if len(CNTL3.newValue) != 0 {
			offset = CNTL3.newValue[0]
		}
		data := offset | byte(0x01)
		CNTL3.newValue = []byte{data}
		CNTL3.update = true
	}
}

func SetDisableTiltPosition() OptFunc {
	return func() {
		offset := byte(0x98)
		if len(CNTL3.newValue) != 0 {
			offset = CNTL3.newValue[0]
		}
		data := offset & ^byte(0x01)
		CNTL3.newValue = []byte{data}
		CNTL3.update = true
	}
}

func StartRamReboot() OptFunc {
	return func() {
		offset := byte(0x37)
		if len(CNTL2.newValue) != 0 {
			offset = CNTL2.newValue[0]
		}
		data := offset | byte(0x80)

		CNTL2.newValue = []byte{data}
		CNTL2.update = true

	}
}

func SetTestContol() OptFunc {
	return func() {
		offset := byte(0x37)
		if len(CNTL2.newValue) != 0 {
			offset = CNTL2.newValue[0]
		}
		data := offset | byte(0x40)
		CNTL2.newValue = []byte{data}
		CNTL2.update = true
	}
}
