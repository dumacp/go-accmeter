package accmeter

import "fmt"

type Register struct {
	Name      string `json:"name"`
	Addr      byte   `json:"addr"`
	Len       byte   `json:"len"`
	lastValue []byte
	newValue  []byte
	update    bool
}

func (r *Register) Read(d *Device) ([]byte, error) {
	d.mux.Lock()
	defer d.mux.Unlock()
	buf := make([]byte, r.Len)
	if err := d.dev.Tx([]byte{r.Addr}, buf); err != nil {
		return nil, err
	}
	r.lastValue = buf
	return buf, nil
}

func (r *Register) Write(d *Device, data []byte) error {
	d.mux.Lock()
	defer d.mux.Unlock()
	buf := make([]byte, 0)
	buf = append(buf, r.Addr)
	buf = append(buf, data...)

	if _, err := d.dev.Write(buf); err != nil {
		return err
	}
	r.lastValue = data
	return nil
}

func (r *Register) updatereg(d *Device) error {
	// fmt.Println("update register")
	if r.update {

		d.mux.Lock()
		defer d.mux.Unlock()
		buf := make([]byte, 0)
		buf = append(buf, r.Addr)
		buf = append(buf, r.newValue...)
		fmt.Printf("write register: %s, value:[ % 02X ]\n", r.Name, buf)
		if _, err := d.dev.Write(buf); err != nil {
			return err
		}
		r.lastValue = r.newValue
		r.update = false
	}
	return nil
}

func (r *Register) LastKnownValue() []byte {
	return r.lastValue
}
