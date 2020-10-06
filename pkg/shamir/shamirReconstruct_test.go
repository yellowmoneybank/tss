package shamir

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShamirReconstruct(t *testing.T) {
	secret := []byte{10}

	splits, _ := SplitSecret(secret, 3, 3)
	reconstruct, _ := Reconstruct(splits)
	assert.Equal(t, secret, reconstruct)
	fmt.Println(string(reconstruct))
}

func Test_isDeterminantVandermondeZero(t *testing.T) {
	{
		indices := []int{1, 3, 52, 7, 9, -10}
		assert.False(t, isDeterminantVandermondeZero(indices))
	}
	{
		indices := []int{1, 3, 52, 7, 9, -10, 1}
		assert.True(t, isDeterminantVandermondeZero(indices))
	}
}

func Test_reconstructPolynomial(t *testing.T) {
	{
		//x^2
		var points []point
		for i := -50; i < 100; i++ {
			points = append(points, point{
				x: i,
				y: i * i,
			})
		}

		assert.Equal(t, float64(123*123), reconstructPolynom(points)(123), "should be the same")
	}
}

func Test_createBasisPolynomial(t *testing.T) {
	// x^2
	var points []point
	for i := 1; i < 4; i++ {
		points = append(points, point{
			x: i,
			y: i * i,
		})
	}
	basisPolynomial := createBasisPolynomial(points, 0)
	shouldBePolynomial := func(x float64) float64 {
		return 0.5 * ((x * x) - 5*x + 6)
	}
	assert.Equal(t, shouldBePolynomial(5), basisPolynomial(5))

}
