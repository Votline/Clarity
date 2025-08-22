package ui

const (
	minX = -0.3
	maxX = -0.2

	maxY = 0.4
	minY = -0.2
)

func (d *digit) create(offset float32) element {
	switch d.num {
	case 1:
		return &digit{
			vtc: []float32{
				minX + offset, maxY/3, 0.0,
				maxX + offset, maxY,   0.0,
				maxX + offset, minY,   0.0},
			vtq: int32(3)}
	case 2:
		return &digit{
			vtc: []float32{
				minX          + offset, maxY/3, 0.0,
				(maxX+minX)/2 + offset, maxY, 0.0,
				maxX          + offset, maxY/3, 0.0,

				minX + offset, minY, 0.0,
				maxX + offset, minY, 0.0},
			vtq: int32(5)}
	case 3:
		return &digit{
			vtc: []float32{
				minX + offset, maxY,   0.0,
				maxX + offset, maxY,   0.0,

				maxX + offset, maxY/3, 0.0,
				minX + offset, maxY/3, 0.0,
				maxX + offset, maxY/3, 0.0,
				
				maxX + offset, minY,   0.0,
				minX + offset, minY,   0.0,},
			vtq: int32(7)}
	case 4:
		return &digit{
			vtc: []float32{
				minX + offset, maxY,   0.0,
				minX + offset, maxY/3, 0.0,
				
				maxX + offset, maxY/3, 0.0,

				maxX + offset, maxY,   0.0,
				maxX + offset, minY,   0.0},
			vtq: int32(5)}
	case 5:
		return &digit{
			vtc: []float32{
				maxX + offset, maxY,   0.0,
				minX + offset, maxY,   0.0,

				minX + offset, maxY/3, 0.0,
				maxX + offset, maxY/3, 0.0,
				
				maxX + offset, minY,   0.0,
				minX + offset, minY,   0.0},
			vtq: int32(6)}
	case 6:
		return &digit{
			vtc: []float32{
				maxX + offset, maxY,   0.0,
				minX + offset, maxY,   0.0,
				
				minX + offset, minY,   0.0,
				maxX + offset, minY,   0.0,

				maxX + offset, maxY/3, 0.0,
				minX + offset, maxY/3, 0.0},
			vtq: int32(6)}
	case 7:
		return &digit{
			vtc: []float32{
				minX-0.05 + offset, maxY, 0.0,
				maxX      + offset, maxY, 0.0,
				minX      + offset, minY, 0.0},
			vtq: int32(3)}
	case 8:
		return &digit{
			vtc: []float32{
				minX + offset, maxY,   0.0,
				maxX + offset, maxY,   0.0,
				
				maxX + offset, minY,   0.0,
				minX + offset, minY,   0.0,
				
				minX + offset, maxY,   0.0,
				minX + offset, maxY/3, 0.0,
				maxX + offset, maxY/3, 0.0},
			vtq: int32(7)}
	case 9:
		return &digit{
			vtc: []float32{
				maxX + offset, maxY,   0.0,
				minX + offset, maxY,   0.0,
				
				minX + offset, maxY/3, 0.0,
				maxX + offset, maxY/3, 0.0,

				maxX + offset, maxY,   0.0,
				maxX + offset, minY,   0.0},
			vtq: int32(6)}
	case 0:
		return &digit{
			vtc: []float32{
				minX + offset, maxY, 0.0,
				maxX + offset, maxY, 0.0,
				
				maxX + offset, minY, 0.0,
				minX + offset, minY, 0.0,

				minX + offset, maxY, 0.0},
			vtq: int32(5)}
	}
	return &digit{}
}

func (d *digit) getVtc() []float32 {
	return d.vtc
}

func (d *digit) getVtq() int32 {
	return d.vtq
}
