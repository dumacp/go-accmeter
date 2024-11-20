package accmeter

// CircularBuffer es una estructura para manejar los datos
type CircularBuffer struct {
	data  [][2]float64
	size  int
	start int
	count int
}

// NewCircularBuffer crea un nuevo buffer circular con un tamaño fijo
func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		data: make([][2]float64, size),
		size: size,
	}
}

// Add agrega un nuevo valor al buffer, descartando el más viejo si está lleno
func (cb *CircularBuffer) Add(x, y float64) {
	if cb.count < cb.size {
		cb.data[(cb.start+cb.count)%cb.size] = [2]float64{x, y}
		cb.count++
	} else {
		cb.data[cb.start] = [2]float64{x, y} // Sobreescribe el valor más antiguo
		cb.start = (cb.start + 1) % cb.size  // Mueve el inicio hacia el siguiente
	}
}

// Get devuelve los datos actuales del buffer en orden
func (cb *CircularBuffer) Get() [][2]float64 {
	result := make([][2]float64, cb.count)
	for i := 0; i < cb.count; i++ {
		result[i] = cb.data[(cb.start+i)%cb.size]
	}
	return result
}

/**
func main() {
	// Crear un buffer circular para 30 registros
	buffer := NewCircularBuffer(30)

	// Simular la adición de 35 registros
	for i := 0; i < 35; i++ {
		buffer.Add(float64(i))
		fmt.Printf("Añadiendo: %v, Buffer: %v\n", i, buffer.Get())
	}

	// Imprimir el buffer final
	fmt.Println("Buffer final:", buffer.Get())
}
/**/
