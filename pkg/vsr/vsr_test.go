package vsr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidVSRShare(t *testing.T) {
}

func Test_vsrCheck(t *testing.T) {
}

func Test_extractIndizesFromMap(t *testing.T) {
	m := make(map[uint16]uint16)
	m[2] = 2
	m[4] = 7
	m[6] = 2

	indices := extractIndizesFromMap(m)

	assert.Equal(t, 3, len(indices), "there should be 3 indices")
	assert.Contains(t, indices, uint16(2))
	assert.Contains(t, indices, uint16(4))
	assert.Contains(t, indices, uint16(6))
}

func Test_initializeExponents(t *testing.T) {
	indices := []uint16{2, 3, 1}
	x := initializeExponents(indices)

	exp1, _ := x[1].Float64()
	assert.Equal(t, float64(3), exp1, "should be equal")

	exp2, _ := x[2].Float64()
	assert.Equal(t, float64(-3), exp2, "should be equal")

	exp3, _ := x[3].Float64()
	assert.Equal(t, float64(1), exp3, "should be equal")
}
