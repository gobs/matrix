package matrix

import (
	"fmt"
	"testing"
)

var (
	m  = New[string](10, 10, false)
	m2 = New[string](3, 3, false)
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

func TestAdjactent(t *testing.T) {
	cells := m2.Adjacent(1, 1, false)
	fmt.Printf("(1,1) %v %#v\n", len(cells), cells)

	cells = m2.Adjacent(0, 1, false)
	fmt.Printf("(0,1) %v %#v\n", len(cells), cells)

	cells = m2.Adjacent(2, 1, false)
	fmt.Printf("(2,1) %v %#v\n", len(cells), cells)

	cells = m2.Adjacent(0, 0, false)
	fmt.Printf("(1,0) %v %#v\n", len(cells), cells)

	cells = m2.Adjacent(2, 2, false)
	fmt.Printf("(2,2) %v %#v\n", len(cells), cells)
}

func TestAdjactentRoll(t *testing.T) {
	cells := m2.Adjacent(1, 1, true)
	fmt.Printf("(1,1) %v %#v\n", len(cells), cells)

	cells = m2.Adjacent(0, 1, true)
	fmt.Printf("(0,1) %v %#v\n", len(cells), cells)

	cells = m2.Adjacent(2, 1, true)
	fmt.Printf("(2,1) %v %#v\n", len(cells), cells)

	cells = m2.Adjacent(0, 0, false)
	fmt.Printf("(1,0) %v %#v\n", len(cells), cells)

	cells = m2.Adjacent(2, 2, false)
	fmt.Printf("(2,2) %v %#v\n", len(cells), cells)
}
