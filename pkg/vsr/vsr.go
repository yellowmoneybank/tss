package vsr

import (
	"math"
	"math/big"

	"moritzm-mueller.de/tss/pkg/secretSharing"
)

func ValidVSRShare(shares []secretSharing.RedistShare) bool {
	// for _, share := range shares {
	// 	if !feldman.IsValidShare(share.Share) {
	// 		return false
	// 	}
	// }

	for i := 0; i < len(shares[0].Share.Secrets); i++ {
		// index -> g^s_i
		vsrMap := make(map[uint16]uint16)
		for _, share := range shares {
			vsrMap[share.OldIndex] = share.Share.Secrets[i].CheckValues[0]
		}

		if !vsrCheck(shares[0].Share.Secrets[i].GS, vsrMap, shares[0].Share.Q) {
			return false
		}
	}

	return true
}

func vsrCheck(gs uint16, checkValues map[uint16]uint16, q int) bool {
	prod := big.NewFloat(1)
	Q := big.NewInt(int64(q))

	indizes := extractIndizesFromMap(checkValues)

	exponents := initializeExponents(indizes)

	for i, checkvalue := range checkValues {
		exp, exact := exponents[i].Float64()
		if !exact {
			println("Error converting")

			return false
		}

		a := math.Pow(float64(checkvalue), exp)

		prod.Mul(prod, big.NewFloat(a))
		prod = floatMod(prod, q)
	}

	if !prod.IsInt() {
		return false
	}

	pInt, _ := prod.Int64()
	p := big.NewInt(pInt)

	result := p.Mod(p, Q)

	return result == big.NewInt(int64(gs))
}

func floatMod(p *big.Float, q int) *big.Float {
	i, _ := p.Int64()
	a := big.NewInt(i)
	a.Mod(a, big.NewInt(int64(q)))

	// 0.****
	p.Sub(p, big.NewFloat(float64(i)))
	p.Add(p, new(big.Float).SetInt(a))

	return p
}

func extractIndizesFromMap(m map[uint16]uint16) []uint16 {
	var indizes []uint16
	for i := range m {
		indizes = append(indizes, i)
	}

	return indizes
}

func initializeExponents(indices []uint16) map[uint16]*big.Rat {
	exponents := make(map[uint16]*big.Rat)

	for _, i := range indices {
		iInt := int64(i)
		prod := big.NewRat(1, 1)

		for _, l := range indices {
			lInt := int64(l)
			if iInt == lInt {
				continue
			}

			frac := big.NewRat(lInt, lInt-iInt)
			prod.Mul(prod, frac)
		}

		exponents[i] = prod
	}

	return exponents
}
