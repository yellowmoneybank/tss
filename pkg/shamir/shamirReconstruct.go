package shamir

import (
	"errors"
	"fmt"
	"math"

	"moritzm-mueller.de/tss/pkg/feldman"
	"moritzm-mueller.de/tss/pkg/secretSharing"
)

func Reconstruct(shares []secretSharing.Share) ([]byte, error) {
	// TODO Assertions...
	var secret []byte

	for _, share := range shares {
		if !feldman.IsValidShare(share) {
			return nil, errors.New("share is invalid")
		}
	}

	for i := 0; i < len(shares[0].Secrets); i++ {
		byteShares := make(map[int]secretSharing.ByteShare)
		for _, share := range shares {
			byteShares[int(share.ShareIndex)] = share.Secrets[i]
		}

		reconstructedByte, err := reconstructByte(byteShares, shares[0].Threshold, shares[0].Prime)
		if err != nil {
			return nil, err
		}

		secret = append(secret, reconstructedByte)

		fmt.Printf("reconstructed: %d \r", i)
	}

	return secret, nil
}

func reconstructByte(byteShares map[int]secretSharing.ByteShare, threshold uint8, prime int) (byte, error) {
	if !isUniqueSolution(byteShares, threshold) {
		return 0, errors.New("can't find unique solution")
	}

	var points []point
	for x, byteShare := range byteShares {
		points = append(points, point{x, int(byteShare.Share)})
	}

	p := reconstructPolynom(points, prime)

	secret := p(0)

	if secret != float64(int(secret)) {
		fmt.Println("We have a Problem")
	}

	return byte(secret), nil
}

func isUniqueSolution(byteShares map[int]secretSharing.ByteShare, threshold uint8) bool {
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

type point struct {
	x int
	y int
}

// polynom interpolation a la Lagrange.
func reconstructPolynom(points []point, modulo int) func(int) float64 {
	return func(x int) float64 {
		sum := float64(0)

		for i, point := range points {
			basisPolynom := createBasisPolynomial(points, i, modulo)
			sum += float64(point.y) * basisPolynom(x)
		}

		return modFloat(sum, float64(modulo))
	}
}

// TODO this doesnt work for many points.
func createBasisPolynomial(points []point, index, modulo int) func(int) float64 {
	return func(x int) float64 {
		dividend := 1
		divisor := 1

		for j, point := range points {
			if j == index {
				continue
			}

			dividend = dividend * (x - point.x)
			divisor = divisor * (points[index].x - point.x)
		}

		basisPolynomial := float64(dividend) / float64(divisor)

		if math.IsNaN(basisPolynomial) {
			fmt.Println("stop")
		}

		return modFloat(basisPolynomial, float64(modulo))
	}
}

func modInt(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}

	if a < 0 && b > 0 {
		m += b
	}

	return m
}

func modFloat(a, b float64) float64 {
	m := math.Mod(a, b)
	if a < 0 && b < 0 {
		m -= b
	}

	if a < 0 && b > 0 {
		m += b
	}

	return m
}
