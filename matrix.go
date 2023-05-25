package matrix

type Matrix[T any] struct {
	w, h  int  // width and height
	yztop bool // y zero (true=top, false=bottom)
	cells []T  // array of cells
}

type Cell[T any] struct {
	X, Y  int
	Value T
}

func New[T any](w, h int, yztop bool) Matrix[T] {
	m := Matrix[T]{
		w:     w,
		h:     h,
		yztop: yztop,
		cells: make([]T, w*h),
	}

	return m
}

func (m Matrix[T]) Clone() Matrix[T] {
	n := Matrix[T]{
		w:     m.w,
		h:     m.h,
		yztop: m.yztop,
		cells: make([]T, m.w*m.h),
	}

	copy(n.cells, m.cells)
	return n
}

func (m Matrix[T]) Fill(v T) {
	for i := range m.cells {
		m.cells[i] = v
	}
}

func (m Matrix[T]) Fix(y int) int {
	if !m.yztop {
		return m.h - 1 - y
	}

	return y
}

func (m Matrix[T]) Width() int {
	return m.w
}

func (m Matrix[T]) Height() int {
	return m.h
}

func (m Matrix[T]) Get(x, y int) T {
	y = m.Fix(y)
	return m.cells[y*m.w+x]
}

func (m Matrix[T]) Set(x, y int, v T) {
	y = m.Fix(y)
	m.cells[y*m.w+x] = v
}

func (m Matrix[T]) Row(y int) []T {
	p := m.Fix(y) * m.w
	return m.cells[p : p+m.w]
}

func (m Matrix[T]) Col(x int) (col []T) {
	for y := 0; y < m.h; y++ {
		col = append(col, m.Get(x, y))
	}

	return col
}

func (m Matrix[T]) Adjacent(x, y int, roll bool) []Cell[T] {
	var cells []Cell[T]
	var skipxl, skipxr, skipyd, skipyu bool

	xl := x - 1
	if xl < 0 {
		if roll {
			xl = m.w - 1
		} else {
			skipxl = true
		}
	}

	xr := x + 1
	if xr >= m.w {
		if roll {
			xr = 0
		} else {
			skipxr = true
		}
	}

	yd := y - 1
	if yd < 0 {
		if roll {
			yd = m.h - 1
		} else {
			skipyd = true
		}
	}

	yu := y + 1
	if yu >= m.h {
		if roll {
			yu = 0
		} else {
			skipyu = true
		}
	}

	// x-1
	if !skipxl {
		// x-1, y+1
		if !skipyu {
			cells = append(cells, Cell[T]{X: xl, Y: yu, Value: m.Get(xl, yu)})
		}

		// x-1, y
		cells = append(cells, Cell[T]{X: xl, Y: y, Value: m.Get(xl, y)})

		// x-1, y-1
		if !skipyd {
			cells = append(cells, Cell[T]{X: xl, Y: yd, Value: m.Get(xl, yd)})
		}
	}

	// x, y+1
	if !skipyu {
		cells = append(cells, Cell[T]{X: x, Y: yu, Value: m.Get(x, yu)})
	}

	// x, y-1
	if !skipyd {
		cells = append(cells, Cell[T]{X: x, Y: yd, Value: m.Get(x, yd)})
	}

	// x+1
	if !skipxr {
		// x+1, y+1
		if !skipyu {
			cells = append(cells, Cell[T]{X: xr, Y: yu, Value: m.Get(xr, yu)})
		}

		// x+1, y
		cells = append(cells, Cell[T]{X: xr, Y: y, Value: m.Get(xr, y)})

		// x+1, y-1
		if !skipyd {
			cells = append(cells, Cell[T]{X: xr, Y: yd, Value: m.Get(xr, yd)})
		}
	}

	return cells
}
