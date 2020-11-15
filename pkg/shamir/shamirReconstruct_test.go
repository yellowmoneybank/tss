package shamir

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShamirReconstruct(t *testing.T) {
	{
		secret := []byte{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}

		splits, _ := SplitSecret(secret, 3, 3)
		reconstruct, _ := Reconstruct(splits)
		assert.Equal(t, secret, reconstruct)
	}

	{
		secret, _ := ioutil.ReadFile("./sample.txt")
		splits, _ := SplitSecret(secret, 5, 3)
		reconstruct, _ := Reconstruct(splits)
		assert.Equal(t, secret, reconstruct)
	}
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
		var points []point

		shouldBePolynomial := func(x int) float64 {
			return float64((10 + 45*x + 102*x*x + 200*x*x*x) % p)
		}

		for i := 1; i < 20; i++ {
			points = append(points, point{
				x: i,
				y: int(shouldBePolynomial(i)),
			})
		}

		assert.Equal(t, shouldBePolynomial(0), reconstructPolynom(points, p)(0), "should be the same")
	}
	//{
	//	// f(x) = 10 + 45 x + 102 x²
	//	points := []point{{1, 157}, {2, 251}, {3, 35} /*{4, 23}, {5, 215}*/}
	//
	//	actualPolynomial := reconstructPolynom(points, p)
	//
	//	shouldBePolynomial := func(x int) float64 {
	//		return float64(modInt(10+45*x+102*x*x, p))
	//	}
	//
	//	// secret
	//	assert.Equal(t, shouldBePolynomial(0), actualPolynomial(0), "Secret Found")
	//}
}

func Test_createBasisPolynomial(t *testing.T) {
	{ // x^2
		var points []point
		for i := 1; i < 4; i++ {
			points = append(points, point{
				x: i,
				y: i * i,
			})
		}

		basisPolynomial := createBasisPolynomial(points, 0, p)
		shouldBePolynomial := func(x float64) float64 {
			return 0.5 * ((x * x) - 5*x + 6)
		}
		assert.Equal(t, shouldBePolynomial(5), basisPolynomial(5))
	}

	{ // f(x) = 10 + 45 x + 102 x²
		points := []point{{1, 157}, {2, 251}, {3, 35} /*{4, 23}, {5, 215}*/}

		basisPolynomial := createBasisPolynomial(points, 0, p)
		assert.Equal(t, modFloat(float64(3), float64(p)), basisPolynomial(0))

		basisPolynomial = createBasisPolynomial(points, 1, p)
		assert.Equal(t, modFloat(float64(-3), float64(p)), basisPolynomial(0))

		basisPolynomial = createBasisPolynomial(points, 2, p)
		assert.Equal(t, modFloat(float64(1), float64(p)), basisPolynomial(0))
	}
}

func Test_modInt(t *testing.T) {
	assert.Equal(t, 2, modInt(10, 4))
	assert.Equal(t, 2, modInt(12, -5))
	assert.Equal(t, 256, modInt(-1, 257))
}
