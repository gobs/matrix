package matrix

import (
	"fmt"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	m := New[string](10, 10, true)

	m.Set(0, 0, "1")
	m.Set(0, m.Height()-1, "2")
	m.Set(m.Width()-1, 0, "3")
	m.Set(m.Width()-1, m.Height()-1, "4")
	fmt.Printf("%#v\n\n", m)

	for i := 0; i < m.Height(); i++ {
		fmt.Println(m.Row(i))
	}

	fmt.Println()

	for i := 0; i < m.Width(); i++ {
		fmt.Println(m.Col(i))
	}
}
