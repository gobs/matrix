package matrix

import (
	"fmt"
	"testing"
)

var (
	m = New[string](10, 10, true)
)

func init() {
	m.Set(0, 0, "1")
	m.Set(0, m.Height()-1, "2")
	m.Set(m.Width()-1, 0, "3")
	m.Set(m.Width()-1, m.Height()-1, "4")
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
