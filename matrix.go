package matrix

// A Matrix is a 2-dimensional table stored as a 1-dimensional slice
type Matrix[T comparable] struct {
	w, h  int  // width and height
	cart  bool // true: cartesian coordinates (0 at the bottom), false: computer coordinates (0 at the top)
	cells []T  // array of cells
}

// A Cell stores the cell coordinates and its value
type Cell[T any] struct {
	X, Y  int
	Value T
}

// New creates a new Matrix given width, height and if it should use cartesian coordinates (0 at the bottom)
func New[T comparable](w, h int, cart bool) Matrix[T] {
	return Matrix[T]{
		w:     w,
		h:     h,
		cart:  cart,
		cells: make([]T, w*h),
	}
}

// NewLike creates a new matrix with the same dimensions as the input matrix
func NewLike[T comparable](m Matrix[T]) Matrix[T] {
	return Matrix[T]{
		w:     m.w,
		h:     m.h,
		cart:  m.cart,
		cells: make([]T, m.w*m.h),
	}
}

// FromSlice creates a new matrix with the content of the input slice.
// It requires the matrix `width` (cols) and the type of matrix.
// If the length of the input slice is not a multiple of the width the `ok` return boolean will be false.
func FromSlice[T comparable](cols int, cart bool, sl []T) (m Matrix[T], ok bool) {
	rows := len(sl) / cols
	if rows*cols != len(sl) {
		return m, false
	}

	return Matrix[T]{
		w:     cols,
		h:     rows,
		cart:  cart,
		cells: sl,
	}, true
}

// Clone creates a new matrix that is a copy of the input matrix
func (m Matrix[T]) Clone() Matrix[T] {
	n := NewLike(m)
	copy(n.cells, m.cells)
	return n
}

// Copy copies the matrix `n` into `m` starting at the coordinates `x,y`
func (m Matrix[T]) Copy(x, y int, n Matrix[T]) {
	if x >= m.w || y >= m.h {
		return
	}

	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}

	w, h := n.w, n.h

	if x+w > m.w {
		w = m.w - x
	}
	if y+h > m.h {
		h = m.h - y
	}

	for i := 0; i < h; i++ {
		sr := n.Row(i)
		dr := m.Row(y + i)
		copy(dr[x:x+w], sr[:w])
	}
}

// Creates a new matrix that is a subset of the input matrix
func (m Matrix[T]) Submatrix(x, y, w, h int) Matrix[T] {
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	if w+x > m.w {
		w = m.w - x
	}
	if h+y > m.h {
		h = m.h - y
	}

	s := New[T](w, h, m.cart)

	for i := 0; i < h; i++ {
		sr := m.Row(y + i)
		copy(s.Row(i), sr[x:x+w])
	}

	return s
}

// Fill sets all cells of the matrix to the specified value
func (m Matrix[T]) Fill(v T) {
	for i := range m.cells {
		m.cells[i] = v
	}
}

// Fix converts the row number (y) from one coordinate system to the other (cartesian to computer or viceversa)
func (m Matrix[T]) Fix(y int) int {
	if m.cart {
		return m.h - 1 - y
	}

	return y
}

// Width returns the matrix width (number of columns)
func (m Matrix[T]) Width() int {
	return m.w
}

// Height returns the matrix height (number of rows)
func (m Matrix[T]) Height() int {
	return m.h
}

// Cartesian return true for cartesian coordinates (0 at the bottom) and false for computer coordinates (0 at the top)
func (m Matrix[T]) Cartesian() bool {
	return m.cart
}

// Get returns the value for the cell at x,y
func (m Matrix[T]) Get(x, y int) T {
	y = m.Fix(y)
	return m.cells[y*m.w+x]
}

// Set changes the value for the cell at x,y
func (m Matrix[T]) Set(x, y int, v T) {
	y = m.Fix(y)
	m.cells[y*m.w+x] = v
}

// Row returns the list of values for the specified row
func (m Matrix[T]) Row(y int) []T {
	p := m.Fix(y) * m.w
	return m.cells[p : p+m.w]
}

// Column returns the list of values for the specified column
func (m Matrix[T]) Col(x int) (col []T) {
	for y := 0; y < m.h; y++ {
		col = append(col, m.Get(x, y))
	}

	return col
}

// Slice return the backing slice for the Matrix
func (m Matrix[T]) Slice() []T {
	return m.cells
}

// Equal returns true if the two matrices are equal
func (m Matrix[T]) Equals(n Matrix[T]) bool {
	if m.w != n.w || m.h != n.h || m.cart != n.cart {
		return false
	}

	for i := range m.cells {
		if m.cells[i] != n.cells[i] {
			return false
		}
	}

	return true
}

// Adjacent returns a list of Cell(s) adjacent to the one at x,y.
// If wrap is true coordinates that are outside the Matrix boundary will wrap around.
// For example if x=0 return add rightmost cell, if y=top add bottom cell.
func (m Matrix[T]) Adjacent(x, y int, wrap bool) []Cell[T] {
	var cells []Cell[T]
	var skipxl, skipxr, skipyd, skipyu bool

	xl := x - 1
	if xl < 0 {
		if wrap {
			xl = m.w - 1
		} else {
			skipxl = true
		}
	}

	xr := x + 1
	if xr >= m.w {
		if wrap {
			xr = 0
		} else {
			skipxr = true
		}
	}

	yd := y - 1
	if yd < 0 {
		if wrap {
			yd = m.h - 1
		} else {
			skipyd = true
		}
	}

	yu := y + 1
	if yu >= m.h {
		if wrap {
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
