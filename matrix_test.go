package matrix

import (
	"fmt"
	"sort"
	"testing"
)

var (
	m  = New[string](10, 10, false)
	m2 = New[string](3, 3, false)

	center = []Cell[string]{
		{X: 0, Y: 0},
		{X: 0, Y: 1},
		{X: 0, Y: 2},
		{X: 1, Y: 0},
		{X: 1, Y: 2},
		{X: 2, Y: 0},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
	}

	left = []Cell[string]{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 1, Y: 1},
		{X: 0, Y: 2},
		{X: 1, Y: 2},
	}

	left_r = []Cell[string]{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 1, Y: 1},
		{X: 0, Y: 2},
		{X: 1, Y: 2},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
		{X: 2, Y: 0},
	}

	right = []Cell[string]{
		{X: 1, Y: 2},
		{X: 2, Y: 2},
		{X: 1, Y: 1},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
	}

	right_r = []Cell[string]{
		{X: 1, Y: 2},
		{X: 2, Y: 2},
		{X: 1, Y: 1},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 0, Y: 2},
		{X: 0, Y: 1},
		{X: 0, Y: 0},
	}

	bottomleft = []Cell[string]{
		{X: 0, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 0},
	}

	topright = []Cell[string]{
		{X: 1, Y: 2},
		{X: 1, Y: 1},
		{X: 2, Y: 1},
	}

	bottomleft_r = []Cell[string]{
		{X: 0, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 0},
		{X: 2, Y: 1},
		{X: 2, Y: 0},
		{X: 2, Y: 2},
		{X: 0, Y: 2},
		{X: 1, Y: 2},
	}

	topright_r = []Cell[string]{
		{X: 1, Y: 2},
		{X: 1, Y: 1},
		{X: 2, Y: 1},
		{X: 0, Y: 1},
		{X: 0, Y: 2},
		{X: 0, Y: 0},
		{X: 2, Y: 0},
		{X: 1, Y: 0},
	}
)

func init() {
	m.Set(0, 0, "1")
	m.Set(0, m.Height()-1, "2")
	m.Set(m.Width()-1, 0, "3")
	m.Set(m.Width()-1, m.Height()-1, "4")

	m2.Set(0, 0, "TL")
	m2.Set(0, 1, "ML")
	m2.Set(0, 2, "BL")
	m2.Set(1, 0, "TM")
	m2.Set(1, 1, "MM")
	m2.Set(1, 2, "BM")
	m2.Set(2, 0, "TR")
	m2.Set(2, 1, "MR")
	m2.Set(2, 2, "BR")
}

func TestNew(t *testing.T) {
	fmt.Printf("%#v\n", m)
}

func TestByRows(t *testing.T) {
	for i := 0; i < m.Height(); i++ {
		fmt.Println(m.Row(i))
	}
}

func TestByColumns(t *testing.T) {
	for i := 0; i < m.Width(); i++ {
		fmt.Println(m.Col(i))
	}
}

func TestClone(t *testing.T) {
	m2 := m.Clone()
	fmt.Printf("%#v\n", m2)
}

func TestFix(t *testing.T) {
	fmt.Printf("fix 0: %v\n", m.Fix(0))
	fmt.Printf("unfix 9: %v\n", m.Fix(9))
}

func TestFill(t *testing.T) {
	m2 := m.Clone()
	fmt.Printf("%#v\n", m2)

	m2.Fill(".")
	fmt.Printf("%#v\n", m2)
}

func compare(l1, l2 []Cell[string]) bool {
	if len(l1) != len(l2) {
		fmt.Println("XXX compare length", l1, l2)
		return false
	}

	sort.Slice(l1, func(i, j int) bool {
		v1 := l1[i].X*3 + l1[i].Y
		v2 := l1[j].X*3 + l1[j].Y

		return v1 < v2
	})

	sort.Slice(l2, func(i, j int) bool {
		v1 := l2[i].X*3 + l2[i].Y
		v2 := l2[j].X*3 + l2[j].Y

		return v1 < v2
	})

	for i := range l1 {
		if l1[i].X != l2[i].X || l1[i].Y != l2[i].Y {
			fmt.Println("XXX compare", i, l1, l2)
			return false
		}
	}

	return true
}

func TestAdjacent(t *testing.T) {
	cells := m2.Adjacent(1, 1, false)
	if !compare(cells, center) {
		t.Logf("(1,1) %v %#v\n", len(cells), cells)
		t.Fail()
	}

	cells = m2.Adjacent(0, 1, false)
	if !compare(cells, left) {
		t.Logf("(0,1) %v %#v\n", len(cells), cells)
		t.Fail()
	}

	cells = m2.Adjacent(2, 1, false)
	if !compare(cells, right) {
		t.Logf("(2,1) %v %#v\n", len(cells), cells)
		t.Fail()
	}

	cells = m2.Adjacent(0, 0, false)
	if !compare(cells, bottomleft) {
		t.Logf("(0,0) %v %#v\n", len(cells), cells)
		t.Fail()
	}

	cells = m2.Adjacent(2, 2, false)
	if !compare(cells, topright) {
		t.Logf("(2,2) %v %#v\n", len(cells), cells)
		t.Fail()
	}
}

func TestAdjacentRoll(t *testing.T) {
	cells := m2.Adjacent(1, 1, true)
	if !compare(cells, center) {
		t.Logf("(1,1) %v %#v\n", len(cells), cells)
		t.Fail()
	}

	cells = m2.Adjacent(0, 1, true)
	if !compare(cells, left_r) {
		t.Logf("(0,1) %v %#v\n", len(cells), cells)
		t.Fail()
	}

	cells = m2.Adjacent(2, 1, true)
	if !compare(cells, right_r) {
		t.Logf("(2,1) %v %#v\n", len(cells), cells)
		t.Fail()
	}

	cells = m2.Adjacent(0, 0, true)
	if !compare(cells, bottomleft_r) {
		t.Logf("(0,0) %v %#v\n", len(cells), cells)
		t.Fail()
	}

	cells = m2.Adjacent(2, 2, true)
	if !compare(cells, topright_r) {
		t.Logf("(2,2) %v %#v\n", len(cells), cells)
		t.Fail()
	}
}
