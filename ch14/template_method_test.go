package ch14

import "testing"

func TestNewFtoc(t *testing.T) {
	f := NewFtoc()
	f.Run(f)

	m := NewMMY()
	m.Run(m)
}
