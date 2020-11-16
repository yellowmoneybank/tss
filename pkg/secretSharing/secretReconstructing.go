package secretSharing

import (
	"fmt"
	"math"
)

func IsUniqueSolution(byteShares map[int]ByteShare, threshold uint8) bool {
	// The Equation System is a Vandermonde-Matrix. There is a unique
	// solution if the Determinant != 0. This is particularly easy for a
	// Vandermonde-Matrix.
	// if the number of shares is higher than the threshold, determine
	// wether there is a combination of shares that has a unique solution
	if len(byteShares) < int(threshold) {
		return false
	}

	var indices []int
	for x := range byteShares {
		indices = append(indices, x)
	}
	// TODO check for all combinations
	return !isDeterminantVandermondeZero(indices[:threshold])
}

func isDeterminantVandermondeZero(indices []int) bool {
	for i := 0; i < len(indices); i++ {
		for j := 0; j < len(indices); j++ {
			if i == j {
				continue
			}

			if indices[i] == indices[j] {
				return true
			}
		}
	}

	return false
}

type Point struct {
	X int
	Y int
}

// polynom interpolation a la Lagrange.
func ReconstructPolynom(points []Point, modulo int) func(int) float64 {
	return func(x int) float64 {
		sum := float64(0)

		for i, point := range points {
			basisPolynom := createBasisPolynomial(points, i, modulo)
			sum += float64(point.Y) * basisPolynom(x)
		}

		return ModFloat(sum, float64(modulo))
	}
}

// TODO this doesnt work for many points. math/big is probably the solution
func createBasisPolynomial(points []Point, index, modulo int) func(int) float64 {
	return func(x int) float64 {
		dividend := 1
		divisor := 1

		for j, point := range points {
			if j == index {
				continue
			}

			dividend = dividend * (x - point.X)
			divisor = divisor * (points[index].X - point.X)
		}

		basisPolynomial := float64(dividend) / float64(divisor)

		if math.IsNaN(basisPolynomial) {
			fmt.Println("stop")
		}

		return ModFloat(basisPolynomial, float64(modulo))
	}
}

func ModInt(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}

	if a < 0 && b > 0 {
		m += b
	}

	return m
}

func ModFloat(a, b float64) float64 {
	m := math.Mod(a, b)
	if a < 0 && b < 0 {
		m -= b
	}

	if a < 0 && b > 0 {
		m += b
	}

	return m
}
