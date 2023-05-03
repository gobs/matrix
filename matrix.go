package matrix

type Matrix[T any] struct {
	w, h  int  // width and height
	yztop bool // y zero (true=top, false=bottom)
	cells []T  // array of cells
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

func (m Matrix[T]) Width() int {
	return m.w
}

func (m Matrix[T]) Height() int {
	return m.h
}

func (m Matrix[T]) Get(x, y int) T {
	if !m.yztop {
		y += m.h - 1
	}

	return m.cells[y*m.w+x]
}

func (m Matrix[T]) Set(x, y int, v T) {
	if !m.yztop {
		y += m.h - 1
	}

	m.cells[y*m.w+x] = v
}

func (m Matrix[T]) Row(y int) []T {
	if !m.yztop {
		y += m.h - 1
	}

	p := y * m.w
	return m.cells[p : p+m.w]
}

func (m Matrix[T]) Col(x int) (col []T) {
	for y := 0; y < m.h; y++ {
		col = append(col, m.Get(x, y))
	}

	return col
}
