package buffers

type CappedFloats struct {
	capacity int
	data     []float64
}

func NewCappedFloats(capacity int) *CappedFloats {
	return &CappedFloats{
		capacity: capacity,
		data:     []float64{},
	}
}

func (c *CappedFloats) Push(v float64) {
	if len(c.data) < c.capacity {
		c.data = append([]float64{v}, c.data...)
	} else {
		c.data = append([]float64{v}, c.data[:len(c.data)-1]...)
	}
}

func (c *CappedFloats) Data() []float64 {
	return c.data
}
