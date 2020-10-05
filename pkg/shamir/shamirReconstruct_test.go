package shamir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reconstructPolynomial(t *testing.T) {
	{
		//x^2
		var points []point
		for i := 0; i < 100; i++ {
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
