package shamir

import "errors"

func ShamirReconstruct(shares []share) ([]byte, error) {
	// TODO Assertions...
	var secret []byte

	for i := 0; i < len(shares[0].slices); i++ {
		var byteShares []singleByteShare
		for _, share := range shares {
			byteShares := append(byteShares,
				singleByteShare{
					shareIndex: share.shareIndex,
					share:      share.slices[i],
				})
			reconstructedByte, err := reconstructByte(byteShares, share.threshold)
			if err != nil {
				return nil, err
			}
			secret = append(secret, reconstructedByte)
		}
	}

	return secret, nil
}

func reconstructByte(byteShares []singleByteShare, threshold uint8) (byte, error) {
	if !isUniqueSolution(byteShares, threshold) {
		return 0, errors.New("Can't find unique solution")
	}

	var points []point
	for _, byteShare := range byteShares {
		points = append(points, point{int(byteShare.shareIndex), int(byteShare.share)})
	}

	p := reconstructPolynom(points)
	return byte(p(0)), nil
}

func isUniqueSolution(byteShares []singleByteShare, threshold uint8) bool {
	// The Equation System is a Vandermonde-Matrix. There is a unique
	// solution if the Determinant != 0. This is particularly easy for a
	// Vandermonde-Matrix.

	// if the number of shares is higher than the threshold, determine
	// wether there is a combination of shares that has a unique solution

	if len(byteShares) < int(threshold) {
		return false
	}

	var indices []int
	for _, byteshare := range byteShares {
		indices = append(indices, int(byteshare.shareIndex))
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

type point struct {
	x int
	y int
}

// polynom interpolation a la Lagrange
func reconstructPolynom(points []point) func(int) float64 {
	return func(x int) float64 {
		sum := float64(0)
		for i, point := range points {
			basisPolynom := createBasisPolynomial(points, i)
			sum += float64(point.y) * basisPolynom(x)
		}
		return sum
	}
}

//
func createBasisPolynomial(points []point, index int) func(int) float64 {
	return func(x int) float64 {
		dividend := 1
		divisor := 1

		for j, point := range points {
			if j == index {
				continue
			}
			dividend *= (x - point.x)
			divisor *= points[index].x - point.x

		}
		return float64(dividend) / float64(divisor)
	}
}
